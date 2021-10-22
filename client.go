package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr := "/tmp/go-playground.sock"

	conn, err := net.Dial("unix", addr)
	if err != nil {
		panic(err.Error())
	}

	_, err = conn.Write([]byte(os.Args[1]))
	if err != nil {
		panic(err.Error())
	}

	buf := make([]byte, 1024)
	nr, err := conn.Read(buf)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("[Server]: %v\n", string(buf[:nr]))
}
