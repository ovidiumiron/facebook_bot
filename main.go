package main

import (
	"log"
	"net/http"
	"os"

	"github.com/omiron/facebook_bot/handlers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.HandleFunc("/webhook", handlers.WebHook)

}
