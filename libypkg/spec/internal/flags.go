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
)

// BuildFlags are special options that configure the build process
type BuildFlags struct {
	AutoDep    shared.DefaultTrue  `yaml:"autodep,omitempty"`
	AVX2       shared.DefaultFalse `yaml:"avx2,omitempty"`
	Clang      shared.DefaultTrue  `yaml:"clang,omitempty"`
	CCache     shared.DefaultFalse `yaml:"ccache,omitempty"`
	Debug      shared.DefaultTrue  `yaml:"debug,omitempty"`
	Devel      shared.DefaultFalse `yaml:"devel,omitempty"`
	Emul32     shared.DefaultFalse `yaml:"emul32,omitempty"`
	Extract    shared.DefaultTrue  `yaml:"extract,omitempty"`
	LAStrip    shared.DefaultTrue  `yaml:"lastrip,omitempty"`
	LibSplit   shared.DefaultTrue  `yaml:"libsplit,omitempty"`
	Networking shared.DefaultFalse `yaml:"networking,omitempty"`
	Optimize   []string            `yaml:"optimize,omitempty"`
	Strip      shared.DefaultTrue  `yaml:"strip,omitempty"`
}
