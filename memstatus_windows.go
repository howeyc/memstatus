// Copyright 2013 Chris Howey. All rights reserved.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
//
// +build windows

package memstatus

import (
	"syscall"
	"unsafe"
)

type memoryStatusEx struct {
	dwLength                uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

func osMemStatus() (Physical, Virtual MemoryStats) {
	modkernel32 := syscall.NewLazyDLL("kernel32.dll")
	GlobalMemoryStatusEx := modkernel32.NewProc("GlobalMemoryStatusEx")

	var stats memoryStatusEx
	stats.dwLength = uint32(unsafe.Sizeof(stats))
	pStats := &stats
	GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(pStats)))

	Physical.Free = stats.ullAvailPhys
	Physical.Total = stats.ullTotalPhys
	Physical.Used = Physical.Total - Physical.Free

	Virtual.Free = stats.ullAvailVirtual
	Virtual.Total = stats.ullTotalVirtual
	Virtual.Used = Virtual.Total - Virtual.Free
	return
}
