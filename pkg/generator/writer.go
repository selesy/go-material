package generator

import (
	"fmt"
	"html/template"
	"os"

	log "github.com/sirupsen/logrus"
)

func write(spec Spec) error {
	_, err := os.Stat(spec.Name)
	if !os.IsNotExist(err) {
		return fmt.Errorf("file already exists: %s", spec.Name)
	}

	tmplt, err := template.New("mdc").Parse(mdcTemplate)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	// file, err := os.OpenFile(spec.Code, os.O_CREATE, os.ModeAppend)
	file, err := os.Create(spec.Name)
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmplt.Execute(file, spec)
	if err != nil {
		return err
	}

	return nil
}
