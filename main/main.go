package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    fmt.Println("Starting server on port 8080...")

    // Root handler
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Talon.One!")
    })

    // Health handler for liveness probe
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, "ok")
    })

    // Ready handler for readiness probe
    http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, "ready")
    })

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}

