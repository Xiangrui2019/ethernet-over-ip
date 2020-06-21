package main

import (
	"ethernet-over-ip/ethernet"
	"ethernet-over-ip/transport"
	"flag"
	"log"
)

var (
	tap_interface_name string
	local_ip           string
	local_port         int64
	remote_ip          string
	remote_port        int64
)

func main() {
	flag.StringVar(&tap_interface_name, "tap_interface_name", "tap0", "")
	flag.StringVar(&local_ip, "local_ip", "0.0.0.0", "")
	flag.StringVar(&remote_ip, "remote_ip", "0.0.0.0", "")
	flag.Int64Var(&local_port, "local_port", 8888, "")
	flag.Int64Var(&remote_port, "remote_port", 8888, "")
	flag.Parse()
	log.Println("Welcome to use Ethernet Over Internet Protocol Agent.")

	eth := ethernet.NewEthernet(tap_interface_name)
	defer eth.TapInterface.Close()
	log.Println("Ethernet tap driver is init ok.")

	trans := transport.NewUDPTransport(local_ip, int(local_port), remote_ip, int(remote_port), eth.TapInterface)
	defer trans.LocalUDPConnection.Close()
	log.Println("UDP Transport driver is init ok.")

	go trans.L2ToL4()
	go trans.L4ToL2()
	log.Println("This Ethernet Over Internet Protocol Agent is running, Please don't close it.")
	select {}
}
