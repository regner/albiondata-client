package client

import (
	"github.com/mitchellh/mapstructure"
	"github.com/regner/albionmarket-client/dumper"
	"github.com/regner/albionmarket-client/config"
)

func decode(params map[string]interface{}, dumperParam *dumper.UPDstringParams) operation {
	if _, ok := params["253"]; !ok {
		return nil
	}

	code := params["253"].(int16)

	switch code {
	case 67:
		operation := requestBuyOrders{}
		mapstructure.Decode(params, &operation)

		return operation
	default:
		if config.GlobalConfiguration.DumpUnknown == true {
			dumper.UnhandledPacketDumper.AddPacket(dumperParam)
		}
		return nil
	}
}
