package generator

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Spec struct {
	Name string

	Block     Block
	Elements  []Spec
	Modifiers []string

	Rippled bool
}

// type Block struct {
// 	Class string
// 	Tag   string
// }

// type Element struct {
// 	Name string
// 	Spec Spec
// }

func (c Config) read() (map[string]Component, error) {
	files, err := c.files()
	if err != nil {
		return nil, err
	}

	log.Info("--------------------------------------------------")
	log.Info("Targeting files:")

	comps := map[string]Component{}

	for _, file := range files {
		log.Info("  ", file)

		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Printf("yamlFile.Get err   #%v ", err) // TODO
			return nil, err                            // TODO
		}

		dec := yaml.NewDecoder(bytes.NewReader(data))

		for {
			c := Component{}
			if dec.Decode(&c) != nil {
				break
			}
			comps[c.Name] = c
		}
	}

	log.Info("--------------------------------------------------")
	log.Info("Imported component specifications:")

	for k, v := range comps {
		log.Info("  ", k, ":")
		log.Info("    ", v)
	}

	return nil, nil
}

func (c Config) files() ([]string, error) {
	in, err := os.Open(c.InDir)
	if err != nil {
		return nil, err
	}

	fis, err := in.Readdir(0)
	if err != nil {
		return nil, err
	}

	f := []string{}
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}

		if strings.HasSuffix(fi.Name(), c.Extension) {
			f = append(f, filepath.Join(c.InDir, fi.Name()))
		}
	}

	return f, nil
}
