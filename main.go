package main

import (
    "fmt"
    "net/http"
)

func main() {
    // create a new http.ServeMux
    mux := http.NewServeMux()

    // create a new http.Server struct
    // use the new "ServeMux" as the server's handler
    // set the .Addr field to ":8080"
    server := &http.Server {
        Addr: ":8080",
        Handler: mux,
    }

    // use the server's ListenAndServe method to start the server
    fmt.Println("Starting server")
    if err := server.ListenAndServe(); err != nil {
        fmt.Errorf("Error in ListeAndServe(): %w", err)
    }
}
