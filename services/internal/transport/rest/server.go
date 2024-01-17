package rest

import (
	"microservices/services/internal/config"
	"net/http"
)

type Server struct {
	*http.Server
	Config *config.ServerHTTP
}
func NewServer(config *config.ServerHTTP, router *mux.Router) *Server {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server Working!")
	}).Methods(http.MethodGet)

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	return &Server {
		Server: &http.Server {
			Addr: addr,
			Handler: router,
			WriteTimeout: time.Second * 15,
			ReadTimeout: time.Second * 15,
			IdleTimeout: time.Second * 60
		},
		Config: config
	}
}