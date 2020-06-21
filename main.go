package main

import (
	"ethernet-over-ip/ethernet"
	"ethernet-over-ip/transport"
	"log"
	"os"
	"strconv"
)

func main() {
	log.Println("Welcome to use Ethernet Over Internet Protocol Agent.")

	local_port, _ := strconv.Atoi(os.Getenv("LOCAL_PORT"))
	remote_port, _ := strconv.Atoi(os.Getenv("PEER_PORT"))

	eth := ethernet.NewEthernet(os.Getenv("TAP_INTERFACE_NAME"))
	defer eth.TapInterface.Close()
	log.Println("Ethernet tap driver is init ok.")

	trans := transport.NewUDPTransport(os.Getenv("LOCAL_IP"), int(local_port), os.Getenv("PEER_IP"), int(remote_port), eth.TapInterface)
	defer trans.LocalUDPConnection.Close()
	log.Println("UDP Transport driver is init ok.")

	go trans.L2ToL4()
	go trans.L4ToL2()
	log.Println("This Ethernet Over Internet Protocol Agent is running, Please don't close it.")
	select {}
}
