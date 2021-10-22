package main

import (
	"bufio"
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

	go func() {
		buf := make([]byte, 1024)
		var n int
		var err error
		for err == nil {
			n, err = conn.Read(buf)
			fmt.Print(string(buf[:n]))
		}
	}()

	rd := bufio.NewReader(os.Stdin)
	var input string
	for err == nil {
		input, err = rd.ReadString('\n')
		if err != nil {
			continue
		}
		_, err = conn.Write([]byte(input))
	}
}
