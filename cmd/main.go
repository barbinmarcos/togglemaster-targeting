package main

import (
	"fmt"
	"net/http"
	"time"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "togglemaster-targeting running")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/health", health)

	server := &http.Server{
		Addr:         ":8082",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	fmt.Println("Server running on :8082")

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
