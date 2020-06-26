package transport

import (
	"bufio"
	"ethernet-over-ip/shared/ethernet"
	"log"
	"net"
)

type Client struct {
	TcpAddr   *net.TCPAddr
	TcpDialer *net.TCPConn
	Ethernet  *ethernet.Ethernet
}

func NewClient(server_addr string, eth *ethernet.Ethernet) *Client {
	tcpAddr, err := net.ResolveTCPAddr("tcp", server_addr)
	if err != nil {
		log.Fatal(err)
	}

	tcpDialer, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	client := &Client{
		TcpAddr:   tcpAddr,
		TcpDialer: tcpDialer,
		Ethernet:  eth,
	}

	return client
}

func (client *Client) Serve() {
	client.Handler(client.TcpDialer)

	defer client.TcpDialer.Close()
}

func (client *Client) Handler(conn *net.TCPConn) {
	go client.L2ToL4(conn)
	client.L4ToL2(conn)

	conn.Close()
}

func (client *Client) L4ToL2(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	buf := make([]byte, 65535)

	for {
		n, err := reader.Read(buf)

		if err != nil {
			log.Println("Read Ethernet Frames from TCP Network error: ", err)
			continue
		}

		if _, err := client.Ethernet.EthernetIface.Write(buf[:n]); err != nil {
			log.Println("Write Ethernet Frames to Tap driver error: ", err)
			continue
		}
	}
}

func (client *Client) L2ToL4(conn *net.TCPConn) {
	buf := make([]byte, 65535)

	for {
		n, err := client.Ethernet.EthernetIface.Read(buf)

		if err != nil {
			log.Println("Read Ethernet Frames from tap driver error: ", err)
			continue
		}

		if _, err := client.TcpDialer.Write(buf[:n]); err != nil {
			log.Println("Write Ethernet Frames to tcp error: ", err)
			continue
		}
	}
}
