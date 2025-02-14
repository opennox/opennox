package server

import (
	"time"

	ns4 "github.com/opennox/noxscript/ns/v4"
)

func (s *Server) NoxScriptNS() NoxScriptNS {
	return NoxScriptNS{s}
}

type NoxScriptNS struct {
	s *Server
}

func (s NoxScriptNS) Frame() int {
	return int(s.s.Frame())
}

func (s NoxScriptNS) FrameRate() int {
	return int(s.s.TickRate())
}

func (s NoxScriptNS) Time() time.Duration {
	return s.s.FrameTS()
}

func (s NoxScriptNS) RandomFloat(min float32, max float32) float32 {
	return float32(s.s.Rand.Logic.FloatClamp(float64(min), float64(max)))
}

func (s NoxScriptNS) Random(min int, max int) int {
	return s.s.Rand.Logic.IntClamp(min, max)
}

func (s NoxScriptNS) StopScript(val any) {
	panic(val)
}

func (s NoxScriptNS) TimerByHandle(h ns4.TimerHandle) ns4.Timer {
	if h == nil {
		return nil
	}
	id := h.TimerScriptID()
	if id <= 0 {
		return nil
	}
	return nsTimer{s: s.s, id: uint32(id)}
}

func (s NoxScriptNS) NewTimer(dt ns4.Duration, fnc ns4.Func, args ...any) ns4.Timer {
	if dt.IsInfinite() {
		panic("trying to create an infinite timer")
	}
	var arg uint32
	switch len(args) {
	default:
		panic("more than one timer arguments are not supported yet")
	case 0:
		arg = 0
	case 1:
		arg = s.s.NoxScriptVM.AsValue(args[0])
	}
	id := s.s.NewTimer(s.s.AsFrames(dt), s.s.NoxScriptVM.AsFuncIndex("Timer", fnc), arg)
	return nsTimer{s: s.s, id: id}
}

type nsTimer struct {
	s  *Server
	id uint32
}

func (t nsTimer) ScriptID() int {
	return int(t.id)
}

func (t nsTimer) TimerScriptID() int {
	return int(t.id)
}

func (t nsTimer) Cancel() bool {
	return t.s.Activators.Cancel(t.id)
}
