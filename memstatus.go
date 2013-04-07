// Copyright 2013 Chris Howey. All rights reserved.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package memstatus

// Memory info, all values in bytes
type MemoryStats struct {
	Free, Used, Total uint64
}

// Returns memory information in terms of Physical, and Virtual (Swap) memory.
func MemStatus() (Physical, Virtual MemoryStats) {
	return osMemStatus()
}
