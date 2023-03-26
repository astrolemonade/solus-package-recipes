package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// hashFile creates a sha1sum for a given file
func hashFile(path string) error {
	fmt.Fprintf(os.Stdout, "Hashing %s... ", path)
	iFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer iFile.Close()
	h := sha1.New()
	_, err = io.Copy(h, iFile)
	if err != nil {
		return err
	}
	oFile, err := os.Create(path + ".sha1sum")
	if err != nil {
		return err
	}
	defer oFile.Close()
	fmt.Fprintf(oFile, "%x", h.Sum(nil))
	return nil
}

// xzFile compresses a file with xz and keeps the original
func xzFile(path string) error {
	fmt.Fprintf(os.Stdout, "Compressing %s... ", path)
	cmd := exec.Command("xz", "-fkz", path)
	return cmd.Run()
}

// writeIndex writes out the index along with its sha1sum and xz copies
func writeIndex(index []byte) error {
	fmt.Fprintf(os.Stdout, "Writing out %s... ", eopkgIndex)
	err := ioutil.WriteFile(eopkgIndex, index, 0644)
	if err != nil {
		err = fmt.Errorf("Failed to write index, reason: '%s'", err.Error())
		return err
	}
	fmt.Fprintf(os.Stdout, "Done.\n")

	err = xzFile(eopkgIndex)
	if err != nil {
		err = fmt.Errorf("Failed to compress index, reason: '%s'", err.Error())
		return err
	}
	fmt.Fprintf(os.Stdout, "Done.\n")

	err = hashFile(eopkgIndex)
	if err != nil {
		err = fmt.Errorf("Failed to hash index, reason: '%s'", err.Error())
		return err
	}
	fmt.Fprintf(os.Stdout, "Done.\n")

	err = hashFile(eopkgIndex + ".xz")
	if err != nil {
		err = fmt.Errorf("Failed to hash compressed index, reason: '%s'", err.Error())
		return err
	}
	fmt.Fprintf(os.Stdout, "Done.\n")
	return nil
}

// vacuumDeltas removes all trace of Deltas from the index
func vacuumDeltas(index []byte) []byte {
	fmt.Fprintf(os.Stdout, "Vacuuming Deltas... ")
	index = deltaRegex.ReplaceAll(index, []byte(deltaPackages))
	fmt.Fprintf(os.Stdout, "Done.\n")
	return index
}

// fetchIndex gets the index from either a local path or an HTTP URL
func fetchIndex(path string) (index []byte, err error) {
	var iFile io.ReadCloser
	var resp *http.Response
	if strings.HasPrefix(path, "http") {
		fmt.Fprintf(os.Stdout, "Fetching index from %s... ", path)
		resp, err = http.Get(path)
		if err == nil {
			iFile = resp.Body
			defer iFile.Close()
		}
		fmt.Fprintf(os.Stdout, "Done.\n")
	} else {
		iFile, err = os.Open(path)
		defer iFile.Close()
	}
	if err != nil {
		err = fmt.Errorf("Failed fetch index, reason: '%s'", err.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "Reading index file... ")
	index, err = ioutil.ReadAll(iFile)
	if err != nil {
		err = fmt.Errorf("Failed to read all of the index, reason: '%s'", err.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "Done.\n")
	return
}

const (
	eopkgIndex    = "eopkg-index.xml"
	unstableIndex = "http://mirrors.rit.edu/solus/packages/shannon/eopkg-index.xml"
)

var deltaRegex = regexp.MustCompile("(?s)(?:<DeltaPackages>).+?(?:<\\/DeltaPackages>)")

const deltaPackages = "<DeltaPackages>\n            </DeltaPackages>"

func main() {

	// determine index path
	var path string
	switch len(os.Args) {
	case 1:
		path = unstableIndex
	case 2:
		path = os.Args[1]
	default:
		fmt.Fprintf(os.Stderr, "Usage: %s [/path/to/index]\n", os.Args[0])
		os.Exit(1)
	}

	// get and read in the index
	index, err := fetchIndex(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	// get rid of all deltas
	index = vacuumDeltas(index)

	// write out the new index
	err = writeIndex(index)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
