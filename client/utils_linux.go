package client

import (
	"github.com/mitchellh/go-ps"
	"golang.org/x/tools/container/intsets"
	"github.com/shirou/gopsutil"
)

func getProcessPorts(pid int) []int {
	var connections, _ = net.ConnectionsPid("all", int32(pid))
	var result = make([]int, len(connections))

	for i, c := range connections {
		result[i] = int(c.Laddr.Port)
	}

	return result

}
