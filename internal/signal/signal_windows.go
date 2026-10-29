//go:build windows

package isignal

import (
	"os"
	"syscall"
)

var StopSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGBREAK,
}
