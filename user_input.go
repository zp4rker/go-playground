package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your full name: ")
	name, _ := reader.ReadString('\n')

	firstName := strings.Fields(name)[0]
	fmt.Printf("Hello, %v! Thanks for running this application.\n", firstName)

	fmt.Println("Press enter to exit...")
	_, _ = reader.ReadString('\n')
}