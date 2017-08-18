package client

import (
	"github.com/drael/GOnetstat"
	"strconv"
)

func getProcessPorts(pid int) []int {
	var foundProcs []int

	for _, el := range GOnetstat.Tcp() {
		if el.Pid == strconv.Itoa(pid) {
			foundProcs = append(foundProcs, int(el.Port))
		}
	}

	for _, el := range GOnetstat.Udp() {
		if el.Pid == strconv.Itoa(pid) {
			foundProcs = append(foundProcs, int(el.Port))
		}
	}

	for _, el := range GOnetstat.Tcp6() {
		if el.Pid == strconv.Itoa(pid) {
			foundProcs = append(foundProcs, int(el.Port))
		}
	}

	for _, el := range GOnetstat.Udp6() {
		if el.Pid == strconv.Itoa(pid) {
			foundProcs = append(foundProcs, int(el.Port))
		}
	}

	return removeDuplicates(foundProcs)
}
