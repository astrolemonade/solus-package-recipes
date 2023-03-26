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

package v3

import (
	"dev.getsol.us/source/libypkg.git/spec/internal"
	"dev.getsol.us/source/libypkg.git/spec/shared"
	"dev.getsol.us/source/libypkg.git/spec/shared/array"
	"gopkg.in/yaml.v3"
	"os"
)

// PackageYML is the v3 representation of the Package YML specification
type PackageYML struct {
	YPKG         int             `yaml:"YPKG"`
	Name         string          `yaml:"name"`
	Version      string          `yaml:"version"`
	Release      uint            `yaml:"release"`
	Source       []shared.Source `yaml:"source"`
	Homepage     string          `yaml:"homepage,omitempty"`
	License      shared.Licenses `yaml:"license"`
	Component    string          `yaml:"component,omitempty"`
	Components   array.Map       `yaml:"components,omitempty"`
	Summary      string          `yaml:"summary,omitempty"`
	Summaries    array.Map       `yaml:"summaries,omitempty"`
	Description  string          `yaml:"description,omitempty"`
	Descriptions array.Map       `yaml:"descriptions,omitempty"`
	Dependencies PackageDeps     `yaml:"deps,omitempty"`
	Flags        BuildFlags      `yaml:"flags,omitempty"`
	Environment  string          `yaml:"environment,omitempty"`
	Stages       BuildStages     `yaml:",inline"`
	Permanent    array.ListMap   `yaml:"permanent,omitempty"`
	Patterns     array.ListMap   `yaml:"patterns,omitempty"`
	f            *os.File
}

// NewPackage creates a new Package with an optional file argument
func NewPackage(f *os.File) *PackageYML {
	return &PackageYML{
		f:            f,
		Components:   array.NewMap(),
		Summaries:    array.NewMap(),
		Descriptions: array.NewMap(),
		Dependencies: NewPackageDeps(),
		Permanent:    array.NewListMap(),
		Patterns:     array.NewListMap(),
	}
}

// Convert translates a v3.PackageYML to an internal.Package.YML
func (p *PackageYML) Convert() (pkg *internal.PackageYML, err error) {
	pkg = &internal.PackageYML{
		YPKG:         p.YPKG,
		Name:         p.Name,
		Version:      p.Version,
		Release:      p.Release,
		Source:       p.Source,
		Homepage:     p.Homepage,
		License:      p.License,
		Component:    p.Component,
		Components:   p.Components,
		Summary:      p.Summary,
		Summaries:    p.Summaries,
		Description:  p.Description,
		Descriptions: p.Descriptions,
		Dependencies: p.Dependencies.Convert(),
		Flags:        p.Flags.Convert(),
		Environment:  p.Environment,
		Stages:       p.Stages.Convert(),
		Permanent:    p.Permanent,
		Patterns:     p.Patterns,
	}
	return
}

// Modify converts an internal.PackageYML to a v2.PackageYML
func (p *PackageYML) Modify(pkg internal.PackageYML) error {
	p.YPKG = 3
	p.Name = pkg.Name
	p.Version = pkg.Version
	p.Release = pkg.Release
	p.Source = pkg.Source
	p.Homepage = pkg.Homepage
	p.License = pkg.License
	p.Component = pkg.Component
	p.Components = pkg.Components
	p.Summary = pkg.Summary
	p.Summaries = pkg.Summaries
	p.Description = pkg.Description
	p.Descriptions = pkg.Descriptions
	p.Dependencies.Modify(pkg.Dependencies)
	p.Flags.Modify(pkg.Flags)
	p.Environment = pkg.Environment
	p.Stages.Modify(pkg.Stages)
	p.Permanent = pkg.Permanent
	p.Patterns = pkg.Patterns
	return nil
}

// Load populates a v2 PackageYML by reading in the contents from a specific filepath
func (p *PackageYML) Load(path string, mode int) error {
	f, err := os.OpenFile(path, mode, 00644)
	if err != nil {
		return err
	}
	p.f = f
	dec := yaml.NewDecoder(p.f)
	return dec.Decode(p)
}

// File returns a pointer to the underlying file record
func (p *PackageYML) File() *os.File {
	return p.f
}

// Save writes any changes to this PackageYML to the currently open file descriptor
func (p *PackageYML) Save() error {
	// Seek to the beginning of the file
	_, err := p.f.Seek(0, 0)
	if err != nil {
		return err
	}
	// clear the contents of the file
	err = p.f.Truncate(0)
	if err != nil {
		return err
	}
	// write out the new contents to the file
	enc := yaml.NewEncoder(p.f)
	return enc.Encode(&p)
}

// Close close the file descriptor for this PackageYML
func (p *PackageYML) Close() {
	_ = p.f.Close()
}
