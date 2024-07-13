//go:build !server

package ail

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/timshannon/go-openal/openal"

	"github.com/noxworld-dev/opennox-lib/env"
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

type audioStream struct {
	h Stream
	d *audioDriver

	next    *audioStream
	source  openal.Source
	hwbuf   openal.Buffers
	hwready int

	playing bool

	dec audioDecoder
}

func (s *audioStream) Channels() int {
	return s.dec.Channels()
}

type audioTimer struct {
	h    Timer
	t    *time.Ticker
	dt   time.Duration
	f    func(u uint32)
	user uint32
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

func openAudio(path string) (audioDecoder, error) {
	switch strings.ToLower(filepath.Ext(path)) {
	default:
		return nil, fmt.Errorf("unsupported audio file format: %s", path)
	case ".wav":
		r, err := openWav(path)
		if err != nil {
			return nil, err
		}
		return r, nil
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

	dec, err := openAudio(name)
	if err != nil {
		audioLog.Println(err)
		return 0
	}

	s := &audioStream{
		d:   d,
		h:   Stream(handles.New()),
		dec: dec,
	}
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
	s.dec.Close()
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
	s.dec.Seek(offset)
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
	return s.dec.Position()
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
		decoded = decodeADPCMMono(decoded, data)
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
	sampleRate := s.dec.SampleRate()
	minSamples := sampleRate * channels / 10

	s.unqueueBuffers()

	for s.hwready != 0 {
		offset := 0

		for offset < minSamples {
			samples, ok := s.dec.Decode(buffer[offset : minSamples*2])
			if !ok {
				s.playing = false
			}
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
		if channels != 1 {
			format = openal.FormatStereo16
		}

		s.hwbuf[s.hwready].SetDataInt16(format, buffer[:offset], int32(sampleRate))
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
