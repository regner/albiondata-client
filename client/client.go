package client

import (
	"log"
	"github.com/regner/albionmarket-client/config"

	"github.com/regner/albionmarket-client/dumper"
)

type Client struct {
}

func NewClient() *Client {
	dumper.UnhandledPacketDumper = *dumper.NewUnhandledPktDumper()
	return &Client{}
}

func (client *Client) Run() {
	log.Print("Starting the Albion Market Client...")

	if config.GlobalConfiguration.Offline {
		proccessOfflinePcap(config.GlobalConfiguration.OfflinePath)
	} else {
		pw := newProcessWatcher()
		pw.run()
	}
}
