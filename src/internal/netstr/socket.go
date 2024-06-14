package netstr

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/netip"
	"syscall"
)

var DebugSockets = false

func GetAddr(addr net.Addr) netip.AddrPort {
	switch a := addr.(type) {
	case nil:
	case interface{ AddrPort() netip.AddrPort }:
		return a.AddrPort()
	case *net.TCPAddr:
		return a.AddrPort()
	case *net.UDPAddr:
		return a.AddrPort()
	default:
		Log.Printf("unsupported address type: %T", a)
	}
	return netip.AddrPort{}
}

func listenOnFreePort(log *log.Logger, port int) (net.PacketConn, int, error) {
	if port < 1024 || port > 0x10000 {
		return nil, 0, NewConnectErr(-15, errors.New("invalid port"))
	}
	for {
		pc, err := listen(log, netip.AddrPortFrom(netip.IPv4Unspecified(), uint16(port)))
		if err == nil {
			return pc, port, nil
		} else if !ErrIsInUse(err) {
			return nil, 0, NewConnectErr(-1, err)
		}
		port++
	}
}

func Listen(addr netip.AddrPort) (net.PacketConn, error) {
	return listen(Log, addr)
}

func listen(log *log.Logger, addr netip.AddrPort) (net.PacketConn, error) {
	log.Printf("listen udp %s", addr)
	l, err := net.ListenUDP("udp4", net.UDPAddrFromAddrPort(addr))
	if err != nil {
		return nil, err
	}
	return l, nil
}

func readFrom(debug bool, log *log.Logger, pc net.PacketConn) ([]byte, netip.AddrPort, error) {
	const buflen = 65535
	buf := make([]byte, buflen)

	n, src, err := pc.ReadFrom(buf)
	buf = buf[:n]
	ap := GetAddr(src)
	if err != nil {
		Log.Println(err)
		return buf, ap, err
	}
	if debug {
		log.Printf("recv %s -> %s [%d]\n%x", ap, pc.LocalAddr(), n, buf)
	}
	return buf, ap, nil
}

func WriteTo(pc net.PacketConn, buf []byte, addr netip.AddrPort) (int, error) {
	return writeTo(DebugSockets, Log, pc, buf, addr)
}

func writeTo(debug bool, log *log.Logger, pc net.PacketConn, buf []byte, addr netip.AddrPort) (int, error) {
	if pc == nil {
		return 0, errors.New("nil conn")
	}
	if debug {
		log.Printf("send %s -> %s [%d]\n%x", pc.LocalAddr(), addr, len(buf), buf)
	}
	n, err := pc.WriteTo(buf, net.UDPAddrFromAddrPort(addr))
	if err != nil {
		Log.Println(err)
		return 0, err
	}
	return n, nil
}

func canReadConn(debug bool, log *log.Logger, pc net.PacketConn) (int, error) {
	if pc == nil {
		return 0, errors.New("nil conn")
	}
	sc, ok := pc.(syscall.Conn)
	if !ok {
		panic(fmt.Errorf("unexpected type: %T", pc))
	}
	rc, err := sc.SyscallConn()
	if err != nil {
		panic(fmt.Errorf("unexpected type: %T: %w", pc, err))
	}
	var (
		ierr syscall.Errno
		cnt  uint32
	)
	err = rc.Control(func(fd uintptr) {
		cnt, ierr = netCanRead(fd)
	})
	if err != nil {
		return 0, err
	}
	if ierr == 0 {
		return int(cnt), nil
	}
	if debug {
		log.Printf("can read: %d", cnt)
	}
	return int(cnt), ierr
}
