package transport

import (
	"bufio"
	"ethernet-over-ip/shared/ethernet"
	"log"
	"net"
)

type Server struct {
	TcpAddr     *net.TCPAddr
	TcpListener *net.TCPListener
	Ethernet    *ethernet.Ethernet
}

func NewServer(server_addr string, eth *ethernet.Ethernet) *Server {
	tcpAddr, err := net.ResolveTCPAddr("tcp", server_addr)
	if err != nil {
		log.Fatal(err)
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	server := &Server{
		TcpAddr:     tcpAddr,
		TcpListener: tcpListener,
		Ethernet:    eth,
	}

	return server
}

func (server *Server) Serve() {
	for {
		tcpConnection, err := server.TcpListener.AcceptTCP()

		if err != nil {
			log.Println("Accept TCP Connection error.")
			continue
		}

		go server.Handler(tcpConnection)
	}

	defer server.TcpListener.Close()
}

func (server *Server) Handler(conn *net.TCPConn) {
	go server.L2ToL4(conn)
	server.L4ToL2(conn)

	conn.Close()
}

func (server *Server) L4ToL2(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	buf := make([]byte, 65535)

	for {
		n, err := reader.Read(buf)

		if err != nil {
			log.Println("Read Ethernet Frame from TCP error: ", err)
			continue
		}

		if _, err := server.Ethernet.EthernetIface.Write(buf[:n]); err != nil {
			log.Println("Write Ethernet Frame to Ethernet Tap Driver error: ", err)
			continue
		}
	}
}

func (server *Server) L2ToL4(conn *net.TCPConn) {
	buf := make([]byte, 65535)

	for {
		n, err := server.Ethernet.EthernetIface.Read(buf)

		if err != nil {
			log.Println("Read Ethernet Frame from tap driver error: ", err)
			continue
		}

		if _, err := conn.Write(buf[:n]); err != nil {
			log.Println("Write Ethernet Frame to tcp streams error: ", err)
			continue
		}
	}
}
