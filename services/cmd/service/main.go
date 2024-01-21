package main

import (
	"context"
	"log"
	"microservices/services/internal/config"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
)

func main() {
	log.Println("Service starting")
	cfg := &config.ServerHTTP{Host: "localhost", Port: 8080}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	var restServer *rest.Server
	g.Go(func() (err error) {
		router := mux.NewRouter().StrictSlash(true)
		restServer, err = rest.NewServer(cfg, router)
		if err != nil {
			return err
		}

		return restServer.Start()
	})

	log.Println("Service started")
	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	log.Println("Shutdown signal received")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if restServer != nil {
		restServer.Stop(shutdownCtx)
	}

	err := g.Wait()
	if err != nil {
		log.Printf("Server shutdown returned an error")
		defer os.Exit(2)
	}

	log.Println("Service shutdown")

	// router := mux.NewRouter().StrictSlash(true)
	// server, err := rest.NewServer(cfg, router)

	// if err != nil {
	// 	log.Fatalf("Error on server: %+v", err)
	// }

	// server.Start()
	// log.Println("Server shutdown")
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "It is working!")
	// })

	// addr := ":8080"
	// log.Printf("starting HTTP server at '%s'\n", addr)
	// http.ListenAndServe(addr, router)
}
