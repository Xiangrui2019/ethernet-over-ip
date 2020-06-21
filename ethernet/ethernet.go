package ethernet

import (
	"github.com/songgao/water"
)

type Ethernet struct {
	TapInterface *water.Interface
}

func NewEthernet(tap_interface_name string) *Ethernet {
	return &Ethernet{
		TapInterface: water.New(water.Config{
			DeviceType: water.TAP,
			PlatformSpecificParams: water.PlatformSpecificParams{
				Name: tap_interface_name,
			},
		}),
	}
}
