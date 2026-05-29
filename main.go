package main

import (
    "crypto/tls"
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

    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    fmt.Fprintln(w, "DevSecOps RGR App is running")
}

func main() {
    password := "admin123"
    hashed := hashPassword(password)

    fmt.Println("Original password:", password)
    fmt.Println("Hashed password (SHA-256):", hashed)

    http.HandleFunc("/", handler)

    fmt.Println("Server started on https://localhost:8443")

    server := &http.Server{
        Addr: ":8443",
        TLSConfig: &tls.Config{
            MinVersion: tls.VersionTLS12,
        },
    }

    err := server.ListenAndServeTLS("cert.pem", "key.pem")
    if err != nil {
        fmt.Println("HTTPS error:", err)
    }
}