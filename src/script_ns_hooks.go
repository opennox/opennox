package opennox

import (
	ns4 "github.com/opennox/noxscript/ns/v4"

	"github.com/opennox/opennox/v1/server"
)

func (s noxScriptNS) OnChat(fnc ns4.ChatFunc) {
	s.s.OnChat(func(t *server.Team, p *server.Player, obj *server.Object, msg string) string {
		var (
			tm ns4.Team
			pl ns4.Player
			no ns4.Obj
		)
		if t != nil {
			tm = nsTeam{s.s, t}
		}
		if p != nil {
			pl = s.toPlayer(p)
		}
		if obj != nil {
			no = s.toObj(obj)
		}
		return fnc(tm, pl, no, msg)
	})
}
