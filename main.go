package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

var channel = make(chan os.Signal, 1)
var wait time.Duration

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1>13º Meetup Docker! Valeu!!!</H1>")
	log.Println("Endpoint Hit: homePage")
}

func handleRequests(wait time.Duration) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)

	srv := &http.Server{
		Addr: "0.0.0.0:10000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      myRouter, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	//Channel := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(channel, os.Interrupt)

	// Block until we receive our signal.
	<-channel

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down web server")
	//os.Exit(0)
}

// RunHTTPServer executes the web server to deal with return codes
func RunHTTPServer(wait time.Duration) {
	log.Println("OneDrive Go Client - waiting fo the code...")
	handleRequests(wait)
}

func main() {
	RunHTTPServer(wait)
}
