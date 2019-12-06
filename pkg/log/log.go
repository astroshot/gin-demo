package log

import (
	"io"
	"log"
	"os"

	"go.uber.org/zap"
)

var ()

func init() {
	// TODO: config zap log
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
}
