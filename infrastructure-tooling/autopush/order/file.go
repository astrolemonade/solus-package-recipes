package order

import (
	"gopkg.in/yaml.v2"
	"os"
)

type File struct {
	Tiers []*Tier `yaml:"tiers"`
	index int
}

type Tier struct {
	Name     string   `yaml:"name"`
	Packages []string `yaml:"dirs"`
	index    int
}

// Marshal will attempt to marshal the file as YAML
func (f *File) Marshal() ([]byte, error) {
	return yaml.Marshal(&f)
}

func Load(path string) (f *File, err error) {
	raw, err := os.Open(path)
	if err != nil {
		return
	}
	defer raw.Close()
	dec := yaml.NewDecoder(raw)
	f = &File{}
	err = dec.Decode(&f)
	return
}

func (f *File) Next() (t *Tier, ok bool) {
	if f.index >= len(f.Tiers) {
		return
	}
	t = f.Tiers[f.index]
	ok = true
	f.index++
	return
}

func (t *Tier) Next() (dir string, ok bool) {
	if t.index >= len(t.Packages) {
		return
	}
	dir = t.Packages[t.index]
	ok = true
	t.index++
	return
}
