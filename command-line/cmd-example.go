package main

import "os"
import "flag" // supports basic command-line flag parsing
import "fmt"

func main() {

	/*
		START COMMAND-LINE ARGUMENTS
		Command-Line Arguments used to parameterize execution of programs
		e.g. `go run hello.go` uses `run` and `hello.go` arguments to the `go` program
	*/
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	//arg := os.Args[3]  // Get the third argument

	// Usage: `cmd-example.go a b c d e`
	fmt.Println(argsWithProg)    // [./cmd-example a b c d e]
	fmt.Println(argsWithoutProg) // [a b c d e]
	//fmt.Println(arg)             // c
	// END COMMAND-LINE ARGUMENTS

	/*
		START COMMAND-LINE FLAGS
		Command-Line Flags are used to specify options for a command-line programs. For example, in `wc -l` the `-l` is a command-line flag.

		Basic flag declarations are available from the 'flag' package for string, integer, boolean options
		For example, we declare a string flag 'word' with a default value of 'foo' and a short description of 'a string'.
	*/

	// Important to note that a flag.String function returns a string pointer (not a string value)
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// You can declare an option that uses an existing var declared elsewhere in the program. You need to pass in a pointer to the flag declaration function
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse() // once all flags are declared, execute the command-line parsing

	fmt.Println("word:", *wordPtr)    // word: foo
	fmt.Println("numb:", *numbPtr)    // numb: 42
	fmt.Println("fork:", *boolPtr)    // fork: false
	fmt.Println("svar:", svar)        //svar: bar
	fmt.Println("tail:", flag.Args()) // e.g. [a b c d e] when ./cmd-example a b c d e
	// END COMMAND-LINE FLAGS

}
