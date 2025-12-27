package http_server

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	fmt.Println("Server running on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
