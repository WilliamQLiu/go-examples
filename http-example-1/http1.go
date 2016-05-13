// Run using `go run http1.go`

package main

import (
    "fmt"
    "io"
    "net/http"
)

// We have a http.ResponseWriter (and its corresponding stream)
// Because we're using io, we have a response stream
func hello1(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "This is a Stream!")  // this is a writable Stream
}

func main() {
    fmt.Printf("Server is Running...")

    // HandleFunc accepts two args, first arg is string pattern representing the route
    // second arg is a function with the signature func(ResponseWriter, *Request)
    http.HandleFunc("/", hello1)
    http.ListenAndServe(":8000", nil)
}
