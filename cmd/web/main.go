package main

import (
	"flag"
	"log"
	"net/http"
)

type config struct {
	addr      string
	staticDir string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Directory containing static assets")
	flag.Parse()

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(cfg.staticDir))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on port %s\n", cfg.addr)
	log.Fatal(http.ListenAndServe(cfg.addr, mux))
}
