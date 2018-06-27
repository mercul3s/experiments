package main

import (
	"fmt"
	"log"
	"net"
)

func tcpServer() {
	fmt.Println("Starting socket connection")
	listen, err := net.Listen("tcp", ":3030")
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("Accepting messages")
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleTCPConn(conn)
	}
}

func handleTCPConn(c net.Conn) {
	defer c.Close()
	var buff [512]byte

	for {
		_, err := c.Read(buff[0:])
		if err != nil {
			return
		}
		fmt.Println(string(buff[0:]))
	}
}

func udpServer() {
	fmt.Println("Starting UDP server")
	UDPServerAddr, err := net.ResolveUDPAddr("udp", ":3030")
	if err != nil {
		fmt.Println(err)
	}

	listen, err := net.ListenUDP("udp", UDPServerAddr)

	if err != nil {
		fmt.Println(err)
	}
	handleUDPMessages(listen)
}

func handleUDPMessages(c net.PacketConn) {
	defer c.Close()
	var buff [1500]byte
	for {
		readLen, _, err := c.ReadFrom(buff[0:])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(buff[:readLen]))
	}
}

func main() {
	//	tcpServer()
	udpServer()
}
