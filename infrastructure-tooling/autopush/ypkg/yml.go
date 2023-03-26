package ypkg

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type PackageYML struct {
	Name    string `yaml:"name"`
	Release int    `yaml:"release"`
}

func Load(dir string) (pkg PackageYML, err error) {
	path := filepath.Join(dir, "package.yml")
	raw, err := os.Open(path)
	if err != nil {
		return
	}
	defer raw.Close()
	dec := yaml.NewDecoder(raw)
	err = dec.Decode(&pkg)
	return
}
