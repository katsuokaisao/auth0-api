package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/katsuokaisao/auth0-api/handler"
	"github.com/katsuokaisao/auth0-api/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	router := http.NewServeMux()
	router.HandleFunc("/health", handler.HandleHealth)

	// auth_tokenをチェックするミドルウェアを初期化
	ensureValidTokenMiddleware := middleware.EnsureValidToken()

	privateHandler := middleware.BotReject(http.HandlerFunc(handler.HandlePrivate))
	privateHandler = middleware.PanicRecovery(privateHandler)
	privateHandler = ensureValidTokenMiddleware(privateHandler)
	router.Handle("/api/private", privateHandler)

	port := 8080
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	go func() {
		log.Printf("Server listening on http://localhost:%d\n", port)

		// block until server.ListenAndServe returns
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("There was an error with the http server: %v", err)
		}
	}()

	trap(server)
}

func trap(server *http.Server) {
	signals := make(chan os.Signal, 1)

	// SIGINT: Ctrl + C
	// SIGTERM: kill
	// SIGKILL: kill -9
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// block until a signal is received
	sigs := <-signals

	log.Printf("Signal %s received. Shutting down...", sigs)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// block until server.Shutdown returns
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v", err)
	}

	log.Println("Server gracefully shutdown")
}
