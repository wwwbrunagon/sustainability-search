package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    fmt.Println("Server starting on port 8000...")
    if err := http.ListenAndServe(":8000", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
