package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kipitix/picture_spawn/internal/interface/dataserverapi"
	"github.com/rs/zerolog/log"
)

func main() {

	apiServer := dataserverapi.NewServer()

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: apiServer.ServerMux(),
	}

	shutdownChan := make(chan bool, 1)

	go func() {
		if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().AnErr("HTTP server error: %v", err)
		}

		// simulate time to close connections
		time.Sleep(1 * time.Millisecond)

		log.Info().Msg("Stopped serving new connections.")
		shutdownChan <- true
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Fatal().AnErr("HTTP shutdown error: %v", err)
	}

	<-shutdownChan
	log.Info().Msg("Graceful shutdown complete.")
}
