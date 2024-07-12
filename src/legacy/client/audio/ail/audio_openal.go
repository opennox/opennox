//go:build !server

package ail

/*
#include <stdint.h>
#define MINIMP3_ONLY_MP3
#define MINIMP3_NO_SIMD
#include "../mp3/minimp3.h"
*/
import "C"
import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/timshannon/go-openal/openal"

	"github.com/noxworld-dev/opennox-lib/env"
	"github.com/noxworld-dev/opennox-lib/ifs"
	"github.com/noxworld-dev/opennox-lib/log"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
)

var (
	audioLog     = log.New("audio")
	audioDebug   = os.Getenv("NOX_DEBUG_AUDIO") == "true"
	audioSamples struct {
		sync.RWMutex
		byHandle map[Sample]*audioSample
	}
	audioStreams struct {
		sync.RWMutex
		byHandle map[Stream]*audioStream
	}
	audioDrivers struct {
		sync.RWMutex
		byHandle map[Driver]*audioDriver
	}
	audioTimers struct {
		sync.RWMutex
		byHandle map[Timer]*audioTimer
	}
)

func init() {
	audioSamples.byHandle = make(map[Sample]*audioSample)
	audioStreams.byHandle = make(map[Stream]*audioStream)
	audioDrivers.byHandle = make(map[Driver]*audioDriver)
	audioTimers.byHandle = make(map[Timer]*audioTimer)
}

func (dig Driver) get() *audioDriver {
	if dig == 0 {
		return nil
	}
	handles.AssertValid(uintptr(dig))
	audioDrivers.RLock()
	s := audioDrivers.byHandle[dig]
	audioDrivers.RUnlock()
	return s
}

func (h Stream) get() *audioStream {
	if h == 0 {
		return nil
	}
	handles.AssertValid(uintptr(h))
	audioStreams.RLock()
	s := audioStreams.byHandle[h]
	audioStreams.RUnlock()
	return s
}

func (h Sample) get() *audioSample {
	if h == 0 {
		return nil
	}
	handles.AssertValid(uintptr(h))
	audioSamples.RLock()
	s := audioSamples.byHandle[h]
	audioSamples.RUnlock()
	return s
}

func (h Timer) get() *audioTimer {
	if h == 0 {
		return nil
	}
	handles.AssertValid(uintptr(h))
	audioTimers.RLock()
	s := audioTimers.byHandle[h]
	audioTimers.RUnlock()
	return s
}

type audioDriver struct {
	h          Driver
	dev        *openal.Device
	ctx        *openal.Context
	mu         sync.Mutex
	sampleHead *audioSample
	streamHead *audioStream
	t          *time.Ticker
}

type audioSampleBuf struct {
	buf []byte
	pos int
}

type audioSample struct {
	h       Sample
	d       *audioDriver
	playing uint32 // atomic

	next    *audioSample
	source  openal.Source
	hwbuf   openal.Buffers
	hwready int

	bmu     sync.Mutex
	buffers [2]audioSampleBuf
	current byte
	ready   byte

	stereo bool
	adpcm  bool

	playbackRate uint32
	blockSize    uint32

	eob  func()
	eos  func()
	user any
}

func (s *audioSample) IsPlaying() bool {
	return atomic.LoadUint32(&s.playing) != 0
}

func (s *audioSample) Play() {
	atomic.StoreUint32(&s.playing, 1)
}

func (s *audioSample) Stop() {
	atomic.StoreUint32(&s.playing, 0)
}

func (s *audioSample) EOS() {
	s.Stop()
	if s.eos != nil {
		s.d.mu.Unlock()
		s.eos()
		s.d.mu.Lock()
	}
}

func (s *audioSample) EOB() {
	s.bmu.Lock()
	s.ready++
	if s.current == 0 {
		s.current = 1
	} else {
		s.current = 0
	}
	s.bmu.Unlock()
	if s.eob != nil {
		s.d.mu.Unlock()
		s.eob()
		s.d.mu.Lock()
	}
}

type audioStreamType int

const (
	audioPCM = audioStreamType(iota + 1)
	audioADPCM
	audioMP3
)

type audioStream struct {
	h Stream
	d *audioDriver

	next    *audioStream
	source  openal.Source
	hwbuf   openal.Buffers
	hwready int

	stereo       bool
	playbackRate uint32
	playing      bool

	file      *os.File
	filename  string
	fileSize  int
	chunkSize int
	chunkPos  int
	buffered  int

	typ audioStreamType
	pcm struct {
		wf       waveFormat
		position int
	}
	adpcm struct {
		wf       waveFormat
		position int
		samples  int
	}
	mp3 struct {
		wf  waveFormatMP3
		dec C.mp3dec_t
	}

	buffer [16 * 1024 * 3]byte
}

func (s *audioStream) Channels() int {
	if s.stereo {
		return 2
	}
	return 1
}

func (s *audioStream) findData() error {
	var tmp [8]byte

	s.chunkSize = 0
	s.chunkPos = 0

	for {
		off, err := s.file.Seek(0, io.SeekCurrent)
		if err != nil {
			audioLog.Println(err)
			return err
		} else if int(off) >= s.fileSize {
			return io.EOF
		}
		_, err = io.ReadFull(s.file, tmp[:])
		if err != nil {
			audioLog.Println("find", err)
			return err
		}
		size := int(binary.LittleEndian.Uint32(tmp[4:]))
		if string(tmp[:4]) == "data" {
			s.chunkSize = size
			return nil
		}
		_, err = s.file.Seek(int64(size), io.SeekCurrent)
		if err != nil {
			audioLog.Println(err)
			return err
		}
	}
}

type audioTimer struct {
	h    Timer
	t    *time.Ticker
	dt   time.Duration
	f    func(u uint32)
	user uint32
}

func sat16(x int) int16 {
	if x < math.MinInt16 {
		return math.MinInt16
	} else if x > math.MaxInt16 {
		return math.MaxInt16
	}
	return int16(x)
}

func satindex(x int) int {
	if x < 0 {
		return 0
	} else if x > 88 {
		return 88
	}
	return x
}

var (
	imaIndexTable = [16]int{-1, -1, -1, -1, 2, 4, 6, 8, -1, -1, -1, -1, 2, 4, 6, 8}
	imaStepTable  = [89]int{
		7, 8, 9, 10, 11, 12, 13, 14, 16, 17, 19, 21, 23,
		25, 28, 31, 34, 37, 41, 45, 50, 55, 60, 66, 73, 80,
		88, 97, 107, 118, 130, 143, 157, 173, 190, 209, 230, 253, 279,
		307, 337, 371, 408, 449, 494, 544, 598, 658, 724, 796, 876, 963,
		1060, 1166, 1282, 1411, 1552, 1707, 1878, 2066, 2272, 2499, 2749, 3024, 3327,
		3660, 4026, 4428, 4871, 5358, 5894, 6484, 7132, 7845, 8630, 9493, 10442, 11487,
		12635, 13899, 15289, 16818, 18500, 20350, 22385, 24623, 27086, 29794, 32767,
	}
)

func decodeNibble(predictor int16, nibble byte, index int) int16 {
	diff := 0
	step := imaStepTable[index]

	diff = step >> 3
	if nibble&4 != 0 {
		diff += step
	}
	if nibble&2 != 0 {
		diff += step >> 1
	}
	if nibble&1 != 0 {
		diff += step >> 2
	}
	if nibble&8 != 0 {
		predictor = sat16(int(predictor) - diff)
	} else {
		predictor = sat16(int(predictor) + diff)
	}
	return predictor
}

func decodeADPCM(out []int16, data []byte) []int16 {
	predictor := int16(data[0]) | (int16(data[1]) << 8)
	index := int(data[2])

	out = append(out, predictor)
	for i := 4; i < len(data); i++ {
		predictor = decodeNibble(predictor, data[i], index)
		index = satindex(index + imaIndexTable[data[i]&15])
		out = append(out, predictor)

		predictor = decodeNibble(predictor, data[i]>>4, index)
		index = satindex(index + imaIndexTable[data[i]>>4])
		out = append(out, predictor)
	}
	return out
}

func decodeADPCMStereo(out []int16, data []byte) []int16 {
	lpredictor := int16(data[0]) | (int16(data[1]) << 8)
	lindex := int(data[2])
	rpredictor := int16(data[4]) | (int16(data[5]) << 8)
	rindex := int(data[6])

	out = append(out, lpredictor, rpredictor)
	for i := 8; i < len(data); i += 8 {
		for j := 0; j < 4; j++ {
			lpredictor = decodeNibble(lpredictor, data[i+j], lindex)
			lindex = satindex(lindex + imaIndexTable[data[i+j]&15])
			out = append(out, lpredictor)

			rpredictor = decodeNibble(rpredictor, data[i+j+4], rindex)
			rindex = satindex(rindex + imaIndexTable[data[i+j+4]&15])
			out = append(out, rpredictor)

			lpredictor = decodeNibble(lpredictor, data[i+j]>>4, lindex)
			lindex = satindex(lindex + imaIndexTable[data[i+j]>>4])
			out = append(out, lpredictor)

			rpredictor = decodeNibble(rpredictor, data[i+j+4]>>4, rindex)
			rindex = satindex(rindex + imaIndexTable[data[i+j+4]>>4])
			out = append(out, rpredictor)
		}
	}
	return out
}

func (s *audioStream) seek(pos int) {
	switch s.typ {
	case audioPCM:
		s.pcmSeek(pos)
	case audioADPCM:
		s.adpcmSeek(pos)
	case audioMP3:
		s.mp3Seek(pos)
	default:
		panic("unsupported")
	}
}

func (s *audioStream) tell() uint32 {
	switch s.typ {
	case audioPCM:
		return s.pcmTell()
	case audioADPCM:
		return s.adpcmTell()
	case audioMP3:
		return s.mp3Tell()
	default:
		panic("unsupported")
	}
}

func (s *audioStream) decode(out []int16, max int) uint32 {
	switch s.typ {
	case audioPCM:
		return s.pcmDecode(out, max)
	case audioADPCM:
		return s.adpcmDecode(out, max)
	case audioMP3:
		return s.mp3Decode(out, max)
	default:
		panic("unsupported")
	}
}

func (s *audioStream) pcmDecode(out []int16, max int) uint32 {
	remaining := s.chunkSize - s.chunkPos

	if remaining == 0 {
		remaining = s.decodeFind()
		if remaining == 0 {
			return 0
		}
	}

	if remaining/2 >= max {
		remaining = max * 2
	}

	channels := s.Channels()
	// TODO(dennwc): reuse
	buf := make([]byte, remaining)
	io.ReadFull(s.file, buf)
	for i := 0; i+1 < len(buf); i += 2 {
		out[i/2] = int16(binary.LittleEndian.Uint16(buf[i:]))
	}
	s.pcm.position += remaining / 2 / channels
	return uint32(remaining / 2)
}

func (s *audioStream) pcmSeek(position int) {
	// TODO
	s.pcm.position = 0
}

func (s *audioStream) pcmTell() uint32 {
	return uint32(s.pcm.position)
}

func (s *audioStream) decodeFind() int {
	s.findData()
	remaining := s.chunkSize - s.chunkPos
	if remaining == 0 {
		s.playing = false
	}
	return remaining
}

func (s *audioStream) adpcmDecode(out []int16, max int) uint32 {
	blockSize := int(s.adpcm.wf.blockAlign)

	samples := 0
	if s.adpcm.position >= s.adpcm.samples {
		s.playing = false
		return 0
	}

	if s.buffered < blockSize {
		remaining := s.chunkSize - s.chunkPos

		if remaining == 0 {
			remaining = s.decodeFind()
			if remaining == 0 {
				return 0
			}
		}

		if remaining > blockSize-s.buffered {
			remaining = blockSize - s.buffered
		}
		io.ReadFull(s.file, s.buffer[s.buffered:s.buffered+remaining])
		s.buffered += remaining
	}

	if s.stereo {
		out = decodeADPCMStereo(out[:0], s.buffer[:blockSize])
	} else {
		out = decodeADPCM(out[:0], s.buffer[:blockSize])
	}
	samples = len(out) // TODO(dennwc): is it true for stereo?

	s.buffered -= blockSize
	if s.buffered != 0 {
		copy(s.buffer[0:], s.buffer[blockSize:blockSize+s.buffered])
	}
	channels := s.Channels()
	if samples/channels+s.adpcm.position >= s.adpcm.samples {
		samples = (s.adpcm.samples - s.adpcm.position) * channels
	}
	s.adpcm.position += samples / channels
	return uint32(samples)
}

func (s *audioStream) adpcmSeek(position int) {
	blockSize := int(s.adpcm.wf.blockAlign)
	channels := s.Channels()
	samplesPerBlock := (blockSize/channels - 4) * 2
	blocks := position / samplesPerBlock
	s.adpcm.position = 0

	for s.adpcm.position < position {
		s.findData()
		if s.chunkSize == 0 {
			break
		}
		chunkBlocks := s.chunkSize / blockSize
		if blocks < chunkBlocks {
			s.file.Seek(int64(blocks*blockSize), io.SeekCurrent)
			s.adpcm.position += blocks * samplesPerBlock
			break
		} else {
			s.file.Seek(int64(s.chunkSize), io.SeekCurrent)
			s.adpcm.position += chunkBlocks * samplesPerBlock
		}
	}

	// TODO seek within an ADPCM block
}

func (s *audioStream) adpcmTell() uint32 {
	return uint32(s.adpcm.position)
}

func (s *audioStream) mp3Decode(out []int16, max int) uint32 {
	remaining := s.chunkSize - s.chunkPos
	samples := 0
	var info C.mp3dec_frame_info_t

	if s.buffered < len(s.buffer) {
		if remaining == 0 && s.buffered == 0 {
			remaining = s.decodeFind()
			if remaining == 0 {
				return 0
			}
		}
		if remaining != 0 {
			if remaining > len(s.buffer)-s.buffered {
				remaining = len(s.buffer) - s.buffered
			}
			io.ReadFull(s.file, s.buffer[s.buffered:s.buffered+remaining])
			s.buffered += remaining
		}
	}

	for {
		samples = int(C.mp3dec_decode_frame(&s.mp3.dec, (*C.uchar)(unsafe.Pointer(&s.buffer[0])), C.int(s.buffered), (*C.short)(unsafe.Pointer(&out[0])), &info))
		if samples != 0 {
			break
		}
		if remaining == 0 {
			remaining = s.decodeFind()
			if remaining == 0 {
				return 0
			}
		}
		if remaining != 0 {
			if remaining > len(s.buffer)-s.buffered {
				remaining = len(s.buffer) - s.buffered
			}
			io.ReadFull(s.file, s.buffer[s.buffered:s.buffered+remaining])
			s.buffered += remaining
		}
	}

	frameBytes := int(info.frame_bytes)
	s.buffered -= frameBytes
	if s.buffered != 0 {
		copy(s.buffer[0:], s.buffer[frameBytes:frameBytes+s.buffered])
	}

	return uint32(samples)
}

func (s *audioStream) mp3Seek(position int) {
	C.mp3dec_init(&s.mp3.dec)
}

func (s *audioStream) mp3Tell() uint32 {
	return 0
}

func audioCheckError() bool {
	if err := openal.Err(); err != nil {
		audioLog.Println(err)
		return false
	}
	return true
}

func (s *audioSample) unqueueBuffers() {
	processed := int(s.source.Geti(openal.AlBuffersProcessed))
	if !audioCheckError() {
		return
	}
	if processed != 0 {
		s.source.UnqueueBuffers(s.hwbuf[s.hwready : s.hwready+processed])
		if !audioCheckError() {
			return
		}
	}
	s.hwready += processed
}

func (s *audioStream) unqueueBuffers() {
	processed := int(s.source.Geti(openal.AlBuffersProcessed))
	if !audioCheckError() {
		return
	}
	if processed != 0 {
		s.source.UnqueueBuffers(s.hwbuf[s.hwready : s.hwready+processed])
		if !audioCheckError() {
			return
		}
	}
	s.hwready += processed
}

func (dig Driver) AllocateSample() Sample {
	if env.IsE2E() {
		h := handles.New()
		return Sample(h)
	}
	if audioDebug {
		audioLog.Println("AIL_allocate_sample_handle")
	}
	d := dig.get()
	if d == nil {
		return 0
	}

	s := &audioSample{
		h:      Sample(handles.New()),
		d:      d,
		source: openal.NewSource(),
	}
	if !audioCheckError() {
		return 0
	}

	s.hwbuf = openal.NewBuffers(4)
	s.hwready = 4
	if !audioCheckError() {
		s.source.Delete()
		return 0
	}

	d.mu.Lock()
	s.next = d.sampleHead
	d.sampleHead = s
	d.mu.Unlock()

	audioSamples.Lock()
	audioSamples.byHandle[s.h] = s
	audioSamples.Unlock()
	return s.h
}

func (s Sample) Release() {
	if audioDebug {
		audioLog.Println("AIL_release_sample_handle")
	}
	handles.AssertValid(uintptr(s))
	if v := audioSamples.byHandle[s]; v != nil {
		v.source.Delete()
		delete(audioSamples.byHandle, s)
		if len(v.hwbuf) > 0 {
			v.hwbuf.Delete()
		}
	}
}

func audioOpenStream(path string) (*audioStream, error) {
	f, err := ifs.Open(path)
	if err != nil {
		return nil, err
	}
	s := &audioStream{file: f, filename: path}
	if err = s.open(); err != nil {
		_ = f.Close()
		return nil, err
	}
	return s, nil
}

type waveFormat struct {
	format         uint16
	channels       uint16
	samplesPerSec  uint32
	avgBytesPerSec uint32
	blockAlign     uint16
}

type waveFormatMP3 struct {
	waveFormat
	bitsPerSample  uint16
	size           uint16
	id             uint16
	flags          uint32
	blockSize      uint16
	framesPerBlock uint16
	codecDelay     uint16
}

func (s *audioStream) open() error {
	var tmp [256]byte
	_, err := io.ReadFull(s.file, tmp[:20])
	if err != nil {
		return err
	}
	if magic := string(tmp[0:4]); magic != "RIFF" {
		return fmt.Errorf("invalid audio file magic: %q", magic)
	}
	s.fileSize = int(binary.LittleEndian.Uint32(tmp[4:8]))
	if str := string(tmp[8:12]); str != "WAVE" {
		return fmt.Errorf("invalid audio file section: %q", str)
	}
	// +12
	size := int(binary.LittleEndian.Uint32(tmp[16:20]))
	if str := string(tmp[12:16]); str != "fmt " {
		return fmt.Errorf("invalid audio file section 2: %q", str)
	}
	if size < 14 || size > cap(tmp) {
		return fmt.Errorf("invalid audio header size: %d", size)
	}
	_, err = io.ReadFull(s.file, tmp[:size])
	if err != nil {
		return err
	}
	wf := waveFormat{
		format:         binary.LittleEndian.Uint16(tmp[0:2]),
		channels:       binary.LittleEndian.Uint16(tmp[2:4]),
		samplesPerSec:  binary.LittleEndian.Uint32(tmp[4:8]),
		avgBytesPerSec: binary.LittleEndian.Uint32(tmp[8:12]),
		blockAlign:     binary.LittleEndian.Uint16(tmp[12:14]),
	}
	switch wf.format {
	case 0x01: // PCM
		s.typ = audioPCM
		s.playbackRate = wf.samplesPerSec
		s.stereo = wf.channels > 1
		s.pcm.wf = wf
		audioLog.Printf("PCM stream: %q, %d channels", s.filename, wf.channels)
		return nil
	case 0x11: // ADPCM
		s.typ = audioADPCM
		s.playbackRate = wf.samplesPerSec
		s.stereo = wf.channels > 1
		s.adpcm.wf = wf
		_, err = io.ReadFull(s.file, tmp[:12])
		if err != nil {
			return err
		}
		if str := string(tmp[0:4]); str != "fact" {
			return fmt.Errorf("invalid ADPCM file section: %q", str)
		}
		if sz := binary.LittleEndian.Uint32(tmp[4:8]); sz != 4 {
			return fmt.Errorf("unsupported ADPCM size: %q", sz)
		}
		s.adpcm.samples = int(binary.LittleEndian.Uint32(tmp[8:12]))
		audioLog.Printf("ADPCM stream: %q, %d channels", s.filename, wf.channels)
		return nil
	case 0x55: // MP3
		if size < 14+16 {
			return fmt.Errorf("invalid MP3 header size: %d", size)
		}
		s.typ = audioMP3
		wf3 := waveFormatMP3{waveFormat: wf}
		wf3.bitsPerSample = binary.LittleEndian.Uint16(tmp[14:16])
		wf3.size = binary.LittleEndian.Uint16(tmp[16:18])
		wf3.id = binary.LittleEndian.Uint16(tmp[18:20])
		wf3.flags = binary.LittleEndian.Uint32(tmp[20:24])
		wf3.blockSize = binary.LittleEndian.Uint16(tmp[24:26])
		wf3.framesPerBlock = binary.LittleEndian.Uint16(tmp[26:28])
		wf3.codecDelay = binary.LittleEndian.Uint16(tmp[28:30])
		s.playbackRate = wf3.samplesPerSec
		s.stereo = wf3.channels > 1
		s.mp3.wf = wf3
		C.mp3dec_init(&s.mp3.dec)
		audioLog.Printf("MP3 stream: %q, %d channels", s.filename, wf3.channels)
		return nil
	default:
		return fmt.Errorf("unsupported format: 0x%x", wf.format)
	}
}

func (dig Driver) OpenStream(name string) Stream {
	if audioDebug {
		audioLog.Println("AIL_open_stream")
	}
	if env.IsE2E() {
		h := handles.New()
		return Stream(h)
	}
	d := dig.get()
	if d == nil {
		return 0
	}

	s, err := audioOpenStream(name)
	if err != nil {
		audioLog.Println(err)
		return 0
	}
	s.d = d
	s.h = Stream(handles.New())
	s.source = openal.NewSource()
	if !audioCheckError() {
		return 0
	}
	s.hwbuf = openal.NewBuffers(2)
	s.hwready = 2
	if !audioCheckError() {
		s.source.Delete()
		return 0
	}

	d.mu.Lock()
	s.next = d.streamHead
	d.streamHead = s
	d.mu.Unlock()

	audioStreams.Lock()
	audioStreams.byHandle[s.h] = s
	audioStreams.Unlock()
	return s.h
}

func (h Stream) Close() error {
	if audioDebug {
		audioLog.Println("AIL_close_stream")
	}
	if env.IsE2E() {
		handles.AssertValid(uintptr(h))
		return nil
	}
	s := h.get()
	if s == nil {
		return nil
	}
	s.d.mu.Lock()
	s.playing = false
	s.d.mu.Unlock()
	return nil
}

func RegisterTimer(f func(u uint32)) Timer {
	if audioDebug {
		audioLog.Println("AIL_register_timer")
	}
	if env.IsE2E() {
		h := handles.New()
		return Timer(h)
	}
	h := Timer(handles.New())
	t := &audioTimer{h: h, f: f}

	audioTimers.Lock()
	audioTimers.byHandle[h] = t
	audioTimers.Unlock()
	return h
}

func (h Timer) Release() {
	if audioDebug {
		audioLog.Println("AIL_release_timer_handle")
	}
	handles.AssertValid(uintptr(h))
	if t := audioTimers.byHandle[h]; t != nil {
		if t.t != nil {
			t.t.Stop()
		}
		delete(audioTimers.byHandle, h)
	}
}

func (h Sample) GetSource() *uint32 {
	if env.IsE2E() {
		return nil
	}
	s := h.get()
	if s == nil {
		return nil
	}
	return (*uint32)(&s.source)
}

func (h Sample) End() {
	if audioDebug {
		audioLog.Println("AIL_end_sample")
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}

	s.d.mu.Lock()
	defer s.d.mu.Unlock()
	if s.IsPlaying() {
		s.EOB()
		s.EOS()
	}
}

func (h Sample) Init() {
	if audioDebug {
		audioLog.Println("AIL_init_sample")
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}

	s.d.mu.Lock()
	defer s.d.mu.Unlock()

	s.bmu.Lock()
	s.current = 0
	s.ready = 2
	s.bmu.Unlock()

	s.source.Setf(openal.AlPitch, 1)
	s.source.Setf(openal.AlGain, 1)
	h.SetPan(63)

	s.Stop()
	s.source.Stop()
	s.unqueueBuffers()
}

func LastError() string {
	if audioDebug {
		audioLog.Println("AIL_last_error")
	}
	return ""
}

func (h Sample) LoadBuffer(num uint32, buf []byte) {
	if audioDebug {
		audioLog.Printf("AIL_load_sample_buffer: [%d]", len(buf))
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}
	s.bmu.Lock()
	defer s.bmu.Unlock()
	sb := &s.buffers[num]
	sb.buf = buf
	sb.pos = 0
	s.Play()
}

func (h Stream) Pause(pause bool) {
	if audioDebug {
		audioLog.Println("AIL_pause_stream")
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}
	// TODO: mutex?
	s.playing = !pause
}

func (h Sample) RegisterEOBCallback(f func()) {
	if audioDebug {
		audioLog.Println("AIL_register_EOB_callback")
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}
	s.eob = f
}

func (h Sample) RegisterEOSCallback(f func()) {
	if audioDebug {
		audioLog.Println("AIL_register_EOS_callback")
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}
	s.eos = f
}

func (h Sample) BufferReady() int {
	if audioDebug {
		audioLog.Println("AIL_sample_buffer_ready")
	}
	if env.IsE2E() {
		return -1
	}
	s := h.get()
	if s == nil {
		return -1
	}
	s.bmu.Lock()
	defer s.bmu.Unlock()
	if s.ready == 0 {
		return -1
	}
	if s.ready == 2 {
		s.ready--
		return 0
	}
	s.ready--
	if s.current == 0 {
		return 1
	}
	return 0
}

func (h Sample) UserData() any {
	if audioDebug {
		audioLog.Println("AIL_sample_user_data")
	}
	if env.IsE2E() {
		return nil
	}
	s := h.get()
	if s == nil {
		return nil
	}
	return s.user
}

func Serve() {
	if audioDebug {
		audioLog.Println("AIL_serve")
	}
	for _, d := range audioDrivers.byHandle {
		select {
		case <-d.t.C:
			d.doWork()
		default:
		}
	}
	for _, t := range audioTimers.byHandle {
		select {
		case <-t.t.C:
			t.f(t.user)
			t.t.Reset(t.dt)
		default:
		}
	}
}

func (h Sample) SetADPCMBlockSize(block uint32) {
	if audioDebug {
		audioLog.Println("AIL_set_sample_adpcm_block_size")
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}
	s.blockSize = block
}

func (h Sample) SetPan(pan int) {
	if audioDebug {
		audioLog.Println("AIL_set_sample_pan")
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}
	pos := [3]float32{float32(pan-63) / 64.0, 0, 0}
	pos[2] = float32(math.Sqrt(float64(1 - pos[0]*pos[0])))
	s.source.Setfv(openal.AlPosition, pos[:])
}

func (h Sample) SetPlaybackRate(rate int) {
	if audioDebug {
		audioLog.Println("AIL_set_sample_playback_rate")
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}
	s.playbackRate = uint32(rate)
}

func (h Sample) SetType(format int32, flags uint32) {
	if audioDebug {
		audioLog.Println("AIL_set_sample_type")
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}
	if flags != 0 {
		panic("abort")
	}
	s.stereo = format&2 != 0
	s.adpcm = format&4 != 0
}

func (h Sample) SetUserData(value any) {
	if audioDebug {
		audioLog.Println("AIL_set_sample_user_data")
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}
	s.user = value
}

func (h Sample) SetVolume(volume int) {
	if audioDebug {
		audioLog.Println("AIL_set_sample_volume")
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}
	s.source.Setf(openal.AlGain, float32(volume)/127.0)
}

func (h Stream) SetPosition(offset int) {
	if audioDebug {
		audioLog.Println("AIL_set_stream_position")
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}
	s.d.mu.Lock()
	defer s.d.mu.Unlock()
	s.chunkSize = 0
	s.chunkPos = 0
	s.buffered = 0
	s.file.Seek(12, io.SeekStart)
	s.seek(offset)
}

func (h Stream) SetVolume(volume int) {
	if audioDebug {
		audioLog.Println("AIL_set_stream_volume", int(volume))
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}
	s.source.Setf(openal.AlGain, float32(volume)/127.0)
}

func (h Timer) SetFrequency(hertz uint) {
	if audioDebug {
		audioLog.Println("AIL_set_timer_frequency")
	}
	if env.IsE2E() {
		return
	}
	t := h.get()
	if t == nil {
		return
	}
	t.dt = time.Duration(1000.0/float32(hertz)) * time.Millisecond
}

func (h Stream) Start() {
	if audioDebug {
		audioLog.Println("AIL_start_stream")
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}
	s.playing = true
}

func (h Timer) Start() {
	if audioDebug {
		audioLog.Println("AIL_start_timer")
	}
	if env.IsE2E() {
		return
	}
	t := h.get()
	if t == nil {
		return
	}
	if t.t != nil {
		t.t.Stop()
	}
	t.t = time.NewTicker(t.dt)
}

func Startup() int {
	if audioDebug {
		audioLog.Println("AIL_startup")
	}
	return -1
}

func Shutdown() {
	if audioDebug {
		audioLog.Println("AIL_shutdown")
	}
	if env.IsE2E() {
		return
	}
	audioTimers.Lock()
	for _, t := range audioTimers.byHandle {
		if t.t != nil {
			t.t.Stop()
		}
	}
	audioTimers.Unlock()
	audioDrivers.Lock()
	for _, d := range audioDrivers.byHandle {
		if d.t != nil {
			d.t.Stop()
		}
	}
	audioDrivers.Unlock()
}

func (h Sample) Stop() {
	if audioDebug {
		audioLog.Println("AIL_stop_sample")
	}
	if env.IsE2E() {
		return
	}
	s := h.get()
	if s == nil {
		return
	}
	s.Stop() // TODO: anything else here?
}

func (h Timer) Stop() {
	if audioDebug {
		audioLog.Println("AIL_stop_timer")
	}
	if env.IsE2E() {
		return
	}
	t := h.get()
	if t == nil {
		return
	}
	if t.t != nil {
		t.t.Stop()
	}
}

func (h Stream) Position() int {
	if audioDebug {
		audioLog.Println("AIL_stream_position")
	}
	if env.IsE2E() {
		return -1
	}
	s := h.get()
	if s == nil {
		return -1
	}
	return int(s.tell())
}

func (h Stream) Status() int {
	if env.IsE2E() {
		return 2
	}
	s := h.get()
	if s == nil {
		return 2
	}
	if s.playing {
		return 4
	}
	return 2
}

func (dig Driver) Close() error {
	if audioDebug {
		audioLog.Println("AIL_waveOutClose")
	}
	if env.IsE2E() {
		return nil
	}
	d := dig.get()
	if d == nil {
		return nil
	}
	if d.t != nil {
		d.t.Stop()
	}
	return nil
}

func (s *audioSample) playADPCM(data []byte) {
	if !s.adpcm {
		panic("abort")
	}

	var decoded []int16
	if s.stereo {
		decoded = decodeADPCMStereo(decoded, data)
	} else {
		decoded = decodeADPCM(decoded, data)
	}

	s.hwready--

	format := openal.FormatMono16
	if s.stereo {
		format = openal.FormatStereo16
	}

	s.hwbuf[s.hwready].SetDataInt16(format, decoded, int32(s.playbackRate))
	if !audioCheckError() {
		s.hwready++
		return
	}
	s.source.QueueBuffer(s.hwbuf[s.hwready])
	if !audioCheckError() {
		s.hwready++
		return
	}
}

func (s *audioSample) work() {
	if !s.IsPlaying() {
		return
	}

	s.bmu.Lock()
	if s.ready == 2 {
		s.bmu.Unlock()
		s.EOS()
		return
	}

	buf := &s.buffers[s.current]
	if len(buf.buf) == 0 {
		s.bmu.Unlock()
		s.EOB()
		s.EOS()
		return
	}
	defer s.bmu.Unlock()
	s.unqueueBuffers()

	for s.hwready != 0 {
		s.playADPCM(buf.buf[buf.pos : buf.pos+int(s.blockSize)])
		buf.pos += int(s.blockSize)

		if s.IsPlaying() && s.source.State() != openal.Playing {
			s.source.Play()
		}

		if buf.pos >= len(buf.buf) {
			s.bmu.Unlock()
			s.EOB()
			s.bmu.Lock()
			buf = &s.buffers[s.current]
			if len(buf.buf) == 0 {
				break
			}
		}
	}
}

func (s *audioStream) work() {
	if !s.playing {
		return
	}
	var buffer [16 * 1024 * 3]int16
	channels := s.Channels()
	// We need to queue 100ms of audio data.
	minSamples := int(s.playbackRate) * channels / 10

	s.unqueueBuffers()

	for s.hwready != 0 {
		offset := 0

		for offset < minSamples {
			samples := s.decode(buffer[offset:], minSamples*2-offset)
			if samples == 0 {
				break
			}
			offset += int(samples)
		}

		if offset == 0 {
			break
		}

		s.hwready--

		format := openal.FormatMono16
		if s.stereo {
			format = openal.FormatStereo16
		}

		s.hwbuf[s.hwready].SetDataInt16(format, buffer[:offset], int32(s.playbackRate))
		if !audioCheckError() {
			s.hwready++
			return
		}
		s.source.QueueBuffer(s.hwbuf[s.hwready])
		if !audioCheckError() {
			s.hwready++
			return
		}

		state := s.source.State()
		if s.playing && state != openal.Playing {
			s.source.Play()
		}
	}
}

func (d *audioDriver) doWork() {
	d.mu.Lock()
	defer d.mu.Unlock()
	for s := d.sampleHead; s != nil; s = s.next {
		s.work()
	}
	for s := d.streamHead; s != nil; s = s.next {
		s.work()
	}
}

func WaveOutOpen() Driver {
	if audioDebug {
		audioLog.Println("AIL_waveOutOpen")
	}
	if env.IsE2E() {
		return 0
	}
	h := Driver(handles.New())

	d := &audioDriver{h: h}
	d.dev = openal.OpenDevice("")
	if d.dev == nil {
		if err := openal.Err(); err != nil {
			audioLog.Println("error opening device:", err)
		} else {
			audioLog.Println("cannot open device")
		}
		return 0
	}
	audioLog.Println("device ok")
	d.ctx = d.dev.CreateContext()
	if d.ctx == nil {
		if err := openal.Err(); err != nil {
			audioLog.Println("error creating context:", err)
		} else {
			audioLog.Println("cannot create context")
		}
		return 0
	}
	audioLog.Println("context ok")
	d.ctx.Activate()
	if err := openal.Err(); err != nil {
		audioLog.Println("error activating context:", err)
		return 0
	}

	openal.Listener{}.Setf(openal.AlGain, 1)
	openal.Listener{}.Setfv(openal.AlPosition, []float32{0, 0, 0})
	d.t = time.NewTicker(time.Second / 20)

	audioDrivers.Lock()
	audioDrivers.byHandle[h] = d
	audioDrivers.Unlock()

	return h
}
