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

package array

import (
	"dev.getsol.us/source/libypkg.git/spec/shared/constant"
	"gopkg.in/yaml.v3"
	"strings"
	"testing"
)

func TestMapMarshalEmpty(t *testing.T) {
	var am Map
	var out strings.Builder
	enc := yaml.NewEncoder(&out)
	if err := enc.Encode(am); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if result := out.String(); result != "null\n" {
		t.Fatalf("Expected null result, found: %s", result)
	}
}

func TestMapMarshalSingle(t *testing.T) {
	expected := "component: system.devel # comment\n"
	value := struct {
		Components Map `yaml:"component"`
	}{
		Components: Map{
			constant.DefaultPackage: &yaml.Node{
				Kind:        yaml.ScalarNode,
				Value:       "system.devel",
				LineComment: "# comment",
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

func TestMapMarshalMultiple(t *testing.T) {
	expected := `component:
    - system.devel
    - devel: programming.devel
`
	value := struct {
		Components Map `yaml:"component"`
	}{
		Components: Map{
			constant.DefaultPackage: &yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: "system.devel",
			},
			"devel": &yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: "programming.devel",
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

func TestMapUnmarshalEmpty(t *testing.T) {
	input := "component:"
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		Components Map `yaml:"component"`
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if l := len(value.Components); l != 0 {
		t.Fatalf("expected exactly 0 entries, found: %d", l)
	}
}

func TestMapUnmarshalSingle(t *testing.T) {
	input := "component: system.devel # comment"
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		Components Map `yaml:"component"`
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if l := len(value.Components); l != 1 {
		t.Fatalf("expected exactly one entry, found: %d", l)
	}
	v := value.Components[constant.DefaultPackage]
	if v.Value != "system.devel" {
		t.Errorf("expected 'system.devel', found: %s", v.Value)
	}
	if v.LineComment != "# comment" {
		t.Errorf("expected 'system.devel', found: %s", v.Value)
	}
}

func TestMapUnmarshalMultiple(t *testing.T) {
	input := `component:
     - system.devel
     - devel: programming.devel
`
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		Components Map `yaml:"component"`
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if l := len(value.Components); l != 2 {
		t.Fatalf("expected exactly two entries, found: %d", l)
	}
	if v := value.Components[constant.DefaultPackage].Value; v != "system.devel" {
		t.Errorf("expected 'system.devel', found: %s", v)
	}
	if v := value.Components["devel"].Value; v != "programming.devel" {
		t.Errorf("expected 'programming.devel', found: %s", v)
	}
}
