package main

import (
	"context"
	"os"

	"github.com/selesy/go-material/pkg/generator"
	log "github.com/sirupsen/logrus"
)

func main() {
	// err := generator.ReadPackages(context.TODO(), "")
	// if err != nil {
	// 	log.WithError(err).Error("Generator failed")
	// }

	g := generator.Generator{
		Github: generator.DefaultGithub(generator.MDCOwner, generator.MDCRepository),
	}

	tagName, err := g.Github.ReadLatestRelease(context.TODO())
	if err != nil {
		log.WithError(err).Error("Generator failed")
		os.Exit(1)
	}

	sha, err := g.Github.GetTagSha(context.TODO(), tagName)
	if err != nil {
		log.WithError(err).Error("Generator failed")
		os.Exit(1)
	}

	err = g.Github.GetTree(context.TODO(), sha)
	if err != nil {
		log.WithError(err).Error("Generator failed")
		os.Exit(1)
	}

	os.Exit(0)
}
