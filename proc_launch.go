package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	outFile, err := os.Create("outfile")
	if err != nil {
		panic(err.Error())
	}
	errFile, err := os.Create("errfile")
	if err != nil {
		panic(err.Error())
	}

	attr := &os.ProcAttr{
		Dir:   "",
		Env:   nil,
		Files: []*os.File{os.Stdin, outFile, errFile},
		Sys:   nil,
	}

	bin, err := exec.LookPath("go")
	if err != nil {
		panic(err.Error())
	}
	bin, err = filepath.Abs(bin)
	if err != nil {
		panic(err.Error())
	}

	proc, err := os.StartProcess(bin, strings.Fields("go run long_running.go"), attr)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Successfully started process with pid of %v\n", proc.Pid)
}
