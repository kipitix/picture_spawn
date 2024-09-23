package main

import (
	"github.com/kipitix/wallpaperswide_deskup/pkg/downloader"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	log.Info("Hello World")

	downloader.ParseWallPages()
}
