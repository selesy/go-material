package main

import (
	"flag"
	"os"

	"github.com/selesy/go-material/pkg/generator"
	log "github.com/sirupsen/logrus"
)

func main() {

	var ext = flag.String("ext", "yaml", "Help for ext")
	var in = flag.String("in", "./", "Help for in")
	var out = flag.String("out", "./", "Help for out")
	var pkg = flag.String("pkg", "", "Help for pkg")

	flag.Parse()

	err := (&generator.Config{
		Extension: *ext,
		InDir:     *in,
		OutDir:    *out,
		Package:   *pkg,
	}).Generate()

	if err != nil {
		log.WithError(err).Error("MDC Generator failed")
		os.Exit(1)
	}

	os.Exit(0)
}
