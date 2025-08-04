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
	case "add":
		task := os.Args[2]
		fmt.Printf("Task added: %s\n", task)
	default:
		fmt.Println("Unknown command. Use 'help'")
	}
}
