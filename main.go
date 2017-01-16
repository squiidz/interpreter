package main

import (
	"fmt"
	"local/interpret/repl"
	"os"
)

func main() {
	fmt.Printf("<- Monkey REPL ->\n")
	repl.Start(os.Stdin, os.Stdout)
}
