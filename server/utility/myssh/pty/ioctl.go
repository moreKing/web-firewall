package pty

import (
	"os"
	"syscall"
)

func ioctl(f *os.File, cmd, ptr uintptr) error {
	return ioctlInner(f.Fd(), cmd, ptr) // Fall back to blocking io.
}

// Local syscall const values.
const (
	TIOCGWINSZ = syscall.TIOCGWINSZ
	TIOCSWINSZ = syscall.TIOCSWINSZ
)

func ioctlInner(fd, cmd, ptr uintptr) error {
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, ptr)
	if e != 0 {
		return e
	}
	return nil
}
