package generator

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/go-github/v33/github"
	log "github.com/sirupsen/logrus"
)

const (
	MDCPackageURL = "https://github.com/material-components/material-components-web/tree/v9.0.0/packages"

	MDCOwner      = "material-components"
	MDCRepository = "material-components-web"
	MDCPath       = "packages"
)

func ReadPackages(ctx context.Context, version string) error {
	client := github.NewClient(nil)

	fctn, dctn, rsp, err := client.Repositories.GetContents(ctx, MDCOwner, MDCRepository, MDCPath, nil)
	if err != nil {
		return err
	}

	if rsp.StatusCode != 200 {
		return fmt.Errorf("expected 200 response status code - got %d with message %s", rsp.StatusCode, rsp.Status)
	}

	if fctn != nil {
		return errors.New("file contents should have been nil")
	}

	for _, ctn := range dctn {
		log.Infof("Directory entry name/type: %s/%s", ctn.GetName(), ctn.GetType())
		if ctn.GetType() == "dir" {
			fctn, dctn, rsp, err := client.Repositories.GetContents(ctx, MDCOwner, MDCRepository, ctn.GetPath(), nil)
			if err != nil {
				return err
			}

			if rsp.StatusCode != 200 {
				return fmt.Errorf("expected 200 response status code - got %d with message %s", rsp.StatusCode, rsp.Status)
			}

			if fctn != nil {
				return errors.New("file contents should have been nil")
			}

			for _, ctn := range dctn {
				log.Infof("    Sub-directory entry name/type: %s/%s", ctn.GetName(), ctn.GetType())
			}
		}
		// for _, sctn := range ctn.GetContent() {
		// 	if

		// }
	}

	return nil
}
