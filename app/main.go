package main

import (
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {
	var i string
	fmt.Fprint(os.Stdout, "$ ")
	fmt.Scanln(&i)

	if i != nil {
		fmt.Fprint(os.Stdout, i, ": command not found")
	}

}
