// Copyright 2013 Chris Howey. All rights reserved.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package memstatus

type MemoryStats struct {
	Free, Used, Total uint64
}

func MemStatus() (Physical, Virtual MemoryStats) {
	return osMemStatus()
}
