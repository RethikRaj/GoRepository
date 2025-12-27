package http_server

import (
	"fmt"
	"net/http"
)

func StartServer2() {
	// 1. Create a new router (ServeMux)
	mux := http.NewServeMux()

	// 2. Register route + handler
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong2")
	})

	// 3. Create HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server running on http://localhost:8080")
	server.ListenAndServe()
}
