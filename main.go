package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const port = ":4200"

func main() {
	listener, err := net.Listen("tcp", port)
	fmt.Println("Server is listening on port", port)

	if err != nil {
		log.Fatalf("Error opening TCP connection: %v", err)
	}
	defer listener.Close()
	var conn net.Conn	

	for {
		conn, err = listener.Accept()
		if err != nil {
			log.Fatalf("Error accepting connection: %v", err.Error())
		}
		fmt.Printf("Accepted connection from %v", conn.RemoteAddr())

		linesChannel := getLinesChannel(conn)
		for currentLine := range linesChannel {
			if len(currentLine) > 0 {
				fmt.Printf("read: %s\n", currentLine)
			}
		}
		fmt.Println("Connection to", conn.RemoteAddr(), "closed")
	}
}

func getLinesChannel(connection io.ReadCloser) <-chan string {
	var currString string
	var parts []string

	data := make([]byte, 8);
	linesChannel := make(chan string);

	go func() {
		defer connection.Close()
		defer close(linesChannel)
		for {
			n, err := connection.Read(data)
			if err != nil || n == 0 {
				break
			}

			parts = strings.Split(string(data[:n]), "\n");
			if len(parts) > 1 {
				linesChannel <- currString + parts[0];
				currString = parts[1];
			} else {
				currString += parts[0];
			}
		}

		if len(currString) > 0 {
			linesChannel <- currString;
		}
	}()
	return linesChannel;
}
