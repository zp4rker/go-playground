package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	input := flag.String("input", "", "the input to send to the process")
	flag.Parse()

	inFile, err := os.OpenFile("infile", os.O_WRONLY, 0755)
	if err != nil {
		panic(err.Error())
	}

	_, err = fmt.Fprint(inFile, *input + "\n")
	if err != nil {
		panic(err.Error())
	}
}