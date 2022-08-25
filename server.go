package main

import (
	"fmt"
	"net/http"
)

func startServer() {
	fmt.Printf("Server started at 127.0.0.1:%d", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
