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
}

func NewUDPTransport(ip string, port int, l2_interface *water.Interface) *UDPTransport {
	return &UDPTransport{
		LocalUDPConnection: net.ListenUDP("udp", &net.UDPAddr{
			IP:   ip,
			Port: port,
		}),
		TapInterface: l2_interface,
	}
}

func (transport *UDPTransport) L4ToL2() {
	buf := make([]byte, 65535)

	for {
		n, err := transport.TapInterface.Read(buf)

		if err != nil {
			log.Println(err)
			time.Sleep(time.Millisecond * 2)
			continue
		}

		if _, err := transport.LocalUDPConnection.WriteToUDP(buf[:n]); err != nil {
			log.Println(err)
			time.Sleep(time.Millisecond * 2)
			continue
		}
	}
}

func (transport *UDPTransport) L2ToL4() {
	buf = make([]byte, 65535)

	for {
		n, _, err := transport.LocalUDPConnection.ReadFromUDP(buf)

		if err != nil {
			log.Println(err)
			time.Sleep(time.Millisecond * 2)
			continue
		}

		if _, err := transport.TapInterface.Write(buf[:n]); err != nil {
			time.Sleep(time.Millisecond)
			continue
		}
	}
}
