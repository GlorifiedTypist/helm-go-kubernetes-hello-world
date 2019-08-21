package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	// HTTP	Server listening port.
	Port string `envconfig:"PORT" default:"80"`
}

func main() {
	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Printf("[INFO] Server listening on :%s", env.Port)
	if err := http.ListenAndServe(":"+env.Port, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)); err != nil {
		log.Printf("[ERROR] %s", err)
	}
}
