package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func GetLineChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1) //rds between the channel and the function


	go func() {
		defer f.Close()
		defer close(out)

		str := ""

		for {

			data := make([]byte, 8)
			n, err := f.Read(data) // counts the bytes in the line
			if err != nil {
				break
			}
			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 { // index the data while looking for the end of the line to break the line
				str += string(data[:i]) // adds the line to the string after converting from bytes to 
				data = data[i+1:] // beggins the new line
				out <- str
				str = ""
			}
			str += string(data)

			if len(str) != 0 {
				out <- str
			}

		} // closes for loop

	}() // closes and invokes go func

	return out
}

func main() {
	listener, err := net.Listen("tcp", ":42069")

	if err != nil {
		log.Fatal("error", "error", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}

		for line := range GetLineChannel(conn) {
			fmt.Printf("read: %s\n", line)

		}

	}

}
