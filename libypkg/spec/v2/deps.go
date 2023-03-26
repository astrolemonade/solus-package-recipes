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

package v2

import (
	"dev.getsol.us/source/libypkg.git/spec/internal"
	"dev.getsol.us/source/libypkg.git/spec/shared/array"
	"gopkg.in/yaml.v3"
)

// PackageDeps includes the dependencies required for the Build, Check, and Run Stages
type PackageDeps struct {
	Replaces  array.ListMap `yaml:"replaces,omitempty"`
	Conflicts array.ListMap `yaml:"conflicts,omitempty"`
	Build     []yaml.Node   `yaml:"builddeps,omitempty"`
	Run       array.ListMap `yaml:"rundeps,omitempty"`
}

// NewPackageDeps retuns an empty PackageDeps
func NewPackageDeps() PackageDeps {
	return PackageDeps{
		Replaces:  array.NewListMap(),
		Conflicts: array.NewListMap(),
		Run:       array.NewListMap(),
	}
}

// Convert translates a v2.PackageDeps to an internal.PackageDeps
func (deps PackageDeps) Convert() internal.PackageDeps {
	return internal.PackageDeps{
		Replaces:  deps.Replaces,
		Conflicts: deps.Conflicts,
		Build:     deps.Build,
		Run:       deps.Run,
	}
}

// Modify translates an internal.PackageDeps to a v2.PackageDeps
func (deps *PackageDeps) Modify(changes internal.PackageDeps) {
	deps.Replaces = changes.Replaces
	deps.Conflicts = changes.Conflicts
	deps.Build = append(changes.Build, changes.Check...)
	deps.Run = changes.Run
}
