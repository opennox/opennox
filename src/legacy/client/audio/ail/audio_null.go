//go:build server

package ail

import (
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
)

func (dig Driver) AllocateSample() Sample {
	h := Sample(handles.New())
	return h
}

func (s Sample) Release() {
	handles.AssertValid(uintptr(s))
}

func (dig Driver) OpenStream(name string) Stream {
	h := handles.New()
	return Stream(h)
}

func (h Stream) Close() error {
	handles.AssertValid(uintptr(h))
	return nil
}

func RegisterTimer(f func(u uint32)) Timer {
	h := handles.NewPtr()
	return Timer(h)
}

func (h Timer) Release() {
	handles.AssertValid(uintptr(h))
}

func (h Sample) GetSource() *uint32 {
	return nil
}

func (h Sample) End() {}

func (h Sample) Init() {}

func LastError() string { return "" }

func (h Sample) LoadBuffer(num uint32, buf []byte) {}

func (h Stream) Pause(pause bool) {}

func (h Sample) RegisterEOBCallback(f func()) {}

func (h Sample) RegisterEOSCallback(f func()) {}

func (h Sample) BufferReady() int { return -1 }

func (h Sample) UserData() any { return nil }

func Serve() {}

func (h Sample) SetADPCMBlockSize(block uint32) {}

func (h Sample) SetPan(pan int) {}

func (h Sample) SetPlaybackRate(rate int) {}

func (h Sample) SetType(format int32, flags uint32) {}

func (h Sample) SetUserData(value any) {}

func (h Sample) SetVolume(volume int) {}

func (h Stream) SetPosition(offset int) {}

func (h Stream) SetVolume(volume int) {}

func (h Timer) SetFrequency(hertz uint) {}

func Shutdown() {}

func (h Stream) Start() {}

func (h Timer) Start() {}

func Startup() int32 { return -1 }

func (h Sample) Stop() {}

func (h Timer) Stop() {}

func (h Stream) Position() int { return -1 }

func (h Stream) Status() int { return 2 }

func (dig Driver) Close() error { return nil }

func WaveOutOpen() Driver {
	return 0
}
