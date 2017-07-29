package main

import (
	"flag"
	"log"
	"sync"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"./assemblers"
	"./utils"
)

func main() {
	log.Print("Starting the Albion Market Client...")
	config := utils.ClientConfig{}

	flag.StringVar(&config.IngestUrl, "i", "https://albion-market.com/api/v1/ingest/", "URL to send market data to.")
	flag.BoolVar(&config.DisableUpload, "d", false, "If specified no attempts will be made to upload data to remote server.")
	flag.BoolVar(&config.SaveLocally, "s", false, "If specified all market orders will be saved locally.")
	flag.Parse()

	if config.DisableUpload {
		log.Print("Remote upload of market orders is disabled!")
	} else {
		log.Printf("Using the following ingest: %v", config.IngestUrl)
	}

	if config.SaveLocally {
		log.Print("Saving market orders locally.")
	}

	devices, err := pcap.FindAllDevs()

	if err != nil {
		log.Fatal(err)
	}
	if len(devices) == 0 {
		log.Fatal("Unable to find network device.")
	}

	var wg sync.WaitGroup
	wg.Add(len(devices))

	for _, device := range devices {
		go captureDeviceTraffic(device.Name, config, wg)
	}

	wg.Wait()
}

func captureDeviceTraffic(deviceName string, config utils.ClientConfig, wg sync.WaitGroup) {
	defer wg.Done()

	handle, err := pcap.OpenLive(deviceName, 2048, false, pcap.BlockForever)

	if err != nil {
		log.Fatal(err)
	}

	defer handle.Close()

	var filter = `
		src 192.168.2.15 &&
		dst 158.85.26.40`
	err = handle.SetBPFFilter(filter)

	if err != nil {
		log.Fatal(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	source.NoCopy = true

	assembler := assemblers.NewMarketAssembler(config)

	log.Printf("Starting to process packets for %s...", deviceName)
	for packet := range source.Packets() {
		assembler.ProcessPacket(packet)
	}
}

// (udp.port == 54332 || udp.port == 61621) && ip.src == 192.168.2.15 && ip.dst == 158.85.26.40 && udp.length == 142
// (udp.port == 54332 || udp.port == 61621) && ip.src == 192.168.2.15 && ip.dst == 158.85.26.40 && udp.length == 108
// udp port 61621 && src 192.168.2.15 && dst == 158.85.26.40 && less 142 && greater 142
// (udp.port == 54332 || udp.port == 61621) && ip.src == 192.168.2.15 && ip.dst == 158.85.26.40 && (udp.length == 116 || udp.length == 148 || udp.length == 104)
// pine
// iron
// sandstone
// (udp.port == 54332 || udp.port == 61621) && ip.src == 192.168.2.15 && ip.dst == 158.85.26.40 && (udp.length == 116 || udp.length == 148 || udp.length == 104)
// 34433, 1
