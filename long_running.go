package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
