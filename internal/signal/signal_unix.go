//go:build unix

package isignal

import (
	"os"
	"syscall"
)

var StopSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGHUP,
	syscall.SIGQUIT,
}
