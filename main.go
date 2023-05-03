package main

import (
	"fmt"
	"log"
	"net/http"

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
	log.Printf("Server listening on http://localhost:%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), router); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
