//go:build windows && guiapp && !server

package opennox

import (
	stdlog "log"
	"os"
	"syscall"

	"github.com/opennox/libs/log"
)

func init() {
	// attach stdout/stderr to the console of the parent process
	// this way we can keep building the main binary as GUI App,
	// but still allow seeing console if started via cmd
	const ATTACH_PARENT_PROCESS = ^uint32(0) // (uint32_t)-1
	modkernel32 := syscall.NewLazyDLL("kernel32.dll")
	procAttachConsole := modkernel32.NewProc("AttachConsole")

	r1, _, errno := syscall.Syscall(procAttachConsole.Addr(), 1, uintptr(ATTACH_PARENT_PROCESS), 0, 0)
	if errno != 0 || r1 == 0 {
		return
	}
	hout, err := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		return
	}
	herr, err := syscall.GetStdHandle(syscall.STD_ERROR_HANDLE)
	if err != nil {
		return
	}
	stdout := os.NewFile(uintptr(hout), "/dev/stdout")
	stderr := os.NewFile(uintptr(herr), "/dev/stderr")
	os.Stdout = stdout
	os.Stderr = stderr
	stdlog.SetOutput(stderr)
	log.AddHandler(log.NewTextHandler(stderr))
}
