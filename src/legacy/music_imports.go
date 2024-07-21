package legacy

/*
#include <stdint.h>

extern uint32_t dword_5d4594_816356;
extern uint32_t dword_5d4594_816368;
extern uint32_t dword_5d4594_816340;
extern uint32_t dword_5d4594_816372;
extern uint32_t dword_5d4594_816376;
extern uint32_t dword_5d4594_816348;
extern uint32_t dword_587000_93156;
extern uint32_t dword_587000_93160;
extern void* dword_587000_81128;
*/
import "C"

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/music"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

var (
	MusicModule *music.Module
)

func initMusic() {
	var (
		dword_5d4594_816368      *uint32            = (*uint32)(&C.dword_5d4594_816368)
		dword_5d4594_816372      *uint32            = (*uint32)(&C.dword_5d4594_816372)
		dword_5d4594_816376      *ail.Driver        = (*ail.Driver)(unsafe.Pointer(&C.dword_5d4594_816376))
		dword_587000_93156       *uint32            = (*uint32)(&C.dword_587000_93156)
		dword_587000_93160       *uint32            = (*uint32)(&C.dword_587000_93160)
		dword_5d4594_816340      *uint32            = (*uint32)(&C.dword_5d4594_816340)
		counter_5d4594_816244    *timer.TimerGroup  = memmap.PtrT[timer.TimerGroup](0x5d4594, 816244)
		ptr_counter_587000_81128 **timer.TimerGroup = (**timer.TimerGroup)(unsafe.Pointer(&C.dword_587000_81128))
		dword_5d4594_816348      *uint32            = (*uint32)(&C.dword_5d4594_816348)
	)
	MusicModule = music.NewModule(
		"music",
		dword_5d4594_816368,
		dword_5d4594_816372,
		dword_5d4594_816376,
		dword_587000_93156,
		dword_587000_93160,
		dword_5d4594_816340,
		counter_5d4594_816244,
		ptr_counter_587000_81128,
		dword_5d4594_816348,
		Sub_43F130,
		checkDialogs,
		Sub_413890,
		PlatformTicks,
	)
}

func Get_block_816060() *music.MusicState {
	return memmap.PtrT[music.MusicState](0x5D4594, 816060)
}

func PtrUint32_5d4594_816344() *uint32 {
	return memmap.PtrUint32(0x5d4594, 816344)
}

//export sub_43D9E0
func sub_43D9E0(a1p unsafe.Pointer) {
	Sub_43D9E0((*music.MusicState)(a1p))
}

//export sub_43DD10
func sub_43DD10(ret unsafe.Pointer) {
	*(*music.MusicState)(ret) = MusicModule.GetCurrentBlock()
}

// Stop playing music
func Sub_43D990() {
	v1 := music.MusicState{}
	v1.MusicIdx = 0
	Sub_43D9E0(&v1)
}

//export sub_43DBD0
func sub_43DBD0() {
	MusicModule.Sub_43DBD0()
}

//export sub_43DBE0
func sub_43DBE0() {
	MusicModule.Sub_43DBE0()
}

func checkDialogs() bool {
	return Dialogs.IsFallbackMode() != 0 && Dialogs.Sub_44D930()
}

func Sub_43D650() {
	MusicModule.Sub_43D650()
}

func Sub_43D9B0(a1, a2 uint32) {
	blk := music.MusicState{MusicIdx: a1, Volume: a2, Position: 0, D: 0}
	Sub_43D9E0(&blk)
}

//export sub_43D9B0
func sub_43D9B0(a1, a2 int) {
	Sub_43D9B0(uint32(a1), uint32(a2))
}

func Sub_43D8E0() int {
	MusicModule.Init()
	return 1
}

func Sub_43D440() {
	MusicModule.Update()
}

// Update current playing music to a1
func Sub_43D9E0(a1 *music.MusicState) {
	MusicModule.SetNextMusic(a1)
}

func Sub_43D2D0() {
	MusicModule.Sub_43D2D0()
}

func Sub_43D3C0(a1, a2 int) {
	MusicModule.Sub_43D3C0(ail.Stream(a1), uint32(a2))
}

func Sub_43DC40() int {
	return MusicModule.Sub_43DC40()
}

//export sub_43DD70
func sub_43DD70(a1, a2 int) {
	MusicModule.Sub_43DD70(uint32(a1), uint32(a2))
}
