package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"  // web server
}

type Person struct {
	ID		string `json:"id,omitempty"`
}

func main() {
	fmt.Println("vim-go")
}
