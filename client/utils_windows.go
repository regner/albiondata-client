package client

import (
	"github.com/drael/GOnetstat"
	"golang.org/x/tools/container/intsets"
	"github.com/mitchellh/go-ps"
	"strconv"
)

func findProcess(processName string) []int {
	var results []int
	var processes, _ = ps.Processes()

	for _, proc := range processes {
		if proc.Executable() == processName {
			results = append(results, proc.Pid())
		}
	}

	return results
}

func removeDuplicates(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func getProcessPorts(pid int) []int {
	tcpProcs := GOnetstat.Tcp()
	udpProcs := GOnetstat.Udp()
	tcp6Procs := GOnetstat.Tcp6()
	udp6Procs := GOnetstat.Udp6()

	var foundProcs []int

	for _, el := range tcpProcs {
		if el.Pid == strconv.Itoa(pid) {
			foundProcs = append(foundProcs, int(el.Port))
		}
	}

	for _, el := range udpProcs {
		if el.Pid == strconv.Itoa(pid) {
			foundProcs = append(foundProcs, int(el.Port))
		}
	}

	for _, el := range tcp6Procs {
		if el.Pid == strconv.Itoa(pid) {
			foundProcs = append(foundProcs, int(el.Port))
		}
	}

	for _, el := range udp6Procs {
		if el.Pid == strconv.Itoa(pid) {
			foundProcs = append(foundProcs, int(el.Port))
		}
	}

	var result = removeDuplicates(foundProcs)


	return result
}

func diffIntSets(a []int, b []int) ([]int, []int) {
	var aSparse = intsets.Sparse{}
	var bSparse = intsets.Sparse{}

	for _, k := range a {
		aSparse.Insert(k)
	}

	for _, k := range b {
		bSparse.Insert(k)
	}

	var addedSparse = intsets.Sparse{}
	addedSparse.Difference(&bSparse, &aSparse)
	var addedSlice []int = addedSparse.AppendTo(make([]int, 0))

	var removedSparse = intsets.Sparse{}
	removedSparse.Difference(&aSparse, &bSparse)
	var removedSlice []int = removedSparse.AppendTo(make([]int, 0))

	var added = make([]int, addedSparse.Len())
	var removed = make([]int, removedSparse.Len())

	for i := 0; i < addedSparse.Len(); i++ {
		added[i] = addedSlice[i]
	}

	for i := 0; i < removedSparse.Len(); i++ {
		removed[i] = removedSlice[i]
	}

	return added, removed
}
