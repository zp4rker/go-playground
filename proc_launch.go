package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	goFile := flag.String("go-file", "long_running.go", "the go file to execute")
	flag.Parse()

	inFile, err := os.Create("infile")
	if err != nil {
		panic(err.Error())
	}
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
		Files: []*os.File{inFile, outFile, errFile},
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

	proc, err := os.StartProcess(bin, strings.Fields(fmt.Sprintf("go run %v", *goFile)), attr)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Successfully started process with pid of %v\n", proc.Pid)
}
