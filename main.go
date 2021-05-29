package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/girdharsourabh/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.Flags())

	sm := http.NewServeMux()

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	sm.Handle("/products", handlers.NewProducts(l))
	s := &http.Server{
		Addr:         ":9000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("received terminate , graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
