package eopkg

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"github.com/ulikunitz/xz"
	"io"
	"io/ioutil"
	"net/http"
)

type Index struct {
	XMLName  xml.Name  `xml:"PISI"`
	Packages []Package `xml:"Package"`
}

var (
	ErrInvalidIndex  = errors.New("index hash mismatch")
	ErrInvalidSHASum = errors.New("SHA sum invalid length")
)

const (
	IndexURL = "https://mirrors.rit.edu/solus/packages/unstable/eopkg-index.xml.xz"
	IndexSHA = "https://mirrors.rit.edu/solus/packages/unstable/eopkg-index.xml.xz.sha1sum"
)

func FetchIndex() (i *Index, err error) {
	sum, err := FetchSHA()
	if err != nil {
		return
	}
	h := sha1.New()
	resp, err := http.Get(IndexURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	tee := io.TeeReader(resp.Body, h)
	raw, err := xz.NewReader(tee)
	if err != nil {
		return
	}
	dec := xml.NewDecoder(raw)
	i = &Index{}
	if err = dec.Decode(i); err != nil {
		return
	}
	if !bytes.Equal(sum, h.Sum(nil)) {
		err = ErrInvalidIndex
	}
	return
}

func FetchSHA() (sum []byte, err error) {
	resp, err := http.Get(IndexSHA)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	sum = make([]byte, hex.DecodedLen(len(raw)))
	if len(sum) != sha1.Size {
		err = ErrInvalidSHASum
	}
	if _, err = hex.Decode(sum, raw); err != nil {
		return
	}
	return
}

func (i Index) HasRelease(pkg string, rel int) bool {
	for _, p := range i.Packages {
		if p.Name != pkg {
			continue
		}
		return p.HasRelease(rel)
	}
	return false
}

type Package struct {
	Name    string   `xml:"Name"`
	History []Update `xml:"History>Update"`
}

func (p Package) HasRelease(rel int) bool {
	for _, u := range p.History {
		switch {
		case u.Release < rel:
			continue
		case u.Release > rel:
			return false
		default:
			return true
		}
	}
	return false
}

type Update struct {
	Release int `xml:"release,attr"`
}
