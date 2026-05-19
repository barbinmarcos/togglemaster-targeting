package main

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "togglemaster-targeting running")
}

func main() {

	http.HandleFunc("/health", health)

	fmt.Println("Server running on :8082")

	http.ListenAndServe(":8082", nil)
}
