package middleware

import (
	"fmt"
	"net/http"
)

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(res http.ResponseWriter, req *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("panic occurred:", err)
					res.Header().Set("Content-Type", "application/json")
					res.WriteHeader(http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(res, req)
		},
	)
}
