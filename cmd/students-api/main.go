package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/aashishdubey1/students-api/internal"
)

func main() {
	// laod config
	cfg := config.MustLoad()

	// router setup
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World")) // by byte ?
	})

	// server setup

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Printf("server Started at %s", cfg.HttpServer.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}

	// db setup

}
