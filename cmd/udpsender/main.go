package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "localhost:4200"

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatalf("Error resolving address: %v", err)
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf("Error opening UDP connection: %v", err)
	}
	defer udpConn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading from stdin: %v", err)
		}

		n, err := udpConn.WriteToUDP([]byte(line), udpAddr)
		if err != nil {
			log.Fatalf("Error sending UDP packet: %v", err)
		}
		fmt.Printf("Sent %d bytes\n", n)
	}

}