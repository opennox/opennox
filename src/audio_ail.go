package opennox

import (
	"encoding/binary"
	"fmt"
	"math"
	"path/filepath"
	"sort"
	"strings"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/ifs"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/internal/binfile"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
)

var (
	nox_enable_audio     = 1
	dword_5d4594_805980  *audioStructXxx
	dword_5d4594_1193336 uint32
)

var (
	audioTimer93944 ail.Timer = math.MaxUint32
	audioDev        ail.Driver
)

func sub_43EFD0(a1 unsafe.Pointer) int {
	s := *(**legacy.AudioSample)(unsafe.Add(a1, 272))
	s.Smp.End()
	if s.Flag7 == 0 {
		ptr := s.Field1
		fptr := (*unsafe.Pointer)(unsafe.Add(ptr, 284))
		ccall.CallVoidPtr(*fptr, ptr)
		s.Flag7 = 1
	}
	return 0
}

func sub_43EC30(a1 unsafe.Pointer) int {
	smp := audioDev.AllocateSample()
	s, _ := alloc.New(legacy.AudioSample{})
	*(*unsafe.Pointer)(unsafe.Add(a1, 272)) = unsafe.Pointer(s)
	s.Dev = audioDev
	s.Field1 = a1
	s.Smp = smp
	b1, _ := alloc.Make([]byte{}, 16384)
	b2, _ := alloc.Make([]byte{}, 16384)
	s.Data1 = &b1[0]
	s.Data2 = &b2[0]
	if smp == 0 {
		return -2147221504 // 0x80040000
	}
	smp.SetUserData(s)
	return 0
}

func sub_43ECB0(a1 unsafe.Pointer) int {
	p := *(**legacy.AudioSample)(unsafe.Add(a1, 272))
	if p.Smp != 0 {
		p.Smp.Release()
	}
	if p2 := p.Data1; p2 != nil {
		alloc.Free(p2)
	}
	if p2 := p.Data2; p2 != nil {
		alloc.Free(p2)
	}
	alloc.Free(p)
	return 0
}

func sub_43E940(a1 unsafe.Pointer) int {
	ail.Startup()
	audioTimer93944 = ail.RegisterTimer(func(u uint32) {
		sub_486EF0()
		legacy.Sub_43D2D0()
		legacy.Sub_486620((*legacy.TimerGroup)(legacy.Get_dword_587000_127004()))
	})
	if audioTimer93944 == math.MaxUint32 {
		return -2147221504 // 0x80040000
	}
	audioTimer93944.SetFrequency(30)
	audioTimer93944.Start()
	legacy.Sub_42EBB0(1, legacy.Get_sub_43E910(), 0, "Audio")
	legacy.Sub_42EBB0(2, legacy.Get_sub_43E8E0(), 0, "Audio")
	*(*uint32)(unsafe.Add(a1, 20)) = 1
	return 0
}

func sub_43E9F0() {
	if audioTimer93944 != math.MaxUint32 {
		audioTimer93944.Stop()
		audioTimer93944.Release()
		audioTimer93944 = math.MaxUint32
	}
	ail.Shutdown()
}

func noxAudioServeT(a1 int) {
	noxAudioServe()
}

func noxAudioServe() { // sub_43F1A0
	if audioDev != 0 {
		ail.Serve()
	}
}

func sub_43EA20(a1 unsafe.Pointer) int {
	//char v2[16]; // [esp+4h] [ebp-10h]

	v2 := sub_43EA90(unsafe.Add(a1, 60))
	audioDev = sub_43EAD0(v2)
	if audioDev == 0 {
		return -2147221504 // 0x80040000
	}
	*memmap.PtrInt32(0x587000, 93948) = 0
	*(*uint32)(unsafe.Add(a1, 196)) = 24
	return 0
}

type audioStruct1 struct {
	Field0  uint16
	Field2  uint16
	Field4  uint32
	Field8  uint32
	Field12 uint16
	Field14 uint16
}

func sub_43EA90(a2 unsafe.Pointer) *audioStruct1 {
	return &audioStruct1{
		Field0:  1,
		Field2:  *(*uint16)(unsafe.Add(a2, 12)),
		Field4:  *(*uint32)(unsafe.Add(a2, 8)),
		Field8:  *(*uint32)(unsafe.Add(a2, 20)),
		Field12: *(*uint16)(unsafe.Add(a2, 16)) * *(*uint16)(unsafe.Add(a2, 12)),
		Field14: 8 * *(*uint16)(unsafe.Add(a2, 16)),
	}
}

func sub_43EAD0(a1 *audioStruct1) ail.Driver {
	if dr := ail.WaveOutOpen(); dr != 0 {
		return dr
	}
	_ = ail.LastError()
	return 0
}

func sub_43EC10() int {
	if audioDev != 0 {
		audioDev.Close()
		audioDev = 0
	}
	return 0
}

func sub_43F130() int {
	return int(audioDev)
}

func sub_43ED00(a1p unsafe.Pointer) int {
	a1 := unsafe.Slice((*uint32)(a1p), 73)
	s := *(**legacy.AudioSample)(unsafe.Pointer(&a1[68]))
	smp := s.Smp
	smp.Init()
	s.Field4 = *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(uintptr(a1[72])), 20))
	if s.Field4 == nil {
		s.Field4 = unsafe.Pointer(uintptr(a1[33]) + 60)
	}
	v3 := int32(legacy.Sub_43F0E0(s.Field4))
	smp.SetType(v3, 0)
	if *(*uint32)(unsafe.Add(s.Field4, 4)) == 2 {
		smp.SetADPCMBlockSize(*(*uint32)(unsafe.Add(s.Field4, 24)))
	}
	sub_43F060(unsafe.Pointer(&a1[0]))
	s.Flag7 = 0
	s.Field3 = 0
	smp.RegisterEOBCallback(func() {
		v := smp.UserData().(*legacy.AudioSample)
		legacy.Sub_43EE00(v)
	})
	smp.RegisterEOSCallback(func() {
		legacy.Sub_43EDB0(smp)
	})
	legacy.Sub_43EE00(s)
	return 0
}

func sub_43D650() {
	if s := legacy.Get_dword_5d4594_816364(); s != 0 {
		s.Close()
		legacy.Set_dword_5d4594_816364(0)
	}
	legacy.Set_dword_5d4594_816092(0)
}

func sub_43D680() {
	if s := legacy.Get_dword_5d4594_816364(); s != 0 {
		s.Pause(true)
	}
}

func sub_43D6A0() {
	if s := legacy.Get_dword_5d4594_816364(); s != 0 {
		s.Pause(false)
	}
}

func sub_44D640() {
	if s := legacy.Get_dword_5d4594_831088(); s != 0 {
		s.Close()
		legacy.Set_dword_5d4594_831088(0)
	}
}

func sub_44D7E0(a1 int) int {
	s := legacy.Get_dword_5d4594_831088()
	if s == 0 {
		return 0
	}
	legacy.Sub_44D5C0(s, a1)
	s.Start()
	return 1
}

var audioTable = []string{
	"",
	"chap1", "chap2wiz", "chap2con", "chap2war", "chap3", "chap4", "chap5", "chap6", "chap7", "chap8",
	"chap9", "chapa", "chapb", "title", "town1", "town2", "town3", "sub1", "sub2", "sub3",
	"wander1", "wander2", "wander3", "credits", "shell", "action1", "action2", "action3", "wander4",
}

func nox_xxx_musicStartPlay_43D6C0(a1p unsafe.Pointer) int {
	if legacy.Get_dword_5d4594_816376() == 0 {
		return 0
	}
	a1 := unsafe.Slice((*uint32)(unsafe.Pointer(a1p)), 4)

	ind := int(a1[0])
	if ind <= 0 || ind >= len(audioTable) {
		return 0
	}
	aname := audioTable[ind]
	sub_43D650()
	*memmap.PtrUint32(0x587000, 93160) = 0
	path := filepath.Join("music", aname)
	if !strings.Contains(path, ".") {
		path += ".wav"
	}
	s := legacy.Get_dword_5d4594_816376().OpenStream(path, 204800)
	if s == 0 {
		if legacy.Get_dword_587000_122856() != 0 && sub_44D930() {
			return 0
		}
		v5 := legacy.Sub_413890()
		if v5 == "" {
			return 0
		}
		*memmap.PtrUint32(0x587000, 93160) = 1
		path2 := filepath.Join(v5, path)
		s = legacy.Get_dword_5d4594_816376().OpenStream(path2, 204800)
		if s == 0 {
			return 0
		}
	}
	s.SetPosition(int(a1[2]))
	legacy.Sub_486320(memmap.PtrOff(0x5D4594, 816148), 0)
	legacy.Sub_486350(memmap.PtrOff(0x5D4594, 816148), 0x4000)
	legacy.Sub_43D3C0(s, int(a1[1]))
	s.Start()
	legacy.Set_dword_5d4594_816092(ind)
	*memmap.PtrUint32(0x5D4594, 816096) = a1[1]
	*memmap.PtrUint32(0x5D4594, 816100) = a1[2]
	*memmap.PtrUint32(0x5D4594, 816104) = a1[3]
	a1[2] = 0
	legacy.Set_dword_5d4594_816364(s)
	return 1
}

func sub_44D930() bool {
	if legacy.Get_dword_587000_122848() == 0 {
		return false
	}
	return legacy.Get_dword_5d4594_830872() != 0 || legacy.Get_dword_5d4594_831088() != 0
}

func sub_44D660(name string) bool {
	sub_44D640()
	path := filepath.Join("dialog", name)
	legacy.Set_dword_587000_122856(0)
	if !strings.Contains(path, ".") {
		path += ".wav"
	}
	s := legacy.Get_dword_5d4594_831092().OpenStream(path, 51200)
	legacy.Set_dword_5d4594_831088(s)
	if s != 0 {
		return true
	}
	v4 := legacy.Sub_413890()
	if v4 == "" {
		return s != 0
	}
	legacy.Set_dword_587000_122856(1)
	path2 := filepath.Join(v4, path)
	s = legacy.Get_dword_5d4594_831092().OpenStream(path2, 51200)
	legacy.Set_dword_5d4594_831088(s)
	return s != 0
}

func sub_43F060(a1p unsafe.Pointer) int {
	a1 := unsafe.Slice((*uint32)(unsafe.Pointer(a1p)), 69)
	s := *(**legacy.AudioSample)(unsafe.Pointer(&a1[68]))
	smp := s.Smp
	smp.SetVolume(int((127 * (a1[45] >> 16)) >> 14))
	smp.SetPan(int((127 * (a1[61] >> 16)) >> 14))
	p2 := s.Field4
	v2 := legacy.Sub_486640(unsafe.Pointer(&a1[44]), int(*(*uint32)(unsafe.Add(p2, 8))))
	smp.SetPlaybackRate(v2)
	return 0
}

func nox_xxx_parseSoundSetBin_424170(path string) error {
	if legacy.Nox_xxx_parseSoundSetBin_424170(path) == 0 {
		return fmt.Errorf("failed to load sounds")
	}
	return nil
}

func nox_audio_initall(a3 int) int {
	if nox_enable_audio == 0 {
		return 1
	}
	if a3 != 0 {
		sub_486F30()
		if sub_4311F0() != 0 {
			legacy.Set_dword_587000_81128(unsafe.Add(legacy.Get_dword_5d4594_805984(), 88))
			dword_5d4594_805980 = sub_4866F0("audio", "audio")
		}
	}
	legacy.Sub_4864A0(memmap.PtrOff(0x5D4594, 805884))
	legacy.Sub_4864A0(legacy.Get_dword_587000_93164())
	legacy.Sub_4864A0(legacy.Get_dword_587000_122852())
	legacy.Sub_4864A0(legacy.Get_dword_587000_127004())
	legacy.Nox_xxx_WorkerHurt_44D810()
	legacy.Sub_43D8E0()
	legacy.Sub_451850(legacy.Get_dword_5d4594_805984(), unsafe.Pointer(dword_5d4594_805980))
	v1 := configGetVolume(VolumeMusic)
	if v1 == 0 {
		legacy.Sub_43DC00()
	}
	legacy.Sub_486320(legacy.Get_dword_587000_93164(), v1)
	v2 := configGetVolume(VolumeDialog)
	if v2 == 0 {
		legacy.Sub_44D960()
	}
	legacy.Sub_486320(legacy.Get_dword_587000_122852(), v2)
	v3 := configGetVolume(VolumeFX)
	if v3 == 0 {
		legacy.Sub_453050()
	}
	legacy.Sub_486320(legacy.Get_dword_587000_127004(), v3)
	return 1
}

func sub_4311F0() int {
	legacy.Sub_486FA0(int(memmap.Int32(0x587000, 94032)))
	v2, free := alloc.Make([]int32{}, 7)
	defer free()
	v2[2] = 22050
	v2[1] = 0
	v2[3] = 2
	v2[4] = 2
	v2[0] = 4
	legacy.Sub_487D00(unsafe.Pointer(&v2[0]))
	v0 := legacy.Sub_487150(-1, unsafe.Pointer(&v2[0]))
	legacy.Set_dword_5d4594_805984(v0)
	return bool2int(v0 != nil && legacy.Sub_487790(v0, 16) == 16)
}

func sub_486F30() int {
	nox_common_list_clear_425760(legacy.Get_dword_587000_155144())
	nox_common_list_clear_425760(unsafe.Add(legacy.Get_dword_587000_155144(), 12))
	*(*uint32)(unsafe.Add(legacy.Get_dword_587000_155144(), 24)) = 0
	*memmap.PtrPtr(0x5D4594, 1193340) = unsafe.Add(legacy.Get_dword_587000_155144(), 32)
	legacy.Sub_4864A0(unsafe.Add(legacy.Get_dword_587000_155144(), 32))
	dword_5d4594_1193336 = 1
	return 0
}

func sub_486EF0() {
	if dword_5d4594_1193336 != 0 {
		if *(*uint32)(unsafe.Add(legacy.Get_dword_587000_155144(), 24)) == 0 {
			v1 := *(*unsafe.Pointer)(unsafe.Add(legacy.Get_dword_587000_155144(), 12))
			for it := unsafe.Add(legacy.Get_dword_587000_155144(), 12); v1 != it; v1 = *(*unsafe.Pointer)(v1) {
				if (*(*int32)(unsafe.Add(v1, 4*3)) & 2) == 0 {
					ccall.CallVoidPtr(*(*unsafe.Pointer)(unsafe.Add(v1, 4*54)), v1)
				}
			}
		}
	}
}

var _ = [1]struct{}{}[288-unsafe.Sizeof(audioStructXxx{})]

type audioStructXxx struct {
	arr0       *audioStructYyy
	size4      uint
	path2_8    [260]byte
	bagfile268 *legacy.FILE // 67, 268
	field272   *legacy.FILE // 68, 272
	field276   uint32       // 69, 276
	field280   uint32       // 70, 280
	field284   uint32       // 71, 284
}

func (p *audioStructXxx) C() unsafe.Pointer {
	return unsafe.Pointer(p)
}

func (p *audioStructXxx) Free() {
	if p == nil {
		return
	}
	if p.bagfile268 != nil {
		legacy.Nox_fs_close(p.bagfile268)
		p.bagfile268 = nil
	}
	if p.field272 != nil {
		legacy.Nox_fs_close(p.field272)
		p.field272 = nil
	}
	if p.arr0 != nil {
		alloc.Free(p.arr0)
		p.arr0 = nil
	}
	alloc.Free(p)
}

var _ = [1]struct{}{}[36-unsafe.Sizeof(audioStructYyy{})]

type audioStructYyy struct {
	field0  [32]byte
	field32 uint32
}

func (p *audioStructYyy) C() unsafe.Pointer {
	return unsafe.Pointer(p)
}

func sub_4866F0(path1 string, path2 string) *audioStructXxx {
	idxPath := path1
	if i := strings.LastIndexByte(idxPath, '.'); i >= 0 {
		idxPath = idxPath[:i]
	}
	idxPath += ".idx"
	fi, err := ifs.Open(idxPath)
	if err != nil {
		return nil
	}
	bf1 := binfile.NewFile(fi)
	cf1 := legacy.NewFileHandle(bf1)
	defer legacy.Nox_fs_close(cf1)

	bagPath := path1
	if i := strings.LastIndexByte(bagPath, '.'); i >= 0 {
		bagPath = bagPath[:i]
	}
	bagPath += ".bag"
	fb, err := ifs.Open(bagPath)
	if err != nil {
		return nil
	}

	p, _ := alloc.New(audioStructXxx{})
	p.bagfile268 = legacy.NewFileHandle(binfile.NewFile(fb))
	if p.bagfile268 == nil {
		p.Free()
		return nil
	}
	var hdr [12]byte
	if _, err := bf1.Read(hdr[:12]); err != nil {
		p.Free()
		return nil
	}
	_ = binary.LittleEndian.Uint32(hdr[0:])
	vers := binary.LittleEndian.Uint32(hdr[4:])
	p.size4 = uint(binary.LittleEndian.Uint32(hdr[8:]))
	var arr []audioStructYyy
	if p.size4 > 0 {
		arr, _ = alloc.Make([]audioStructYyy{}, p.size4)
		p.arr0 = &arr[0]
		var buf [36]byte
		if vers != 1 {
			for i := uint(0); i < p.size4; i++ {
				if _, err := bf1.Read(buf[:36]); err != nil {
					p.Free()
					return nil
				}
				copy(arr[i].field0[:], buf[:32])
				arr[i].field32 = binary.LittleEndian.Uint32(buf[32:])
			}
		} else {
			for i := uint(0); i < p.size4; i++ {
				if _, err := bf1.Read(buf[:32]); err != nil {
					p.Free()
					return nil
				}
				copy(arr[i].field0[:], buf[:32])
				arr[i].field32 = 0
			}
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		s1, s2 := alloc.GoStringS(arr[i].field0[:]), alloc.GoStringS(arr[j].field0[:])
		s1, s2 = strings.ToLower(s1), strings.ToLower(s2)
		return s1 < s2
	})
	p.field276 = 0
	if path2 != "" {
		alloc.StrCopyZero(p.path2_8[:], path2)
		if i := alloc.StrLenS(p.path2_8[:]); p.path2_8[i] == '\\' {
			p.path2_8[i] = 0
		}
		var find legacy.WIN32_FIND_DATAA

		if h := legacy.FindFirstFileA(&p.path2_8[0], &find); int(h) != -1 {
			if find.FileAttributes&0x10 != 0 {
				p.field276 = 1
			} else {
				for legacy.FindNextFileA(h, &find) != 0 {
					if find.FileAttributes&0x10 != 0 {
						p.field276 = 1
						break
					}
				}
			}
			legacy.FindClose(h)
		}
		// TODO: strlen()-1 ?
		if i := alloc.StrLenS(p.path2_8[:]); p.path2_8[i] != '\\' {
			p.path2_8[i] = '\\'
			p.path2_8[i+1] = 0
		}
	}
	return p
}
