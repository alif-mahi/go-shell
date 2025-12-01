package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
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

		input = strings.TrimSuffix(input, "\n")
		args := strings.Split(input, " ")
		command := args[0]

		// First: check if it's a built-in command
		switch command {
		case "exit":
			os.Exit(0)

		case "echo":
			if len(input) > 5 {
				fmt.Println(input[5:])
			} else {
				fmt.Println()
			}
			continue

		case "type":
			if len(args) < 2 {
				fmt.Println("type: missing argument")
				continue
			}
			target := args[1]

			if slices.Contains(builtInCommands, target) {
				fmt.Println(target, "is a shell builtin")
			} else if ok, loc := checkPath(target); ok {
				fmt.Println(target, "is", loc)
			} else {
				fmt.Println(target + ": not found")
			}
			continue
		}

		// If not a builtin, try running it as an external command
		if ok, path := checkPath(command); ok {
			runExternal(path, args[1:])
			continue
		}

		// Otherwise: unknown command
		fmt.Println(command + ": command not found")
	}
}

func checkPath(command string) (bool, string) {
	pathEnv := os.Getenv("PATH")
	for _, dir := range strings.Split(pathEnv, string(os.PathListSeparator)) {
		full := filepath.Join(dir, command)
		info, err := os.Stat(full)
		if err != nil || info.IsDir() {
			continue
		}

		// Check if executable
		if info.Mode()&0111 != 0 {
			return true, full
		}
	}
	return false, ""
}

func runExternal(path string, args []string) {
	cmd := exec.Command(filepath.Base(path), args...)
	cmd.Path = path

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "error executing:", err)
	}
}
