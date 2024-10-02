package logger

import (
	"os"

	"github.com/UnnoTed/horizontal"
	"github.com/rs/zerolog/log"
)

func SetupZerolog() {
	log.Logger = log.Output(horizontal.ConsoleWriter{Out: os.Stderr})
}
