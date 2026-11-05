package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: todo add|list|done|delete")
		os.Exit(1)
	}

	cmd := os.Args[1]
	fmt.Println("command:", cmd)
}
