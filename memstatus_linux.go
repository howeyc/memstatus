// Copyright 2013 Chris Howey. All rights reserved.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
//
// +build linux

package memstatus

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func osMemStatus() (Physical, Virtual MemoryStats) {
	contents, err := ioutil.ReadFile("/proc/meminfo")

	if err != nil {
		return
	}

	reader := bufio.NewReader(bytes.NewBuffer(contents))

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fields := strings.Split(string(line), ":")
		fieldName := fields[0]

		fieldNumber := strings.TrimLeft(fields[1], " ")
		nums := strings.Split(fieldNumber, " ")
		val, numerr := strconv.ParseUint(nums[0], 10, 64)
		if numerr == nil && len(nums) > 1 && nums[1] == "kB" {
			val *= 1024
		}

		switch fieldName {
		case "MemTotal":
			Physical.Total = val
		case "MemFree":
			Physical.Free = val
		case "SwapTotal":
			Virtual.Total = val
		case "SwapFree":
			Virtual.Free = val
		}
	}
	Physical.Used = Physical.Total - Physical.Free
	Virtual.Used = Virtual.Total - Virtual.Free
	return
}
