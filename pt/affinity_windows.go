//go:build windows

package pt

import (
	"golang.org/x/sys/windows"
)

func setAffinity(i int) {
	mask := uintptr(1 << i)
	thread, _ := windows.GetCurrentThread()
	windows.SetThreadAffinityMask(thread, mask)
}
