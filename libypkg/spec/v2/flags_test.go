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
	"dev.getsol.us/source/libypkg.git/spec/shared"
	"gopkg.in/yaml.v3"
	"strings"
	"testing"
)

func TestBuildFlagsMarshalEmpty(t *testing.T) {
	expected := "name: golang\n"
	var out strings.Builder
	enc := yaml.NewEncoder(&out)
	value := struct {
		Name  string     `yaml:"name"`
		Flags BuildFlags `yaml:",omitempty,inline"`
	}{
		Name: "golang",
	}
	if err := enc.Encode(value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if result := out.String(); result != expected {
		t.Fatalf("Expected '%s' result, found: %s", expected, result)
	}
}

func TestBuildFlagsMarshalSingle(t *testing.T) {
	expected := `name: golang
autodep: no
`
	value := struct {
		Name  string     `yaml:"name"`
		Flags BuildFlags `yaml:",omitempty,inline"`
	}{
		Name: "golang",
		Flags: BuildFlags{
			AutoDep: shared.DefaultTrue{
				Valid: true,
				Bool:  false,
			},
		},
	}
	var out strings.Builder
	enc := yaml.NewEncoder(&out)
	if err := enc.Encode(value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if result := out.String(); result != expected {
		t.Fatalf("Expected %s, found: %s", expected, result)
	}
}

func TestBuildFlagsMarshalMultiple(t *testing.T) {
	expected := `name: golang
autodep: no
avx2: yes
`
	value := struct {
		Name  string     `yaml:"name"`
		Flags BuildFlags `yaml:",omitempty,inline"`
	}{
		Name: "golang",
		Flags: BuildFlags{
			AutoDep: shared.DefaultTrue{
				Valid: true,
				Bool:  false,
			},
			AVX2: shared.DefaultFalse{
				Valid: true,
				Bool:  true,
			},
		},
	}
	var out strings.Builder
	enc := yaml.NewEncoder(&out)
	if err := enc.Encode(value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if result := out.String(); result != expected {
		t.Fatalf("Expected %s, found: %s", expected, result)
	}
}

func TestBuildFlagsUnmarshalEmpty(t *testing.T) {
	input := "name: golang"
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		Name  string     `yaml:"name"`
		Flags BuildFlags `yaml:",omitempty,inline"`
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if value.Name != "golang" {
		t.Fatalf("expected '%s', found: %s", "golang", value.Name)
	}
}

func TestBuildFlagsUnmarshalSingle(t *testing.T) {
	input := `name: golang
autodep: no
`
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		Name  string     `yaml:"name"`
		Flags BuildFlags `yaml:",inline"`
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if value.Name != "golang" {
		t.Fatalf("expected '%s', found: %s", "golang", value.Name)
	}
	autodep := value.Flags.AutoDep
	if !autodep.Valid {
		t.Error("autodep should be valid!")
	}
	if autodep.Bool {
		t.Error("autodep should be false!")
	}
}

func TestBuildFlagsUnmarshalMultiple(t *testing.T) {
	input := `name: golang
autodep: no
avx2: yes
clang: yes
ccache: no
`
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		Name  string     `yaml:"name"`
		Flags BuildFlags `yaml:",inline"`
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if value.Name != "golang" {
		t.Fatalf("expected '%s', found: %s", "golang", value.Name)
	}
	autodep := value.Flags.AutoDep
	if !autodep.Valid {
		t.Error("autodep should be valid!")
	}
	if autodep.Bool {
		t.Error("autodep should be false!")
	}
	avx2 := value.Flags.AVX2
	if !avx2.Valid {
		t.Error("avx2 should be valid!")
	}
	if !avx2.Bool {
		t.Error("avx2 should be true!")
	}
	clang := value.Flags.Clang
	if !clang.Valid {
		t.Error("clang should be valid!")
	}
	if !clang.Bool {
		t.Error("clang should be true!")
	}
	ccache := value.Flags.CCache
	if !ccache.Valid {
		t.Error("ccache should be valid!")
	}
	if ccache.Bool {
		t.Error("ccache should be false!")
	}
	debug := value.Flags.Debug
	if debug.Valid {
		t.Error("debug should not be valid!")
	}
	devel := value.Flags.Devel
	if devel.Valid {
		t.Error("devel should not be valid!")
	}
}
