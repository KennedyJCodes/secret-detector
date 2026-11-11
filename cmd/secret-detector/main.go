package main

import (
	"log"
	"os"

	"github.com/KennedyJCodes/secret-detector/internal/adapter/bufioreader"
	"github.com/KennedyJCodes/secret-detector/internal/app"
	"github.com/KennedyJCodes/secret-detector/internal/config"
)

func main() {
	settings, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	lineReader, err := bufioreader.New(bufioreader.Config{MaxLineBytes: settings.MaxLineBytes})
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(lineReader, os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
