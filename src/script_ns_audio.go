package opennox

import (
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/legacy"
)

func (s noxScriptNS) Music(music int, volume int) {
	if noxflags.HasGame(noxflags.GameModeCoop) {
		legacy.Sub_43D9B0(uint32(music), uint32(volume))
	} else {
		s.s.NetMusic(music, volume)
	}
}

func (s noxScriptNS) MusicPushEvent() {
	if noxflags.HasGame(noxflags.GameModeCoop) {
		legacy.Sub_43DA80()
	} else {
		s.s.NetMusicPushEvent()
	}
}

func (s noxScriptNS) MusicPopEvent() {
	if noxflags.HasGame(noxflags.GameModeCoop) {
		legacy.Sub_43DAD0()
	} else {
		s.s.NetMusicPopEvent()
	}
}

func (s noxScriptNS) MusicEvent() {
	if noxflags.HasGame(noxflags.GameModeCoop) {
		legacy.Sub_43D9B0(0, 0)
	} else {
		s.s.NetMusicEvent()
	}
}
