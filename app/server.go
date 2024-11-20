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

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		defer conn.Close()

		errCh := make(chan error)
		fmt.Println("Channel: ", <-errCh)
		go handlePing(conn, errCh)
		// if <-errCh != nil {
		// 	fmt.Println("Error handling connection: ", err.Error())
		// 	os.Exit(1)
		// }
	}
}

func handlePing(conn net.Conn, errCh chan<- error) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		errCh <- err
		return
	}

	_, err = conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		errCh <- err
		return
	}

	errCh <- nil
}
