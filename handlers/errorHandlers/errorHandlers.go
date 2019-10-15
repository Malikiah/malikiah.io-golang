package errorhandlers

import (
	"io"
	"log"
	"net/http"
)

func HealthCheckHandler(response http.ResponseWriter, request *http.Request) {
	// A very simple health check.
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(response, `{"alive": true}`)
}

func NotFoundHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusNotFound)
	log.Println("IP Address: " + request.RemoteAddr)
	log.Println("Requested URI: " + request.RequestURI)
	log.Println("Status Code: ", http.StatusNotFound)
	log.Println(request.Header)
}

func MethodNotAllowedHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	log.Println("IP Address: " + request.RemoteAddr)
	log.Println("Requested URI: " + request.RequestURI)
	log.Println("Status Code: ", http.StatusMethodNotAllowed)
	log.Println(request.Header)
}
