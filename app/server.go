package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	for {
		var b = make([]byte, 128)
		_, err := conn.Read(b)
		if err != nil {
			conn.Close()
			break
		}
		fmt.Println(string(b))
		conn.Write([]byte("+PONG\r\n"))
	}
}
