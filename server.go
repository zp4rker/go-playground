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

	newConn := make(chan net.Conn)
	quit := make(chan bool, 1)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				panic(err.Error())
			}

			newConn<- conn
		}
	}()

	loop: for {
		select {
		case <-quit:
			break loop
		case conn := <-newConn:
			handleConnection(conn, quit)
		}
	}
}

func handleConnection(conn net.Conn, close chan bool) {
	username := "Client"

	buf := make([]byte, 1024)
	nr, err := conn.Read(buf)
	if err != nil {
		panic(err.Error())
	}

	input := string(buf[:nr])
	fmt.Printf("[%v]: %v\n", username, input)
	if input == "close" {
		fmt.Println("Closing connection now...")
		_, err := conn.Write([]byte("Closing connection now..."))
		if err != nil {
			panic(err.Error())
		}
		conn.Close()
		close<- true
	} else {
		_, err := conn.Write(buf[:nr])
		if err != nil {
			panic(err.Error())
		}
	}
}
