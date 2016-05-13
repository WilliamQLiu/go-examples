package main

import (
    "fmt"
    "math/rand"
)

var c, python, java bool  // the 'type' (e.g. bool) goes after the variable name
var d, e int = 1, 2  // you can initialize variables

const Pi = 3.14  // you can make a constant with =. constants cannot be used with := syntax

func add(x, y int) int {  // same as: x int, y int
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

func split(sum int) (x, y int) {  // pass in a named variable 'sum'
    x = sum * 4 / 9
    y = sum - x
    return  // a 'naked' return (i.e. no arguments) returns the named values
}

func main() {
    fmt.Printf("Hello, world\n")  // Default format with Printf
    fmt.Println("My favorite number is", rand.Intn(10))  // Print line

    fmt.Println("These two numbers added together are:", add(42, 13))  // using the 'add' function we made

    a, b := swap("First", "Second")  // using the 'swap' function we made
    fmt.Println(a, b)

    fmt.Println(split(17))  // using the 'split' function we made

    var i int
    fmt.Println(i, c, python, java)  // by default, things are 0 or false (depending on type)

    var c, python, java = true, false, "No!"  // you can initialize variables
    fmt.Println(d, e, c, python, java)

    k := 3  // a short variable declaration means let Go figure out the type
    fmt.Println(k)

    w := 42  // same as: var w int = 42
    x := float64(w)  // doing a type conversion from int to float64, same as: var f float64 = float64(i)
    y := uint(x)  // a uint is a 'byte', same as: var u uint = uint(f)
    z := rune(y)  // a rune is an `int32`, mainly for Unicode code point
    fmt.Println(z)  // If you declare a variable, you have to use it or else it's an error!
    fmt.Println("Pi is", Pi)
}
