package ethernet

import (
	"github.com/songgao/water"
	"log"
)

type Ethernet struct {
	EthernetIface *water.Interface
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
		EthernetIface: eth_interface,
	}
}
