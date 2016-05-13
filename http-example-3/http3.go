// Run using `go run http3.go`
// We added in a Mux to help with routing

package main

import (
    "fmt"
    "io"  // input output utility functions
    "net/http"  // web server
)

// We have a http.ResponseWriter (and its corresponding stream)
// Because we're using io, we have a response stream
func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello")  // this is a writable Stream
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
    fmt.Printf("Server is Running...")
    server := http.Server{
        Addr: ":8000",
        Handler: &myHandler{},
    }

    mux = make(map[string]func(http.ResponseWriter, *http.Request))
    mux["/"] = hello

    server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    if h, ok := mux[r.URL.String()]; ok {
        h(w, r)
        return
    }

    io.WriteString(w, "My server: "+r.URL.String())
}
