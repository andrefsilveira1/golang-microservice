package main

import (
	"log"
	"microservices/services/internal/config"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Service starting")
	cfg := config.Config{}
	router := mux.NewRouter().StrictSlash(true)
	server, err := rest.NewServer(cfg.ServerHTTP, router)

	if err != nil {
		log.Fatalf("Error on server: %+v", err)
	}

	server.Start()
	log.Println("Server shutdown")
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "It is working!")
	// })

	// addr := ":8080"
	// log.Printf("starting HTTP server at '%s'\n", addr)
	// http.ListenAndServe(addr, router)
}
