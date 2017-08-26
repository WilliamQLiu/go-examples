package main

import "os"
import "fmt"

func main() {

	// Command-Line Arguments used to parameterize execution of programs
	// e.g. `go run hello.go` uses `run` and `hello.go` arguments to the `go` program
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg := os.Args[3]

	// Usage: `cmd-example.go a b c d e`
	fmt.Println(argsWithProg)    // [./cmd-example a b c d e]
	fmt.Println(argsWithoutProg) // [a b c d e]
	fmt.Println(arg)             // c
}
