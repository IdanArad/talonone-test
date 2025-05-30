package main

import _ "github.com/stretchr/testify/assert"

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    fmt.Println("Starting server on port 8080...")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Talon.One!")
    })
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}

