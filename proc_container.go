package main

import (
	"bufio"
	"fmt"
	"github.com/google/shlex"
	"os"
	"os/exec"
)

func main() {
	cmdSplit, err := shlex.Split(os.Args[1])
	cmd := exec.Command(cmdSplit[0], cmdSplit[1:]...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err.Error())
	}

	if err := cmd.Start(); err != nil {
		panic(err.Error())
	}

	dir := fmt.Sprintf("/tmp/go-playground/%v", cmd.Process.Pid)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err.Error())
	}
	outfile, err := os.OpenFile(fmt.Sprintf("%v/out", dir), os.O_WRONLY | os.O_CREATE, 0644)
	if err != nil {
		panic(err.Error())
	}
	defer outfile.Close()
	wr := bufio.NewWriter(outfile)

	go func() {
		buf := make([]byte, 64)
		var n int
		for err == nil {
			n, err = stdout.Read(buf)
			if n > 0 {
				fmt.Print(string(buf[:n]))
				if _, err := wr.Write(buf[:n]); err != nil {
					panic(err.Error())
				}
			}
		}
	}()

	if err := cmd.Wait(); err != nil {
		panic(err.Error())
	}

	if err := wr.Flush(); err != nil {
		panic(err.Error())
	}
}
