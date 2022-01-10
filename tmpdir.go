package main

import (
	"os"
)

func main() {
	os.Mkdir("/tmp/tmpdir/", 0777)
	defer os.RemoveAll("/tmp/tmpdir/")

	dir, _ := os.MkdirTemp("/tmp/tmpdir", "tmpdir-")

	println(dir)
}
