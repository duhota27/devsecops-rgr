package main

import (
    "crypto/sha256"
    "fmt"
    "net/http"
)

func hashPassword(password string) string {
    hash := sha256.Sum256([]byte(password))
    return fmt.Sprintf("%x", hash)
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("X-Frame-Options", "DENY")
    fmt.Fprintln(w, "DevSecOps RGR App is running")
}

func main() {
    // имитация "пароля системы"
    password := "admin123"
    hashed := hashPassword(password)

    // вывод в консоль
    fmt.Println("Original password:", password)
    fmt.Println("Hashed password (SHA-256):", hashed)

    http.HandleFunc("/", handler)

    fmt.Println("Server started on :8080")
    http.ListenAndServe(":8080", nil)
}