package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Go app!")
    })

    // Создаём кастомный сервер с таймаутами
    srv := &http.Server{
        Addr:              ":8080",
        ReadHeaderTimeout: 15 * time.Second,
        ReadTimeout:       15 * time.Second,
        WriteTimeout:      10 * time.Second,
        IdleTimeout:       30 * time.Second,
    }

    log.Println("Starting server on :8080")
    if err := srv.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}
