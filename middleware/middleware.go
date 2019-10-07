package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println("IP Address: " + r.RemoteAddr)
		log.Println("Requested URI: " + r.RequestURI)
		log.Println("Status Code: ", http.StatusOK)
		log.Println(r.Header)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
