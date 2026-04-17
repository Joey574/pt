//go:build linux

package pt

import "golang.org/x/sys/unix"

func setAffinity(i int) {
	var cpuset unix.CPUSet
	cpuset.Set(i)
	unix.SchedSetaffinity(0, &cpuset)
}
