package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) > 2 {
		fmt.Println("Usage: todo add \"task description\"")
		return
	}

	command := os.Args[1]
	switch command {
		case "add":
			task := os.Args[2]
			fmt.Println("Adding task", task)

		default:
			fmt.Println("Unknown command", command)
	}
}
