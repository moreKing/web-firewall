package pty

import (
	"os"
	"syscall"
	"unsafe"
)

// Winsize describes the terminal size.
type Winsize struct {
	Rows uint16 // ws_row: Number of rows (in cells).
	Cols uint16 // ws_col: Number of columns (in cells).
	X    uint16 // ws_xpixel: Width in pixels.
	Y    uint16 // ws_ypixel: Height in pixels.
}

// Getsize returns the number of rows (lines) and cols (positions
// in each line) in terminal t.
func Getsize(t *os.File) (rows, cols int, err error) {
	ws, err := GetsizeFull(t)
	if err != nil {
		return 0, 0, err
	}
	return int(ws.Rows), int(ws.Cols), nil
}

// Setsize resizes t to s.
func Setsize(t *os.File, ws *Winsize) error {
	//nolint:gosec // Expected unsafe pointer for Syscall call.
	return ioctl(t, syscall.TIOCSWINSZ, uintptr(unsafe.Pointer(ws)))
}

// GetsizeFull returns the full terminal size description.
func GetsizeFull(t *os.File) (size *Winsize, err error) {
	var ws Winsize

	//nolint:gosec // Expected unsafe pointer for Syscall call.
	if err := ioctl(t, syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&ws))); err != nil {
		return nil, err
	}
	return &ws, nil
}

// InheritSize applies the terminal size of pty to tty. This should be run
// in a signal handler for syscall.SIGWINCH to automatically resize the tty when
// the pty receives a window size change notification.
func InheritSize(pty, tty *os.File) error {
	size, err := GetsizeFull(pty)
	if err != nil {
		return err
	}
	return Setsize(tty, size)
}
