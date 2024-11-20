package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	for {
		err = handlePing(conn)
		if err != nil {
			fmt.Println("Error handling connection: ", err.Error())
			os.Exit(1)
		}
	}
}

func handlePing(conn net.Conn) error {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return err
	}

	fmt.Printf("Recieved: %s\n", buf[:n])

	_, err = conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		return err
	}

	return nil
}
