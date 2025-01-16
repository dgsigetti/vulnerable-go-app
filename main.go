package main

import (
    "fmt"
    "net/http"
    "vulnerable-go-app/utils"
)

func handler(w http.ResponseWriter, r *http.Request) {
    data := "example"
    hashedData := utils.Hash(data)
    fmt.Fprintf(w, "Hashed Data: %s", hashedData)
}

func main() {
    http.HandleFunc("/", handler) // Route to handle HTTP requests
    fmt.Println("Server is running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

