package test_dialog

import (
	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/dialog"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

var (
	dword_587000_81128             = &timer.TimerGroup{}
	dword_587000_122852            = &timer.TimerGroup{}
	dword_5d4594_830864 uint32     = 0
	dword_5d4594_830868 uint32     = 0
	dword_5d4594_830872 *byte      = nil
	dword_5d4594_830972 *byte      = nil
	dword_5d4594_831076 uint32     = 0
	dword_5d4594_831084 uint32     = 0
	dword_5d4594_831088 ail.Stream = 0
	dword_5d4594_831092 ail.Driver = 0
	dword_587000_122856 uint32     = 0
	dword_587000_122848 uint32     = 0
)

func Get_string_5d4594_830872() string {
	return GoString(dword_5d4594_830872)
}

func Get_string_5d4594_830972() string {
	return GoString(dword_5d4594_830972)
}

func GoString(p *byte) string {
	return alloc.GoString(p)
}
func InternCStr(p string) *byte {
	return alloc.InternCString(p)
}
func CreateInstance() *dialog.Dialog {
	return dialog.NewDialog(
		&dword_587000_122848,
		&dword_5d4594_830868,
		&dword_5d4594_830872,
		&dword_5d4594_830972,
		&dword_5d4594_831088,
		&dword_5d4594_830864,
		&dword_5d4594_831084,
		&dword_587000_122856,
		&dword_5d4594_831076,
		&dword_5d4594_831092)
}

func Get_counter_587000_81128_ptr() *timer.TimerGroup {
	return dword_587000_81128
}

func Get_counter_587000_122852_ptr() *timer.TimerGroup {
	return dword_587000_122852
}

func GetIsInitialized() uint32 {
	return dword_5d4594_831076
}

func GetState() uint32 {
	return dword_5d4594_830864
}

func GetAudioStream() ail.Stream {
	return dword_5d4594_831088
}
