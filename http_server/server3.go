// Implementing graceful shutdown
package http_server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServer3() {
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

	// Till above we create a http server. Now we need to start it and gracefully shutdown it .

	// Channel to listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Run server in a goroutine
	// Why goroutine ? Because server.ListenAndServe() blocks and thus if we don't run it in a goroutine, the execution would never reach the shutdown logic.
	go func() {
		fmt.Println("Server running on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil {
			if err != context.Canceled {
				log.Println("server error:", err)
			}
		}
	}()

	// Wait for signal => Block until signal received
	<-stop
	// Only when signal comes execution reaches here otherwise it is blocked by above line
	// Shutdown server gracefully
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("server shutdown error:", err)
	}

	fmt.Println("Server stopped cleanly")
}
