package util

import (
	"strings"

	"github.com/drazen-todorovic/netstatex/model"
)

//ParseNetstatOutput transforms netstat input to slice of NetstatTCPItem, NetstatUDPItem
func ParseNetstatOutput(input string) ([]model.NetstatTCPItem, []model.NetstatUDPItem) {
	lines := strings.Split(input, "\n")

	itemsTCP := []model.NetstatTCPItem{}
	itemsUDP := []model.NetstatUDPItem{}

	for _, line := range lines {
		values := strings.Fields(line)
		valuesLen := len(values)
		if valuesLen > 0 {
			switch values[0] {
			case model.TCP:
				itemTCP := model.NetstatTCPItem{
					Protocol:       model.TCP,
					LocalAddress:   values[1],
					ForeignAddress: values[2],
					State:          values[3],
				}
				itemsTCP = append(itemsTCP, itemTCP)
			case model.UDP:
				itemUDP := model.NetstatUDPItem{
					Protocol:       model.UDP,
					LocalAddress:   values[1],
					ForeignAddress: values[2],
				}
				itemsUDP = append(itemsUDP, itemUDP)
			}

		}
	}
	return itemsTCP, itemsUDP
}
