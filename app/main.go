package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

func main() {
	builtInCommands := []string{"echo", "exit", "type"}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			os.Exit(1)
		}

		args := strings.Split(input, " ")
		command := strings.TrimSpace(args[0])

		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(input[5 : len(input)-1])
		case "type":
			if slices.Contains(builtInCommands, strings.TrimSpace(args[1])) {
				fmt.Println(strings.TrimSpace(args[1]), "is a shell builtin")
			} else {
				fmt.Println(strings.TrimSpace(args[1]) + ": command not found")
			}
		default:
			fmt.Println(input[:len(input)-1] + ": command not found")
		}
	}
}
