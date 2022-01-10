package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

// func main() {
// 	os.Mkdir("/tmp/tmpdir/", 0777)
// 	defer os.RemoveAll("/tmp/tmpdir/")

// 	dir, _ := os.MkdirTemp("/tmp/tmpdir", "tmpdir-")

// 	println(dir)
// }

func main() {
	os.Mkdir("/tmp/tmpdir/", 0777)
	defer os.RemoveAll("/tmp/tmpdir/")

	os.Mkdir("/tmp/tmpdir/tmpdir-"+randToken(), 0777)

	entries, _ := os.ReadDir("/tmp/tmpdir/")
	for _, entry := range entries {
		println(entry.Name())
	}
}

func randToken() string {
	b := make([]byte, 4)
	rand.Read(b)
	token := fmt.Sprintf("%x", b)
	println(token)
	return token
}
