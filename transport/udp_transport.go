package transport

import (
	"log"
	"net"
	"time"

	"github.com/songgao/water"
)

type UDPTransport struct {
	LocalUDPConnection *net.UDPConn
	TapInterface       *water.Interface
	RemoteAddr         *net.UDPAddr
}

func NewUDPTransport(ip string, port int, remote_ip string, remote_port int, l2_interface *water.Interface) *UDPTransport {
	udp, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	})

	if err != nil {
		return nil
	}

	return &UDPTransport{
		LocalUDPConnection: udp,
		TapInterface:       l2_interface,
		RemoteAddr: &net.UDPAddr{
			IP:   net.ParseIP(remote_ip),
			Port: remote_port,
		},
	}
}

func (transport *UDPTransport) L4ToL2() {
	buf := make([]byte, 65535)

	for {
		n, err := transport.TapInterface.Read(buf)

		if err != nil {
			log.Println("Read ethernet frames on tap driver error: ", err)
			time.Sleep(time.Millisecond * 10)
			continue
		}

		log.Println("Transfer ethernet frames number: ", n)
		log.Println("Transfer UDP Remote Addr: ", transport.RemoteAddr.String())

		if _, err := transport.LocalUDPConnection.WriteToUDP(buf[:n], transport.RemoteAddr); err != nil {
			log.Println("Write ethernet frames to udp error: ", err)
			time.Sleep(time.Millisecond * 10)
			continue
		}
	}
}

func (transport *UDPTransport) L2ToL4() {
	buf := make([]byte, 65535)

	for {
		n, readdr, err := transport.LocalUDPConnection.ReadFromUDP(buf)

		if err != nil {
			log.Println("Read ethernet frames on udp error: ", err)
			time.Sleep(time.Millisecond * 10)
			continue
		}

		log.Println("Transfer ethernet frames nuumber: ", n)
		log.Println("Transfer UDP Remote Addr: ", readdr)

		if _, err := transport.TapInterface.Write(buf[:n]); err != nil {
			log.Println("Write ethernet frames to tap driver error: ", readdr.String())
			time.Sleep(time.Millisecond * 10)
			continue
		}
	}
}
