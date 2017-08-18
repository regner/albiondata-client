package client

import (
	"github.com/mitchellh/go-ps"
	"golang.org/x/tools/container/intsets"
)

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
