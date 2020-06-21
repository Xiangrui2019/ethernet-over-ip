package transport

import "net"

type Transport interface {
	Write(buf []byte, remote_addr *net.Addr)
	Read(buf []byte)
}
