package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var c, python, java bool // the 'type' (e.g. bool) goes after the variable name
var d, e int = 1, 2      // you can initialize multiple variables

const Pi = 3.14 // you can make a constant with =. constants cannot be used with := syntax

func add(x, y int) int { // same as: x int, y int
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) { // pass in a named variable 'sum'
	x = sum * 4 / 9
	y = sum - x
	return // a 'naked' return (i.e. no arguments) returns the named values
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Print("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Printf("Hello, world\n")                        // Default format with Printf
	fmt.Println("My favorite number is", rand.Intn(10)) // Print line

	fmt.Println("These two numbers added together are:", add(42, 13)) // using the 'add' function we made

	a, b := swap("First", "Second") // using the 'swap' function we made
	fmt.Println(a, b)

	fmt.Println(split(17)) // using the 'split' function we made

	var i int
	fmt.Println(i, c, python, java) // by default, things are 0 or false (depending on type)

	var c, python, java = true, false, "No!" // you can initialize variables
	fmt.Println(d, e, c, python, java)

	w := 42         // same as: var w int = 42
	x := float64(w) // doing a type conversion from int to float64, same as: var f float64 = float64(i)
	y := uint(x)    // a uint is a 'byte', same as: var u uint = uint(f)
	z := rune(y)    // a rune is an `int32`, mainly for Unicode code point
	fmt.Println(z)  // If you declare a variable, you have to use it or else it's an error!
	fmt.Println("Pi is", Pi)
	fmt.Println("The time is", time.Now())

	sum := 0                          // a short variable declaration means let Go figure out the type
	fmt.Println("Running a for-loop") // for-loop
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("The sum is:", sum, "\nThe var 'w' is: ", w) // how to print different lines

	fmt.Println("Running a if statement") // if statement
	if w > 0 {
		fmt.Println("Yo, w is small")
	}

	fmt.Println(pow(3, 2, 10))

	// switch statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("%s", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	defer fmt.Println("world")
	fmt.Println("hello")

	// Pointers
	var d *int // d will redirect to an int
	fmt.Println("d pointer is: ", d)
	// d pointer is: <nil>

	e := 1009
	f := &e // give value to pointer
	fmt.Println("f reference to e is: ", f)
	// f reference to e is: 0xc420000e448

	// set e through pointer f
	*f = 2000 // use * to get the pointer's underlying value (e.g. value of e)
	fmt.Println("e has been set through pointer f:", e)
	fmt.Println("f looks like:", f)
	fmt.Println("*f looks like:", *f)

	// Arrays
	var g [2]string
	g[0] = "Hello"
	g[1] = "World"
	fmt.Println(g[0], g[1])
	fmt.Println(g)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices of array
	var s []int = primes[1:4]
	fmt.Println(s) // [3, 5, 7]

	// Append to array
	s = append(s, 23)
	fmt.Println(s)

	// Range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for index, value := range pow {
		fmt.Printf("2**%d = %d\n", index, value)
	}

	// Maps
	m := make(map[string]int) // Make a map of string to int
	m["Answer"] = 42          // Assign or update map 'm' on key 'Answer' with value '42'
	delete(m, "Answer")       // Delete an element with the key 'Answer'
	v, ok := m["Answer"]      // If 'v' or 'ok' have not been declared yet, use this shorthand
	fmt.Println("v is: ", v, " and 'ok' is: ", ok)

}
