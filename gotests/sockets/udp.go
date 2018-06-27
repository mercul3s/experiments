package main

import (
	"fmt"
	"net"
	"sync"
)

type UDPServer struct {
	stopped chan struct{}
}

func (u *UDPServer) client() net.Conn {
	conn, err := net.Dial("udp", "localhost:5050")
	if err != nil {
		fmt.Println(err)
	}
	return conn
	//defer conn.Close()
}

func (u *UDPServer) udpServer() {
	fmt.Println("Starting UDP server")
	UDPServerAddr, err := net.ResolveUDPAddr("udp", ":5050")
	if err != nil {
		fmt.Println(err)
	}

	listen, err := net.ListenUDP("udp", UDPServerAddr)

	if err != nil {
		fmt.Println(err)
	}
	go func() {
		fmt.Println("Handling messages")
		u.handleUDPMessages(listen)
	}()
}

func (u *UDPServer) handleUDPMessages(c net.PacketConn) {
	fmt.Println("In handle messages")
	var wg sync.WaitGroup
	defer c.Close()
	var buff [1500]byte

	wg.Add(1)
	for {
		select {
		case <-u.stopped:
			fmt.Println("In stopped case")
			return
		default:
			readLen, _, err := c.ReadFrom(buff[0:])
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(string(buff[:readLen]))
		}
	}
	err := c.Close()
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
}

func clientCycle(u UDPServer, payload string) {
	u.udpServer()
	client := u.client()
	client.Write([]byte(payload))
	client.Close()
	close(u.stopped)
}
func main() {
	u := UDPServer{
		stopped: make(chan struct{}),
	}
	payload1 := `{"timestamp":123}`
	payload2 := `{"timestamp":1234}`
	payload3 := `{"timestamp":12345}`
	clientCycle(*u, payload1)
	clientCycle(*u, payload2)
	clientCycle(*u, payload3)
}
