package legacy

/*
#include <stdint.h>
#include "compat.h"
#include "GAME2_2.h"
#include "client/audio/ail/compat_mss.h"
#include "client__io__win95__focus.h"

void sub_43E910(int a1);
void sub_43E8E0(int a1);
void sub_43D2D0();
void sub_43EDB0(HSAMPLE a1);
int sub_43EE00(void* a1);
int sub_43F060(uint32_t* a1);
int sub_43F0E0(uint32_t* a1);
void sub_44D5C0(int a1, int a2);
char* sub_413890();
void sub_43D3C0(int a1, int a2);
int sub_486640(void* a1, int a2);
int nox_xxx_parseSoundSetBin_424170(char* a1);

extern void* dword_587000_127004;
extern uint32_t dword_5d4594_816092;
extern uint32_t dword_5d4594_816364;
extern uint32_t dword_5d4594_816376;
extern uint32_t dword_5d4594_831084;
extern uint32_t dword_5d4594_831092;

void sub_43DC00();
int sub_43D8E0();
void sub_44D960();
void sub_453050();
int sub_451850(int a2, void* a3);
extern void* dword_587000_81128;
extern void* dword_5d4594_805984;
extern void* dword_587000_155144;
*/
import "C"
import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
)

var (
	Sub_43EA20                    func(a1 unsafe.Pointer) int
	Sub_43E9F0                    func()
	Sub_43E940                    func(a1 unsafe.Pointer) int
	Sub_43EFD0                    func(a1 unsafe.Pointer) int
	Sub_43EC10                    func() int
	Sub_43F130                    func() int
	Sub_43ED00                    func(a1p unsafe.Pointer) int
	Sub_43D650                    func()
	Sub_43D680                    func()
	Sub_43D6A0                    func()
	Sub_44D640                    func()
	Sub_44D7E0                    func(a1 int) int
	Nox_xxx_musicStartPlay_43D6C0 func(a1p unsafe.Pointer) int
	Sub_44D660                    func(a1 string) bool
	Sub_43F060                    func(a1p unsafe.Pointer) int
	Sub_43EC30                    func(a1 unsafe.Pointer) int
	Sub_43ECB0                    func(a1 unsafe.Pointer) int
)

type _WIN32_FIND_DATAA = C.struct__WIN32_FIND_DATAA
type WIN32_FIND_DATAA struct {
	FileAttributes     uint32
	ftCreationTime     C.FILETIME
	ftLastAccessTime   C.FILETIME
	ftLastWriteTime    C.FILETIME
	nFileSizeHigh      uint32
	nFileSizeLow       uint32
	dwReserved0        uint32
	dwReserved1        uint32
	cFileName          [C.MAX_PATH]byte
	cAlternateFileName [14]byte
}

func FindFirstFileA(a1 *byte, f *WIN32_FIND_DATAA) uintptr {
	return uintptr(unsafe.Pointer(C.FindFirstFileA((*C.char)(unsafe.Pointer(a1)), (*_WIN32_FIND_DATAA)(unsafe.Pointer(f)))))
}

func FindNextFileA(h uintptr, f *WIN32_FIND_DATAA) int {
	return int(C.FindNextFileA(C.HANDLE(unsafe.Pointer(h)), (*_WIN32_FIND_DATAA)(unsafe.Pointer(f))))
}

func FindClose(h uintptr) {
	C.FindClose(C.HANDLE(unsafe.Pointer(h)))
}

var _ = [1]struct{}{}[32-unsafe.Sizeof(AudioSample{})]

type AudioSample struct {
	Dev    ail.Driver     // 0, 0
	Field1 unsafe.Pointer // 1, 4
	Smp    ail.Sample     // 2, 8
	Field3 uint32         // 3, 12
	Field4 unsafe.Pointer // 4, 16
	Data1  *byte          // 5, 20
	Data2  *byte          // 6, 24
	Flag7  uint32         // 7, 28
}

//export sub_43E8E0
func sub_43E8E0(a1 int) {}

//export sub_43E910
func sub_43E910(a1 int) {}

//export sub_43F050
func sub_43F050() int {
	return 0
}

//export sub_43F0D0
func sub_43F0D0() int {
	return 0
}

//export sub_43F030
func sub_43F030(a1 int) int {
	panic("abort")
}

//export AIL_set_stream_volume
func AIL_set_stream_volume(s C.HSTREAM, volume C.int32_t) {
	ail.Stream(unsafe.Pointer(s)).SetVolume(int(volume))
}

//export AIL_stream_position
func AIL_stream_position(s C.HSTREAM) C.int32_t {
	return C.int32_t(ail.Stream(unsafe.Pointer(s)).Position())
}

//export AIL_stream_status
func AIL_stream_status(s C.HSTREAM) C.int32_t {
	return C.int32_t(ail.Stream(unsafe.Pointer(s)).Status())
}

//export AIL_load_sample_buffer
func AIL_load_sample_buffer(s C.HSAMPLE, num C.uint32_t, buf unsafe.Pointer, sz C.uint32_t) {
	ail.Sample(unsafe.Pointer(s)).LoadBuffer(uint32(num), unsafe.Slice((*byte)(buf), int(sz)))
}

//export AIL_sample_buffer_ready
func AIL_sample_buffer_ready(s C.HSAMPLE) C.int32_t {
	return C.int32_t(ail.Sample(unsafe.Pointer(s)).BufferReady())
}

//export AIL_sample_user_data
func AIL_sample_user_data(s C.HSAMPLE) unsafe.Pointer {
	v := ail.Sample(unsafe.Pointer(s)).UserData()
	if v == nil {
		return nil
	}
	return unsafe.Pointer(v.(*AudioSample))
}

//export sub_43F010
func sub_43F010(a1 unsafe.Pointer) int {
	p := *(**AudioSample)(unsafe.Add(a1, 272))
	p.Smp.Stop()
	return 0
}

//export sub_43EA20
func sub_43EA20(a1 unsafe.Pointer) int {
	return Sub_43EA20(a1)
}

//export sub_43E9F0
func sub_43E9F0() {
	Sub_43E9F0()
}

//export sub_43E940
func sub_43E940(a1 unsafe.Pointer) int {
	return Sub_43E940(a1)
}

//export sub_43EFD0
func sub_43EFD0(a1 unsafe.Pointer) int {
	return Sub_43EFD0(a1)
}

//export sub_43EC10
func sub_43EC10() int {
	return Sub_43EC10()
}

//export sub_43F130
func sub_43F130() int {
	return Sub_43F130()
}

//export sub_43ED00
func sub_43ED00(a1p *C.uint32_t) int {
	return Sub_43ED00(unsafe.Pointer(a1p))
}

//export sub_43D650
func sub_43D650() {
	Sub_43D650()
}

//export sub_43D680
func sub_43D680() {
	Sub_43D680()
}

//export sub_43D6A0
func sub_43D6A0() {
	Sub_43D6A0()
}

//export sub_44D640
func sub_44D640() {
	Sub_44D640()
}

//export sub_44D7E0
func sub_44D7E0(a1 int) int {
	return Sub_44D7E0(a1)
}

//export nox_xxx_musicStartPlay_43D6C0
func nox_xxx_musicStartPlay_43D6C0(a1p *C.int) int {
	return Nox_xxx_musicStartPlay_43D6C0(unsafe.Pointer(a1p))
}

//export sub_44D660
func sub_44D660(a1 *C.char) int {
	return bool2int(Sub_44D660(GoString(a1)))
}

//export sub_43F060
func sub_43F060(a1p *C.uint32_t) int {
	return Sub_43F060(unsafe.Pointer(a1p))
}

//export sub_43EC30
func sub_43EC30(a1p unsafe.Pointer) int {
	return Sub_43EC30(unsafe.Pointer(a1p))
}

//export sub_43ECB0
func sub_43ECB0(a1p unsafe.Pointer) int {
	return Sub_43ECB0(unsafe.Pointer(a1p))
}

func Get_dword_5d4594_831092() ail.Driver {
	return ail.Driver(C.dword_5d4594_831092)
}

func Sub_43D2D0() {
	C.sub_43D2D0()
}

func Sub_486620(v *TimerGroup) {
	C.sub_486620((*C.timerGroup)(unsafe.Pointer(v)))
}

func Sub_42EBB0(a1 int, a2 unsafe.Pointer, a3 int, a4 string) {
	C.sub_42EBB0(C.uint(a1), (*[0]byte)(a2), C.int(a3), internCStr(a4))
}

func Get_sub_43E910() unsafe.Pointer {
	return C.sub_43E910
}

func Get_sub_43E8E0() unsafe.Pointer {
	return C.sub_43E8E0
}

func Get_dword_587000_127004() unsafe.Pointer {
	return C.dword_587000_127004
}

func Sub_43F0E0(v unsafe.Pointer) int {
	return int(C.sub_43F0E0((*C.uint)(v)))
}

func Sub_43EE00(v *AudioSample) {
	C.sub_43EE00(unsafe.Pointer(v))
}

func Sub_43EDB0(v ail.Sample) {
	C.sub_43EDB0((C.HSAMPLE)(unsafe.Pointer(v)))
}

func Get_dword_5d4594_816364() ail.Stream {
	return ail.Stream(C.dword_5d4594_816364)
}

func Set_dword_5d4594_816364(v ail.Stream) {
	C.dword_5d4594_816364 = C.uint(v)
}

func Set_dword_5d4594_816092(v int) {
	C.dword_5d4594_816092 = C.uint(v)
}

func Sub_413890() string {
	return GoString(C.sub_413890())
}

func Nox_xxx_parseSoundSetBin_424170(path string) int {
	return int(C.nox_xxx_parseSoundSetBin_424170(internCStr(path)))
}

func Sub_43D8E0() {
	C.sub_43D8E0()
}

func Sub_43DC00() {
	C.sub_43DC00()
}

func Sub_44D960() {
	C.sub_44D960()
}

func Sub_453050() {
	C.sub_453050()
}

func Sub_44D5C0(s ail.Stream, a2 int) {
	C.sub_44D5C0(C.int(s), C.int(a2))
}

func Get_dword_5d4594_816376() ail.Driver {
	return ail.Driver(C.dword_5d4594_816376)
}

func Sub_486320(p unsafe.Pointer, a2 int) {
	C.sub_486320((*C.timer)(p), C.int(a2))
}

func Sub_486350(p unsafe.Pointer, a2 int) {
	C.sub_486350((*C.timer)(p), C.int(a2))
}

func Sub_43D3C0(s ail.Stream, a2 int) {
	C.sub_43D3C0(C.int(s), C.int(a2))
}

func Sub_486640(a1 unsafe.Pointer, a2 int) int {
	return int(C.sub_486640(a1, C.int(a2)))
}

func Get_dword_5d4594_805984() unsafe.Pointer {
	return C.dword_5d4594_805984
}

func Set_dword_5d4594_805984(v unsafe.Pointer) {
	C.dword_5d4594_805984 = v
}

func Set_dword_587000_81128(v unsafe.Pointer) {
	C.dword_587000_81128 = v
}

func Sub_4864A0(v unsafe.Pointer) {
	C.sub_4864A0((*C.timerGroup)(v))
}

func Sub_451850(a1 unsafe.Pointer, a2 unsafe.Pointer) {
	C.sub_451850(C.int(uintptr(a1)), a2)
}

func Sub_486FA0(a1 int) {
	C.sub_486FA0(C.int(a1))
}

func Sub_487D00(a1 unsafe.Pointer) {
	C.sub_487D00((*C.uint)(a1))
}

func Sub_487150(a1 int, a2 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.sub_487150(C.int(a1), a2))
}

func Sub_487790(a1 unsafe.Pointer, a2 int) int {
	return int(C.sub_487790(C.int(uintptr(a1)), C.int(a2)))
}
