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
		rd := bufio.NewReader(conn)
		var str string
		var err error
		for err == nil {
			str, err = rd.ReadString('\n')
			fmt.Print(str)
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
