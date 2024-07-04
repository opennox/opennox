package test_dialog

import (
	"testing"
	"time"

	"github.com/noxworld-dev/opennox-lib/platform"
	"github.com/noxworld-dev/opennox-lib/strman"
	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
	"github.com/noxworld-dev/opennox/v1/legacy/dialog"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
	"github.com/shoenig/test/must"
)

func platformTicks() uint64 {
	return uint64(platform.Ticks() / time.Millisecond)
}

func Setup(t *testing.T) {
	timer.PlatformTicks = platformTicks

	handles.Init()
	t.Cleanup(func() {
		handles.Release()
	})
	// Initialize audio.
	ail.Startup()
	dr := ail.WaveOutOpen()
	dialog.Sub_43F130 = func() int {
		return int(dr)
	}
	dialog.Sub_43DBD0 = func() {}
	dialog.Sub_43DBE0 = func() {}
	dialog.Sub_43DC40 = func() int { return 0 }
	dialog.Sub_413890 = func() string { return "" }
	dialog.GetStringManager = func() *strman.StringManager {
		mgr := &strman.StringManager{}
		mgr.ReadJSON("dialog_test.json")
		return mgr
	}

	counter_830876 := &timer.TimerGroup{}
	dialog.Get_counter_5d4594_830876 = func() *timer.TimerGroup {
		return counter_830876
	}

	counter_830980 := &timer.TimerGroup{}
	dialog.Get_counter_5d4594_830980 = func() *timer.TimerGroup {
		return counter_830980
	}

	dialog.Set_dword_5d4594_831080 = func(v uint32) {
	}

	dialog.Get_dword_587000_93160 = func() uint32 {
		return 0
	}

	dword_5d4594_830860 := uint32(0)
	dialog.Get_dword_5d4594_830860 = func() uint32 {
		return dword_5d4594_830860
	}
	dialog.Set_dword_5d4594_830860 = func(v uint32) {
		dword_5d4594_830860 = v
	}
	dialog.GoString = GoString
	dialog.InternCStr = InternCStr
	dialog.Instance = CreateInstance()
	dialog.Get_counter_587000_81128_ptr = Get_counter_587000_81128_ptr
	dialog.Get_counter_587000_122852_ptr = Get_counter_587000_122852_ptr
}

func TestDialogIntegration(t *testing.T) {
	Setup(t)

	must.Eq(t, 0, GetIsInitialized())
	must.Eq(t, 0, GetState())
	must.Eq(t, "", Get_string_5d4594_830872()) // FileToRead
	must.Eq(t, "", Get_string_5d4594_830972()) // CurrentPlaying

	dialog.Nox_xxx_WorkerHurt_44D810() // Part of init, one file being load with volume 0
	must.Eq(t, 0, GetState())
	must.Eq(t, "empty", Get_string_5d4594_830872()) // FileToRead
	must.Eq(t, "", Get_string_5d4594_830972())      // CurrentPlaying
	must.Eq(t, 1, GetIsInitialized())

	dialog.Sub_44D3A0()
	must.Eq(t, 2, GetState())
	dialog.Sub_44D3A0()
	must.Eq(t, 3, GetState())
	must.Eq(t, "empty", Get_string_5d4594_830872()) // FileToRead
	must.Eq(t, "empty", Get_string_5d4594_830972()) // CurrentPlaying
	must.Eq(t, Get_string_5d4594_830872(), Get_string_5d4594_830972())

	must.Eq(t, 4, GetAudioStream().Status())
	must.Eq(t, 0, GetAudioStream().Position())
}
