package transport

import (
	"bufio"
	"ethernet-over-ip/shared/ethernet"
	"log"
	"net"
	"time"
)

type Client struct {
	TcpAddr      *net.TCPAddr
	TcpDialer    *net.TCPConn
	Ethernet     *ethernet.Ethernet
	PreSharedKey string
}

func NewClient(server_addr string, eth *ethernet.Ethernet, pre_shared_key string) *Client {
	tcpAddr, err := net.ResolveTCPAddr("tcp", server_addr)
	if err != nil {
		log.Fatal(err)
	}

	client := &Client{
		TcpAddr:      tcpAddr,
		TcpDialer:    nil,
		Ethernet:     eth,
		PreSharedKey: pre_shared_key,
	}

	return client
}

func (client *Client) Serve() {
	var err error

	for {
		client.TcpDialer, err = net.DialTCP("tcp", nil, client.TcpAddr)

		if err != nil {
			log.Println("ReConnect Faild: ", err)
			continue
		}

		client.Handler(client.TcpDialer)
		time.Sleep(time.Second * 20)
	}

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
	error_rate := 0

	for {
		if error_rate > 1000 {
			break
		}

		n, err := reader.Read(buf)

		if err != nil {
			log.Println("Read Ethernet Frames from TCP Network error: ", err)
			time.Sleep(time.Millisecond * 10)
			error_rate = error_rate + 1
			continue
		}

		if _, err := client.Ethernet.EthernetIface.Write(buf[:n]); err != nil {
			log.Println("Write Ethernet Frames to Tap driver error: ", err)
			error_rate = error_rate + 1
			continue
		}
	}
}

func (client *Client) L2ToL4(conn *net.TCPConn) {
	buf := make([]byte, 65535)
	error_rate := 0

	for {
		if error_rate > 1000 {
			break
		}

		n, err := client.Ethernet.EthernetIface.Read(buf)

		if err != nil {
			log.Println("Read Ethernet Frames from tap driver error: ", err)
			error_rate = error_rate + 1
			continue
		}

		if _, err := client.TcpDialer.Write(buf[:n]); err != nil {
			log.Println("Write Ethernet Frames to tcp error: ", err)
			error_rate = error_rate + 1
			continue
		}
	}
}
