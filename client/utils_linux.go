package client

import (
	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil"
	"golang.org/x/tools/container/intsets"
)

func getProcessPorts(pid int) []int {
	var connections, _ = net.ConnectionsPid("all", int32(pid))
	var result = make([]int, len(connections))

	for i, c := range connections {
		result[i] = int(c.Laddr.Port)
	}

	return result

}
