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

package internal

import (
	"dev.getsol.us/source/libypkg.git/spec/shared"
	"dev.getsol.us/source/libypkg.git/spec/shared/array"
	"errors"
	"gopkg.in/yaml.v3"
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
	Component    string          `yaml:"component"`
	Components   array.Map       `yaml:"components"`
	Summary      string          `yaml:"summary"`
	Summaries    array.Map       `yaml:"summaries"`
	Description  string          `yaml:"description"`
	Descriptions array.Map       `yaml:"descriptions"`
	Dependencies PackageDeps     `yaml:"deps,omitempty"`
	Flags        BuildFlags      `yaml:"flags,omitempty"`
	Environment  string          `yaml:"environment,omitempty"`
	Stages       BuildStages     `yaml:",inline"`
	Permanent    array.ListMap   `yaml:"permanent,omitempty"`
	Patterns     array.ListMap   `yaml:"patterns,omitempty"`
}

// NewPackage returns an empty package
func NewPackage() *PackageYML {
	return &PackageYML{
		Components:   array.NewMap(),
		Summaries:    array.NewMap(),
		Descriptions: array.NewMap(),
		Dependencies: NewPackageDeps(),
		Permanent:    array.NewListMap(),
		Patterns:     array.NewListMap(),
	}
}

// Default provides a skeleton for a new package
func Default() *PackageYML {
	return &PackageYML{
		Name:    "Name-Of-Package",
		Version: "1.0.0a",
		Release: 1,
		Source: []shared.Source{
			shared.Source{
				"URI": "HASH",
			},
		},
		License: shared.Licenses{
			yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: "GPL-2.0-or-later # CHECK AND/OR CHANGE ME",
			},
		},
		Component:   "# SET ME",
		Summary:     "# Short description of the package",
		Description: "# Long description of the package\n# Can be multiple lines\n# Just no 80/100 column wrapping!\n",
		Stages: BuildStages{
			Install: "# do some things\n# and stuff\n",
		},
	}
}

// Bump increments the Release number
func (pkg *PackageYML) Bump() {
	pkg.Release++
}

// Update replaces the existing source with newer ones
func (pkg *PackageYML) Update(version string, sources []string) error {
	pkg.Version = version
	return errors.New("Not yet implemented")
}

// Lint checks over the package for any obvious errors or questionable choices
func (pkg *PackageYML) Lint() error {
	return errors.New("Not yet implemented")
}

// Auto creates a new package from a list of sources
func Auto(sources []string) (pkg *PackageYML, err error) {
	pkg = Default()
	if err = pkg.Update(pkg.Version, sources); err != nil {
		return
	}
	// Inspect the first source to fill out as many fields as possible
	err = errors.New("Not yet implemented")
	return
}
