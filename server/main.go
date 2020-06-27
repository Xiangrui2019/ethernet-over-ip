package main

import (
	_ "ethernet-over-ip/conf"
	"ethernet-over-ip/shared/ethernet"
	"ethernet-over-ip/shared/transport"
	"os"
)

func main() {
	eth := ethernet.NewEthernet(os.Getenv("ETHERNET_IFACE_NAME"))
	server := transport.NewServer(os.Getenv("SERVER_ADDR"), eth, os.Getenv("PRE_SHARED_KEY"))

	server.Serve()
}
