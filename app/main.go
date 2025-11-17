package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	i := 0
	for i == 0 {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {

			fmt.Fprintln(os.Stderr, "Error reading input:", err)

			os.Exit(1)

		}

		fmt.Println(command[:len(command)-1] + ": command not found")

	}
}
