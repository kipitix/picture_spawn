package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	log.Fatal().Err(http.ListenAndServe(":8080", nil))

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
