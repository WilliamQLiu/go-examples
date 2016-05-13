// Run using `go run http2.go`
// We added in a Mux to help with routing

package main

import (
    "fmt"
    "io"
    "net/http"
)

// We have a http.ResponseWriter (and its corresponding stream)
// Because we're using io, we have a response stream
func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello")  // this is a writable Stream
}

func world(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "World!")
}

func catcher(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "This is a catch all location")
}

func main() {
    fmt.Printf("Server is Running...")

    // mux stands for HTTP Request multiplexer
    // this router matches incoming requests and calls a handler for the route
    // that matches the URL
    mux := http.NewServeMux()
    mux.HandleFunc("/hello", hello)
    mux.HandleFunc("/world", world)
    mux.HandleFunc("/", catcher)

    http.ListenAndServe(":8000", mux)
}
