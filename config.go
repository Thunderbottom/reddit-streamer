package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/jessevdk/go-flags"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
)

var opt options
var parser = flags.NewParser(&opt, flags.Default)

// getConfig reads and returns the configuration file
func getConfig() config {
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		default:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			log.Fatalf("Error loading configuration file: %v", err)
		}
	}

	var cfg = config{}
	var koanf = koanf.New(".")
	if err := koanf.Load(file.Provider(opt.ConfigFile), toml.Parser()); err != nil {
		log.Fatalf("Error loading configuration file: %v", err)
	}
	if err := koanf.Unmarshal("", &cfg); err != nil {
		log.Fatalf("Error loading configuration file: %v", err)
	}

	return cfg
}

func getLogger(debug bool) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	if debug {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	return logger
}
