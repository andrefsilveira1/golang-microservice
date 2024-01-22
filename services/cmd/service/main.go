package main

import (
	"context"
	"flag"
	"log"
	"microservices/services/internal/config"
	"microservices/services/internal/domain"
	repository "microservices/services/internal/repository/postgres"
	"microservices/services/internal/transport/rest"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"golang.org/x/sync/errgroup"
)

func main() {
	log.Println("Service starting")

	var configPath string
	flag.StringVar(&configPath, "config", "", "...")
	flag.Parse()

	cfg := loadConfig(configPath)
	db := loadDatabase(cfg.Database)

	// Repositories
	itemRepository := repository.NewItemRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)

	itemService := domain.NewItemService(itemRepository, categoryRepository)
	categoryService := domain.NewCategoryService(categoryRepository)

	// Shutdown
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

		rest.NewItemHandler(itemService).Register(router)
		rest.NewCategoryHandler(categoryService).Register(router)
		restServer, err = rest.NewServer(cfg.Server.HTTP, router)
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

func loadConfig(path string) *config.Config {
	if path == "" {
		path = os.Getenv("APP_CONFIG_PATH")
		if path == "" {
			path = "./config.yaml"
		}
	}

	cfg, err := config.NewConfig(path)
	if err != nil {
		log.Printf("Configuration error: %v", err)
		os.Exit(-1)
	}

	return cfg
}

func loadDatabase(cfg *config.Database) *sqlx.DB {
	db, err := database.Connect(cfg)
	if err != nil {
		log.Printf("Database error: %v", err)
		os.Exit(-1)
	}

	return db
}
