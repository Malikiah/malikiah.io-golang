package main

import (
	"context"
	"flag"

	//"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	//Handlers
	errorhandlers "malikiah.io/handlers/errorHandlers"
	userhandlers "malikiah.io/handlers/userHandlers"

	//Middleware
	"malikiah.io/middleware"

	//Services

	//External Packages
	"github.com/gorilla/mux"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	port := "3000"
	// Init Router
	router := mux.NewRouter()

	// route handlers / endpoints
	router.HandleFunc("/health", errorhandlers.HealthCheckHandler).Methods("GET").Name("HealthCheck")
	router.HandleFunc("/login", userhandlers.LoginHandler).Methods("POST").Name("Login")
	router.HandleFunc("/register", userhandlers.RegistrationHandler).Methods("POST").Name("Register")

	router.NotFoundHandler = http.HandlerFunc(errorhandlers.NotFoundHandler)

	router.Use(middleware.LoggingMiddleware)

	log.Println("Gopher army ready and is listening on TCP port " + port + "...")
	// Custom server
	srv := &http.Server{
		Addr: "127.0.0.1:" + port,
		// Prevents Slowloris Attacks
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, //Passing gorilla mux instance
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// Acceptes SIGINT for graceful shutdown.

	signal.Notify(c, os.Interrupt)

	// Block until it receives signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait until the timeout deadline.
	srv.Shutdown(ctx)

	log.Println("shutting down...")
	os.Exit(0)

}
