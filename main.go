// main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add', 'list', or 'done' subcommands")
		return
	}
	switch os.Args[1] {
	case "help":
		fmt.Println("Available commands: add, list, done")
	default:
		fmt.Println("Unknown command. Use 'help'")
	}
}
