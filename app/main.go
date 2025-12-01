package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"os/exec"
	"path/filepath"
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
		commandExists, commandPath := checkPath(command)
		if commandExists && !slices.Contains(builtInCommands, command){
			match := command
		}

		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(input[5:])
		case "type":
			pathExist, pathLoc := checkPath(args[1])
			if slices.Contains(builtInCommands, args[1]) {
				fmt.Println(args[1], "is a shell builtin")
			} else if pathExist {
				fmt.Println(args[1], "is", pathLoc)
			} else {
				fmt.Println(args[1] + ": not found")
			}
		case match:
			runExternal(commandPath, args[1:])
		default:
			fmt.Println(input + ": command not found")
		}
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

		// Check if any execute bit is set
		if info.Mode()&0111 != 0 {
			return true, full
		}
	}
	return false, ""
}

func runExternal(path string, args []string) {
    cmd := exec.Command(path, args...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        fmt.Println("error executing:", err)
    }
}
