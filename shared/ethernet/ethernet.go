package ethernet

import (
	"log"
	"runtime"

	"github.com/songgao/water"
)

type Ethernet struct {
	EthernetInterface *water.Interface
}

func NewEthernet(tap_interface_name string) *Ethernet {
	var err error
	var eth_interface *water.Interface

	eth_interface, err = water.New(water.Config{
		DeviceType: water.TAP,
		PlatformSpecificParams: water.PlatformSpecificParams{
			Name:       tap_interface_name,
			Persist:    true,
			MultiQueue: true,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return &Ethernet{
		EthernetInterface: eth_interface,
	}
}
