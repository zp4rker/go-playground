package main

import (
	"fmt"
	"github.com/google/shlex"
	"os"
	"os/exec"
)

func main() {
	cmdSplit, err := shlex.Split(os.Args[1])
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(cmdSplit[0], cmdSplit[1:]...)

	if err = cmd.Start(); err != nil {
		panic(err)
	}

	fmt.Printf("PID: %v\n", cmd.Process.Pid)
}
