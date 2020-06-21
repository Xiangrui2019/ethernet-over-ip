package ethernet

import (
	"log"
	"runtime"

	"github.com/songgao/water"
)

type Ethernet struct {
	TapInterface *water.Interface
}

func NewEthernet(tap_interface_name string) *Ethernet {
	var err error
	var tap_interface *water.Interface

	if runtime.GOOS == "linux" {
		tap_interface, err = water.New(water.Config{
			DeviceType: water.TAP,
			PlatformSpecificParams: water.PlatformSpecificParams{
				Name:       tap_interface_name,
				Persist:    true,
				MultiQueue: true,
			},
		})
	}

	if runtime.GOOS == "windows" {
		tap_interface, err = water.New(water.Config{
			DeviceType: water.TAP,
		})
	}

	if err != nil {
		log.Fatal(err)
	}

	return &Ethernet{
		TapInterface: tap_interface,
	}
}
