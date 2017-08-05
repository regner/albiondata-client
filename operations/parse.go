package operations

import (
	"log"

	"github.com/mitchellh/mapstructure"
)

func DecodeRequest(params map[string]interface {}) interface{} {
	code, ok := params["253"]

	if !ok {
		return nil
	}

	switch code.(int16) {
		case 67:
			var operation = RequestBuyOrders{}
			mapstructure.Decode(params, &operation)
			return operation
	}

	log.Print("Unhandled Request: ", code.(int16))
	return nil
}

func DecodeResponse(params map[string]interface {}) interface{} {
	code, ok := params["253"]

	if !ok {
		return nil
	}

	switch code.(int16) {
	}

	log.Print("Unhandled Response: ", code.(int16))
	return nil
}
