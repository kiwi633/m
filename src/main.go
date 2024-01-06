package main

import (
	"fmt"
	"os"
	"runtime"
	"src/repl"
)

func main() {
	fmt.Printf("M programming language!\n")
	fmt.Printf("Author: suntong\n")
	fmt.Printf("This is an interpreter program developed using Go \n")
	fmt.Printf("Go version: %s", runtime.Version())
	repl.Start(os.Stdin, os.Stdout)
}
