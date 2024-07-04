package dialog

import (
	"github.com/noxworld-dev/opennox-lib/strman"
	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

var (
	Sub_43F130       func() int                   // get audio device (as int type)
	Sub_43DBD0       func()                       // looks like music? pair with 43DBE0
	Sub_43DBE0       func()                       // looks like music? pair with 43DBD0
	Sub_43DC40       func() int                   // unknown
	Sub_413890       func() string                // fallback audio file name provider
	GetStringManager func() *strman.StringManager // provides string manager.

	Get_counter_5d4594_830876 func() *timer.TimerGroup
	Get_counter_5d4594_830980 func() *timer.TimerGroup
	Set_dword_5d4594_831080   func(uint32)
	Get_dword_587000_93160    func() uint32
	Get_dword_5d4594_830860   func() uint32
	Set_dword_5d4594_830860   func(uint32)

	GoString                      func(*byte) string
	InternCStr                    func(s string) *byte
	Instance                      *Dialog
	Get_counter_587000_81128_ptr  func() *timer.TimerGroup
	Get_counter_587000_122852_ptr func() *timer.TimerGroup
)

func NewDialog(isEnabled *uint32, volume *uint32, fileToRead **byte, currentPlayingFile **byte, currentStream *ail.Stream, state *uint32, someFlag *uint32, isFallbackMode *uint32, isInitialized *uint32, audioDriver *ail.Driver) *Dialog {
	return &Dialog{
		isEnabled:          isEnabled,
		volume:             volume,
		fileToRead:         fileToRead,
		currentPlayingFile: currentPlayingFile,
		currentStream:      currentStream,
		state:              state,
		someFlag:           someFlag,
		isFallbackMode:     isFallbackMode,
		isInitialized:      isInitialized,
		audioDriver:        audioDriver,
	}
}
