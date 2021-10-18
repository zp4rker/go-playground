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

	if *pid == 0 {
		fmt.Println("Invalid process!")
		return
	}

	fmt.Println("Searching for process...")
	proc, err := os.FindProcess(*pid)
	if err != nil {
		fmt.Println("Unable to find process!")
		panic(err.Error())
	}

	err = proc.Signal(syscall.Signal(0))
	if err == nil {
		fmt.Println("Process is still running!")
	} else if err == os.ErrProcessDone {
		fmt.Println("Process is no longer running!")
	}
}
