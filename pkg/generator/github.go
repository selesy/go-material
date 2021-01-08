package generator

import (
	"context"
	"fmt"

	"github.com/google/go-github/v33/github"

	log "github.com/sirupsen/logrus"
)

const (
	GITHUB_OWNER      = ""
	GITHUB_REPOSITORY = ""
)

type Github struct {
	owner      string
	repository string
	client     *github.Client
}

func DefaultGithub(owner, repository string) *Github {
	return &Github{
		owner:      owner,
		repository: repository,
		client:     github.NewClient(nil),
	}
}

func (g *Github) ReadLatestRelease(ctx context.Context) (string, error) {
	rel, resp, err := g.client.Repositories.GetLatestRelease(ctx, g.owner, g.repository)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("expected 200 response status code - got %d with message %s", resp.StatusCode, resp.Status) // TODO
	}

	tag := rel.GetTagName()
	if tag == "" {
		return "", fmt.Errorf("release's tag name should not be empty")
	}

	log.Infof("Latest release: %s", tag) // TODO

	return tag, nil
}

func (g *Github) ReadRelease(ctx context.Context, version string) error {
	return nil // TODO
}

func (g *Github) GetTagSha(ctx context.Context, tagName string) (string, error) {
	tags, resp, err := g.client.Repositories.ListTags(ctx, g.owner, g.repository, nil)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("expected 200 response status code - got %d with message %s", resp.StatusCode, resp.Status) // TODO
	}

	for _, tag := range tags {
		if tag.GetName() == tagName {
			log.Infof("Tag SHA: %s", tag.GetCommit().GetSHA())
			return tag.GetCommit().GetSHA(), nil
		}
	}

	return "", fmt.Errorf("tag name not found for owner/repository %s/%s", g.owner, g.repository)
}

func (g *Github) GetTree(ctx context.Context, sha string) error {
	tree, resp, err := g.client.Git.GetTree(ctx, g.owner, g.repository, sha, true)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("expected 200 response status code - got %d with message %s", resp.StatusCode, resp.Status) // TODO
	}

	for _, item := range tree.Entries {
		log.Infof("Tree item path (type): %s (%s)", item.GetPath(), item.GetType())

		if item.GetContent() != "" {
			log.Errorf("Tree item contents: %v", item.GetContent())
		}
	}

	return nil // TODO
}
