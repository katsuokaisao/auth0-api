package middleware

import (
	"net/http"

	"github.com/mssola/user_agent"
)

func BotReject(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(res http.ResponseWriter, req *http.Request) {
			ua := user_agent.New(req.UserAgent())
			if ua.Bot() {
				res.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(res, req)
		},
	)
}
