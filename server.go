package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr := "/tmp/go-playground.sock"

	if err := os.RemoveAll(addr); err != nil {
		panic(err.Error())
	}

	fmt.Println("Started server!")
	listener, err := net.Listen("unix", addr)
	if err != nil {
		panic(err.Error())
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err.Error())
		}

		if !handleConnection(conn) {
			break
		}
	}
}

func handleConnection(conn net.Conn) bool {
	buf := make([]byte, 512)
	nr, err := conn.Read(buf)
	if err != nil {
		panic(err.Error())
	}

	input := string(buf[:nr])
	fmt.Printf("[Client]: %v\n", input)
	if input == "close" {
		fmt.Println("Closing connection now...")
		_, err := conn.Write([]byte("Closing connection now..."))
		if err != nil {
			panic(err.Error())
		}
		conn.Close()
		return false
	} else {
		_, err := conn.Write(buf[:nr])
		if err != nil {
			panic(err.Error())
		}
		return true
	}
}
