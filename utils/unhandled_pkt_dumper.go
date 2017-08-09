package utils

import (
	"container/list"
	"io/ioutil"
	"fmt"
	"strings"
	"os"
)

const dumpFileName string = "pktDump.go"
const fileHeader string = `package main

`
const pktStart string = "//+"
const pktEnd string = "//-"

type UnhandledPktDumper struct {
	dumpedPacketList* list.List
}

func NewUnhandledPktDumper() *UnhandledPktDumper {
	nDumper := UnhandledPktDumper{ dumpedPacketList: list.New()}
	nDumper.loadDumpedPacketOpcodesFromFile()

	return &nDumper
}

func (dumper *UnhandledPktDumper) loadDumpedPacketOpcodesFromFile() {
	// try to open dumpfile
	data, err := ioutil.ReadFile(dumpFileName)
	if err != nil { // we failed, the file did not exist, so create it
		err = ioutil.WriteFile(dumpFileName, []byte(fileHeader), 0644)
		if err != nil { // we failed again, print the error and bail
			fmt.Println(err)
			return
		}

		// file was created, nothing to be done. this is a virgin file
		return
	}

	// get the data as string from our file
	dataStr := string(data)

	// find first pkt
	startPos := strings.Index(dataStr, pktStart)
	// do nothing if no packet is found
	if startPos == -1 {
		return
	}


}
