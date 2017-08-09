package client

import (
	"github.com/mitchellh/mapstructure"
	"github.com/regner/albionmarket-client/utils"
)

func decode(params map[string]interface{}, dumperParam *utils.UPDstringParams) operation {
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
		//TODO: Insert start param check to stop it from writing to file.
		utils.UnhandledPacketDumper.AddPacket(dumperParam)
		return nil
	}
}
