package transport

import "log"

func TransportFactory(name string, ip string, port int) Transport {
	switch name {
	case "udp":
		transport, err := NewUDPTransport(ip, port)

		if err != nil {
			log.Fatal(err)
		}

		return transport
	}

	return nil
}
