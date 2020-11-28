package generator

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Extension string
	InDir     string
	OutDir    string
	Package   string
}

func (c *Config) Generate() error {
	if c.Package == "" {
		c.Package = os.Getenv("GOPACKAGE")
	}

	if c.Package == "" {
		return errors.New("ConfigurationError") // TODO: create a typed error for this
	}

	log.Info("Started Material Design Components (MDC) generator")
	log.Info("--------------------------------------------------")
	log.Info("Configuration:")
	log.Info("  Extension: ", c.Extension)
	log.Info("  Input directory: ", c.InDir)
	log.Info("  Output directory: ", c.OutDir)
	log.Info("  Package: ", c.Package)

	_, err := c.read()
	if err != nil {
		return err
	}

	log.Info("Finished Material Design Components (MDC) generator")

	return nil
}
