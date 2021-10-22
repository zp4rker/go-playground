package main

import (
	"fmt"
	"github.com/google/shlex"
	"net"
	"os"
	"strings"
)

var connections = make([]net.Conn, 0)

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
				if strings.Contains(err.Error(), "use of closed network connection") {
					return
				} else {
					panic(err.Error())
				}
			}

			newConn <- conn
		}
	}()

loop:
	for {
		select {
		case <-quit:
			broadcast("[Server]: Closing connection now...\n")
			break loop
		case conn := <-newConn:
			connections = append(connections, conn)
			go handleConnection(conn, quit)
			broadcast(fmt.Sprintf("[Server]: User%v joined!\n", len(connections)))
		}
	}
}

func broadcast(msg string) {
	fmt.Print(msg)
	for _, conn := range connections {
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("WARN: Failed to write to a connection!")
		}
	}
}

func handleConnection(conn net.Conn, quit chan bool) {
	username := fmt.Sprintf("User%v", len(connections))

	buf := make([]byte, 1024)
	var n int
	var err error = nil
	for err == nil {
		n, err = conn.Read(buf)
		if n > 0 {
			input := string(buf[:n])
			output := handleInput(&username, input, quit)
			if output != "" {
				msg := fmt.Sprintf("[%v]: %v", username, output)
				broadcast(msg)
			}
		}
	}
}

func handleInput(username *string, input string, quit chan bool) string {
	if input[0] == '/' {
		parts, err := shlex.Split(input[1:])
		if err != nil {
			return ""
		}
		cmd, args := strings.ToLower(parts[0]), make([]string, 0)
		if len(parts) > 1 {
			args = parts[1:]
		}
		switch cmd {
		case "close", "quit":
			quit <- true
			return ""
		case "username":
			if len(args) == 1 {
				oldUsername := *username
				newUsername := args[0]
				username = &newUsername
				return fmt.Sprintf("Changed username from '%v'", oldUsername)
			}
		}
		return ""
	} else {
		return input
	}
}
