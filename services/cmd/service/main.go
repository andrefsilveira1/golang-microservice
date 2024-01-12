package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("service starting")

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "It is working!")
	})

	addr := ":8080"
	log.Printf("starting HTTP server at '%s'\n", addr)
	http.ListenAndServe(addr, router)

	log.Println("service shutdown")
}
