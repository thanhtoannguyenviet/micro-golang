package main

import (
	"fmt"
	"log"
	"net/http"
)

const wwebPort = "8080"

type Config struct{}

func main() {
	app := Config{}
	log.Printf("Starting broker service on port %s\n", wwebPort)
	srv := &http.Server{
		Addr:    fmt.Sprint(":%s", wwebPort),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
}
