package pkg

import (
	"encoding/xml"
	"github.com/getsolus/libeopkg/pspec"
	"os"
)

// LoadPSpec reads in a PSpec and decodes it from XML
func LoadPSpec(path string) (spec pspec.PSpec, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	dec := xml.NewDecoder(f)
	err = dec.Decode(&spec)
	return
}
