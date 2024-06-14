package netstr

import (
	"errors"
	"net"
	"net/netip"

	"github.com/noxworld-dev/opennox-lib/noxnet"
)

type LobbyWaitOptions struct {
	OnResult func(addr netip.AddrPort, data []byte)
	OnCode15 func()
	OnPing   func(addr netip.AddrPort, data []byte)
	OnCode19 func(errcode byte) bool
	OnCode20 func()
	OnCode21 func()
}

func WaitForLobbyResults(conn net.PacketConn, srvAddr netip.Addr, flag RecvFlags, opts LobbyWaitOptions) (int, error) {
	argp := 0
	if flag.Has(RecvCanRead) {
		var err error
		argp, err = canReadConn(DebugSockets, Log, conn)
		if err != nil {
			return 0, err
		} else if argp == 0 {
			// TODO: is it an error at all?
			return 0, errors.New("nothing to read")
		}
	} else {
		argp = 1
	}
	buf := make([]byte, 256)
	for {
		buf = buf[:cap(buf)]
		n, from, err := readFrom(DebugSockets, Log, conn, buf)
		if err != nil {
			return 0, err
		}
		buf = buf[:n]
		op := noxnet.Op(buf[2])
		if op < code32 {
			if op == codeInfo13 || srvAddr == from.Addr() {
				switch op {
				case codeInfo13:
					if from.Addr().IsValid() {
						opts.OnResult(from, buf)
					}
				case codeErr15:
					opts.OnCode15()
				case codePing16:
					opts.OnPing(from, buf)
				case codeErrCode19:
					if !opts.OnCode19(buf[3]) {
						break
					}
				case codeACK20:
					opts.OnCode20()
				case codeErr21:
					opts.OnCode21()
				}
			}
		}
		if flag&RecvCanRead == 0 || (flag&RecvJustOne) != 0 {
			return n, nil
		}
		argp, err = canReadConn(DebugSockets, Log, conn)
		if err != nil {
			return n, err
		} else if argp == 0 {
			return n, nil
		}
	}
}
