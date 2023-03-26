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

// Convert translates a v2.BuildFlags to an internal.BuildFlags
func (flags BuildFlags) Convert() internal.BuildFlags {
	return internal.BuildFlags{
		AutoDep:    flags.AutoDep,
		AVX2:       flags.AVX2,
		Clang:      flags.Clang,
		CCache:     flags.CCache,
		Debug:      flags.Debug,
		Devel:      flags.Devel,
		Emul32:     flags.Emul32,
		Extract:    flags.Extract,
		LAStrip:    flags.LAStrip,
		LibSplit:   flags.LibSplit,
		Networking: flags.Networking,
		Optimize:   flags.Optimize,
		Strip:      flags.Strip,
	}
}

// Modify translates an internal.BuildFlags to a v2.BuildFlags
func (flags *BuildFlags) Modify(changes internal.BuildFlags) {
	flags.AutoDep = changes.AutoDep
	flags.AVX2 = changes.AVX2
	flags.Clang = changes.Clang
	flags.CCache = changes.CCache
	flags.Debug = changes.Debug
	flags.Devel = changes.Devel
	flags.Emul32 = changes.Emul32
	flags.Extract = changes.Extract
	flags.LAStrip = changes.LAStrip
	flags.LibSplit = changes.LibSplit
	flags.Networking = changes.Networking
	flags.Optimize = changes.Optimize
	flags.Strip = changes.Strip
}
