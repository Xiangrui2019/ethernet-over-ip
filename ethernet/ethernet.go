package ethernet

import (
	"log"

	"github.com/songgao/water"
)

type Ethernet struct {
	TapInterface *water.Interface
}

func NewEthernet(tap_interface_name string) *Ethernet {
	tap_interface, err := water.New(water.Config{
		DeviceType: water.TAP,
		PlatformSpecificParams: water.PlatformSpecificParams{
			Name: tap_interface_name,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return &Ethernet{
		TapInterface: tap_interface,
	}
}
