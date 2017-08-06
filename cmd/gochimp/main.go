package main

import (
	"flag"
	"github.com/kendellfab/gochimp/internal/config"
	"github.com/BurntSushi/toml"
	"log"
	"github.com/kendellfab/gochimp/internal/handlers"
	"github.com/kendellfab/gochimp/internal/stats"
	"net/http"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "c", "config.toml", "Set the path for the config file.")
	flag.Parse()

	var c config.Config
	_, err := toml.DecodeFile(configPath, &c)
	// Log fatal as we need the config info to continue.
	if err != nil {
		log.Fatal("Error reading config file.", err)
	}
	log.Println("Starting gochimp server.")
	statsRepo := stats.NewRepo(c.ApiKey, c.ListId)

	corsHandler := handlers.CorsHandler{Origin: c.Origin}
	statsHandler := stats.NewHandler(statsRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("/member-count", corsHandler.CorsMiddleware(statsHandler.HandleListCount))

	srv := &http.Server {
		Addr: ":8080",
		Handler: mux,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Println("Error with server:", err)
	}
}
