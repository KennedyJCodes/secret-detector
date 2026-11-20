package main

import (
	"log"
	"os"

	"github.com/KennedyJCodes/secret-detector/internal/adapter/bufioreader"
	"github.com/KennedyJCodes/secret-detector/internal/app"
	"github.com/KennedyJCodes/secret-detector/internal/config"
	"github.com/KennedyJCodes/secret-detector/internal/rules"
	"github.com/KennedyJCodes/secret-detector/internal/services/aws"
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

	ruleSet := rules.NewRegistry()
	if err := ruleSet.Add(aws.AWSAccessKeyRule()); err != nil {
		log.Fatal(err)
	}
	if err := ruleSet.Add(aws.AWSSecretAccessKeyRule()); err != nil {
		log.Fatal(err)
	}

	if err := app.Run(lineReader, os.Stdin, os.Stdout, ruleSet); err != nil {
		log.Fatal(err)
	}
}
