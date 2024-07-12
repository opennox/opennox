package dialog

import (
	"testing"
	"time"

	"github.com/noxworld-dev/opennox-lib/platform"
	"github.com/noxworld-dev/opennox-lib/strman"
	"github.com/shoenig/test/must"

	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

func platformTicks() uint64 {
	return uint64(platform.Ticks() / time.Millisecond)
}

func newTestDialog(t *testing.T) *Dialog {
	timer.PlatformTicks = platformTicks

	handles.Init()
	t.Cleanup(func() {
		handles.Release()
	})
	// Initialize audio.
	ail.Startup()
	dr := ail.WaveOutOpen()

	sub_43DBD0 := func() {}
	sub_43DBE0 := func() {}
	sub_43DC40 := func() int { return 0 }
	set_dword_5d4594_831080 := func(v uint32) {
	}
	get_dword_587000_93160 := func() uint32 {
		return 0
	}

	var (
		dword_5d4594_830864 uint32
		dword_5d4594_831076 uint32
		dword_5d4594_831084 uint32
		dword_5d4594_831092 ail.Driver
		dword_587000_122856 uint32
		dword_587000_122848 uint32
		dword_5d4594_830860 uint32
	)

	sm := strman.New()
	err := sm.ReadJSON("testdata/dialogs.json")
	must.NoError(t, err)
	return NewDialog(
		"testdata",
		&dword_587000_122848,
		&dword_5d4594_830864,
		&dword_5d4594_831084,
		&dword_587000_122856,
		&dword_5d4594_831076,
		&dword_5d4594_831092,
		&dword_5d4594_830860,
		func() *strman.StringManager {
			return sm
		},
		new(timer.TimerGroup),
		new(timer.TimerGroup),
		new(timer.TimerGroup),
		new(timer.TimerGroup),
		nil,

		func() ail.Driver {
			return dr
		},
		sub_43DBD0,
		sub_43DBE0,
		sub_43DC40,

		set_dword_5d4594_831080,
		get_dword_587000_93160,
	)
}

func TestDialogIntegration(t *testing.T) {
	d := newTestDialog(t)

	must.EqOp(t, false, d.IsInitialized())
	must.EqOp(t, 0, d.State())
	must.EqOp(t, "", d.FileToRead())
	must.EqOp(t, "", d.CurrentPlayingFile())

	d.Nox_xxx_WorkerHurt_44D810() // Part of init, one file being load with volume 0
	must.EqOp(t, 0, d.State())
	must.EqOp(t, "empty", d.FileToRead())
	must.EqOp(t, "", d.CurrentPlayingFile())
	must.EqOp(t, true, d.IsInitialized())

	d.Sub_44D3A0()
	must.EqOp(t, 2, d.State())
	d.Sub_44D3A0()
	must.EqOp(t, 3, d.State())
	must.EqOp(t, "empty", d.FileToRead())
	must.EqOp(t, "empty", d.CurrentPlayingFile())

	must.EqOp(t, 4, d.GetStream().Status())
	must.EqOp(t, 0, d.GetStream().Position())
}
