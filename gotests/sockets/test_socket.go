package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func sendSocketMessage(protocol string, message string) {
	socketClient, _ := net.Dial(protocol, ":3030")
	readData := make([]byte, 500)
	fmt.Printf("Sending message %s over %s \n", message, protocol)
	socketClient.Write([]byte(message))
	if protocol == "tcp" {
		readBytes, _ := bufio.NewReader(socketClient).Read(readData)
		fmt.Printf("Read response from server: %s \n", readData[:readBytes])
	}
	time.Sleep(1 * time.Second)
}

func main() {
	simpleMessage := `{"timestamp": 123}`
	checkMessage := `{"timestamp": 12345, "check":{"config": {"name": "turtles"}, "status": 1}}`
	invalidMessage := `{"timestamp:`

	sendSocketMessage("tcp", simpleMessage)
	sendSocketMessage("udp", simpleMessage)
	sendSocketMessage("tcp", invalidMessage)
	sendSocketMessage("udp", invalidMessage)
	sendSocketMessage("tcp", checkMessage)
	sendSocketMessage("udp", checkMessage)
}
