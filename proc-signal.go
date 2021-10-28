package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
)

func main() {
	pid := flag.Int("pid", 0, "the pid of the process")
	flag.Parse()

	proc, err := os.FindProcess(*pid)
	if err != nil {
		panic(err)
	}

	if err := proc.Signal(syscall.Signal(0x0)); err != nil {
		fmt.Println("Process is not currently running!")
		return
	}

	if err := proc.Signal(syscall.Signal(0x1f)); err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Sent a Signal(0x1f) to the process!")
	}
}
