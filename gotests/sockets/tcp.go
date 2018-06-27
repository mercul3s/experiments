package main

import (
	"bufio"
	"fmt"
	"net"
)

func tcpClient() {
	conn, err := net.Dial("tcp", ":3030")
	if err != nil {
		fmt.Println("Failed to create client")
	}
	bytesWritten, err := conn.Write([]byte(" ping "))
	if err != nil {
		fmt.Println("Error is %s", err.Error())
	}
	fmt.Println("# of bytes written is %d", bytesWritten)
	readData := make([]byte, 4)
	numBytes, err := bufio.NewReader(conn).Read(readData)
	if err != nil {
		fmt.Println("Error is %s", err.Error())
	}
	fmt.Println("read %d bytes from client")
	fmt.Println(string(readData[:numBytes]))
	conn.Close()
}

func main() {
	tcpClient()
}
