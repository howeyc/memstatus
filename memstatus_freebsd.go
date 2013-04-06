// Copyright 2013 Chris Howey. All rights reserved.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
//
// +build freebsd

package memstatus

/*
#include <sys/types.h>
#include <sys/sysctl.h>
#include <unistd.h>

typedef struct {
	long long total;
	long long free;
} mem_stats;

void fbsd_get_mem_stats(mem_stats* mem_stat){
	int mib[2];
	u_long physmem;
	size_t size;
	u_long free_count;
	int pagesize;

	mib[0] = CTL_HW;
	mib[1] = HW_PHYSMEM;
	size = sizeof physmem;
	if (sysctl(mib, 2, &physmem, &size, NULL, 0) < 0) {
		return;
	}
	mem_stat->total = physmem;

	size = sizeof free_count;
	if (sysctlbyname("vm.stats.vm.v_free_count", &free_count, &size, NULL, 0) < 0){
		return;
	}

	size = sizeof pagesize;
	if (sysctlbyname("hw.pagesize", &pagesize, &size, NULL, 0) < 0){
		return;
	}
	mem_stat->free=(free_count*pagesize);
}
*/
import "C"

func osMemStatus() (Physical, Virtual MemoryStats) {
	var memstat C.mem_stats
	C.fbsd_get_mem_stats(&memstat)

	Physical.Free = uint64(memstat.free)
	Physical.Total = uint64(memstat.total)
	Physical.Used = Physical.Total - Physical.Free

	return
}

func main() {
	var memstat C.mem_stats
	C.fbsd_get_mem_stats(&memstat)
}
