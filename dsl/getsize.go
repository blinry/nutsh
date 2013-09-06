package dsl

import (
	"syscall"
	"unsafe"
)

type window_size struct {
    row    uint16
    col    uint16
    xpixel uint16
    ypixel uint16
}

func getsize() (int, int) {
	w := new(window_size)
	res, _, _ := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		syscall.TIOCGWINSZ,
		uintptr(unsafe.Pointer(w)),
	)
	if int(res) == -1 {
		panic("fuck")
	}	
	return int(w.row), int(w.col)
}
