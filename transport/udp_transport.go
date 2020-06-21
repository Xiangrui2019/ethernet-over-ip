package transport

import "net"

type UDPTransport struct {
	LocalConnection *net.UDPConn
}

func NewUDPTransport(ip string, port int) (Transport, error) {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   ip,
		Port: port,
	})

	if err != nil {
		return nil, err
	}

	return &UDPTransport{
		LocalConnection: conn,
	}, nil
}

func (udp *UDPTransport) Write(buf []byte, remoteAddr *net.Addr) {

}

func (udp *UDPTransport) Read(buf []byte) {

}
