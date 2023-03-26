//
// Copyright © 2021 Solus Project <copyright@getsol.us>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package spec

import (
	"dev.getsol.us/source/libypkg.git/spec/internal"
	"dev.getsol.us/source/libypkg.git/spec/v2"
	"dev.getsol.us/source/libypkg.git/spec/v3"
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
)

var (
	// ErrInvalidVersion is returned when the requested YPKG format is unsupported
	ErrInvalidVersion = errors.New("invalid ypkg version specified")
)

// Package is a common interface to all version of the Package YML specification
type Package interface {
	// Load populates a PackageYML version by reading in the contents from a specific filepath
	Load(path string, mode int) error
	// Convert turns a versioned PackageYML into the intermediate internal.PackageYML representation
	Convert() (*internal.PackageYML, error)
	// Modify a versioned PackageYML with the contents of an internal.PackageYML
	Modify(changes internal.PackageYML) error
	// File returns the underlying file record
	File() *os.File
	// Save writes any changes to this PackageSpec to the currently open file descriptor
	Save() error
	// Close close the file descriptor for this PackageSpec
	Close()
}

// Load reads in any supported package.yml from file
func Load(path string) (pkg Package, err error) {
	ypkg, err := DetectFormat(path)
	if err != nil {
		return
	}
	pkg, err = NewPackage(ypkg, nil)
	if err != nil {
		return
	}
	err = pkg.Load(path, os.O_RDWR)
	return
}

// DetectFormat reads a package.yml to check for the version of the format
func DetectFormat(path string) (ypkg int, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	dec := yaml.NewDecoder(f)
	var version struct {
		YPKG string `yaml:"YPKG"`
	}
	if err = dec.Decode(&version); err != nil {
		return
	}
	if len(version.YPKG) == 0 {
		ypkg = 2
		return
	}
	ypkg, err = strconv.Atoi(version.YPKG)
	return
}

// NewPackage creates and empty package of the specified version, if supported
func NewPackage(ypkg int, f *os.File) (pkg Package, err error) {
	switch ypkg {
	case 2:
		pkg = v2.NewPackage(f)
	case 3:
		pkg = v3.NewPackage(f)
	default:
		err = ErrInvalidVersion
	}
	return
}

// Auto generates a new package my inspecting the contents of a list of sources
func Auto(sources []string) (pkg Package, err error) {
	def, err := internal.Auto(sources)
	if err != nil {
		return
	}
	pkg = &v2.PackageYML{}
	if err = pkg.Load("package.yml", os.O_CREATE); err != nil {
		return
	}
	err = pkg.Modify(*def)
	return
}

// Bump increments the release number of a package.yml
func Bump(path string) (pkg Package, err error) {
	if pkg, err = Load(path); err != nil {
		return
	}
	i, err := pkg.Convert()
	if err != nil {
		return
	}
	i.Bump()
	err = pkg.Modify(*i)
	return
}

// Convert a package.yml from any version to another
func Convert(path string, ypkg int) (pkg Package, err error) {
	original, err := Load(path)
	if err != nil {
		return
	}
	i, err := original.Convert()
	if err != nil {
		return
	}
	i.Bump()
	pkg, err = NewPackage(ypkg, original.File())
	if err != nil {
		return
	}
	err = pkg.Modify(*i)
	return
}

// Init creates a new package.yml with the required field pre-populated like a template
func Init(path string) (pkg Package, err error) {
	f, err := os.Create(path)
	if err != nil {
		return
	}
	pkg = v2.NewPackage(f)
	err = pkg.Modify(*internal.Default())
	return
}

// Lint checks for errors and common mistakes in package.yml
func Lint(path string) (pkg Package, err error) {
	if pkg, err = Load(path); err != nil {
		return
	}
	i, err := pkg.Convert()
	if err != nil {
		return
	}
	err = i.Lint()
	return
}

// Update modifies the sources in an existing package.yml and overwrites the existing file
func Update(path, version string, sources []string) (pkg Package, err error) {
	original, err := Load(path)
	if err != nil {
		return
	}
	i, err := original.Convert()
	if err != nil {
		return
	}
	original.Close()
	i.Bump()
	if err = i.Update(version, sources); err != nil {
		return
	}
	pkg = &v2.PackageYML{}
	if err = pkg.Load(path, os.O_CREATE); err != nil {
		return
	}
	err = pkg.Modify(*i)
	return
}
