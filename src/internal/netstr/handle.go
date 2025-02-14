package netstr

import (
	"net/netip"

	"github.com/opennox/opennox/v1/common/ntype"
)

func errHandle(e int) handle {
	if e >= 0 {
		panic("must be negative")
	}
	return handle{nil, e}
}

type handle struct {
	g *Streams
	i int // hiding it in a struct helps prevent direct casts
}

func (h handle) Valid() bool {
	return h.g != nil && h.i >= 0 && h.i < maxStructs
}

func (h handle) IsHost() bool {
	return h.i == 0
}

func (h handle) Player() ntype.PlayerInd {
	return ntype.PlayerInd(h.i - 1)
}

func (h handle) Get() *Conn {
	if !h.Valid() {
		return nil
	}
	return h.g.streams[h.i]
}

func (h handle) IP() netip.Addr {
	return h.Get().IP()
}
