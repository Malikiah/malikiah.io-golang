package errorhandlers

import (
	"io"
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	log.Println("IP Address: " + r.RemoteAddr)
	log.Println("Requested URI: " + r.RequestURI)
	log.Println("Status Code: ", http.StatusNotFound)
	log.Println(r.Header)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	log.Println("IP Address: " + r.RemoteAddr)
	log.Println("Requested URI: " + r.RequestURI)
	log.Println("Status Code: ", http.StatusMethodNotAllowed)
	log.Println(r.Header)
}
