package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt");

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close();
	
	linesChannel := getLinesChannel(file);
	for currString := range linesChannel {
		if len(currString) > 0 {
			log.Printf("read: %s\n", currString);
		}
	}
}

func getLinesChannel(file io.ReadCloser) <-chan string {
	var currString string;
	var parts []string;

	data := make([]byte, 8);
	linesChannel := make(chan string);

	go func() {
		defer close(linesChannel)
		for {
			n, err := file.Read(data);
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
