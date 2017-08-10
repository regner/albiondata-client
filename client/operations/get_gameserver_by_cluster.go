package operations

import (
	"github.com/regner/albionmarket-client/client/albionstate"
	"log"
	"strconv"
	"strings"
)

//10
type GetGameServerByCluster struct {
	ZoneID string    							`mapstructure:"0"`
	NrSessionZoneChanges int32    `mapstructure:"255"`
}

func (op GetGameServerByCluster) Process(state *albionstate.AlbionState) {
	if strings.Index(op.ZoneID, ".") != -1 {
		return
	}
	log.Print("Got GetGameServerByCluster operation.\n Params: ZoneID:" + op.ZoneID + ", NrSessionZoneChanges:" + strconv.Itoa(int(op.NrSessionZoneChanges)))
}

