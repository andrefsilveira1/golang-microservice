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

func (s *Server) Start() error {
	var err error
	log.Printf("HTTP server starting at: '%s: %d \n' ", s.Config.Host, s.Config.Port)
	if s.Config.UseHTTPS {
		log.Println("SSL certificate Enabled")
		path := s.Config.CerPath
		err = s.Server.ListenAndServeTLS(
			fmt.Sprintf("%s /server.crt", path),
			fmt.Sprintf("%s /server.key", path),
		)
	} else {
		log.Println("SSL certificate Disabled")
		err = s.Server.ListenAndServe()
	}

	if err != nil && err != http.ErrServerClosed {
		log.Printf("Unable to start HTTP server: %+v\n", err)
		return err
	}

	return nil
}