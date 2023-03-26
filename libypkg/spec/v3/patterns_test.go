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
	"dev.getsol.us/source/libypkg.git/spec/shared/array"
	"dev.getsol.us/source/libypkg.git/spec/shared/constant"
	"gopkg.in/yaml.v3"
	"strings"
	"testing"
)

func TestPatternsMarshalEmpty(t *testing.T) {
	expected := "name: golang\n"
	var out strings.Builder
	enc := yaml.NewEncoder(&out)
	value := struct {
		Name     string        `yaml:"name"`
		Patterns array.ListMap `yaml:",omitempty"`
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

func TestPatternsMarshalSingle(t *testing.T) {
	expected := `name: golang
patterns:
    - /usr/bin/go # comment1
    - /usr/share/golang # comment2
`
	value := struct {
		Name     string        `yaml:"name"`
		Patterns array.ListMap `yaml:",omitempty"`
	}{
		Name: "golang",
		Patterns: array.ListMap{
			constant.DefaultPackage: []*yaml.Node{
				&yaml.Node{
					Kind:        yaml.ScalarNode,
					Value:       "/usr/bin/go",
					LineComment: "# comment1",
				},
				&yaml.Node{
					Kind:        yaml.ScalarNode,
					Value:       "/usr/share/golang",
					LineComment: "# comment2",
				},
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

func TestPatternsMarshalMultiple(t *testing.T) {
	expected := `name: golang
patterns:
    - /usr/bin/go
    - /usr/share/golang
    - devel:
        - /usr/include
        - /usr/share/golang/include
`
	value := struct {
		Name     string        `yaml:"name"`
		Patterns array.ListMap `yaml:",omitempty"`
	}{
		Name: "golang",
		Patterns: array.ListMap{
			constant.DefaultPackage: []*yaml.Node{
				&yaml.Node{
					Kind:  yaml.ScalarNode,
					Value: "/usr/bin/go",
				},
				&yaml.Node{
					Kind:  yaml.ScalarNode,
					Value: "/usr/share/golang",
				},
			},
			"devel": []*yaml.Node{
				&yaml.Node{
					Kind:  yaml.ScalarNode,
					Value: "/usr/include",
				},
				&yaml.Node{
					Kind:  yaml.ScalarNode,
					Value: "/usr/share/golang/include",
				},
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

func TestPatternsUnmarshalEmpty(t *testing.T) {
	input := "name: golang"
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	value := struct {
		Name     string        `yaml:"name"`
		Patterns array.ListMap `yaml:",omitempty"`
	}{
		Patterns: make(array.ListMap),
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if value.Name != "golang" {
		t.Fatalf("expected '%s', found: %s", "golang", value.Name)
	}
}

func TestPatternsUnmarshalSingle(t *testing.T) {
	input := `name: golang
patterns:
    - /usr/bin/go
    - /usr/share/golang
`
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	value := struct {
		Name     string        `yaml:"name"`
		Patterns array.ListMap `yaml:",omitempty"`
	}{
		Patterns: make(array.ListMap),
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if value.Name != "golang" {
		t.Fatalf("expected '%s', found: %s", "golang", value.Name)
	}
	if len(value.Patterns) != 1 {
		t.Fatalf("expected 1 package, found: %d", len(value.Patterns))
	}
	paths := value.Patterns[constant.DefaultPackage]
	if len(paths) != 2 {
		t.Fatalf("expected 2 paths, found: %d", len(paths))
	}
	if paths[0].Value != "/usr/bin/go" {
		t.Errorf("expected '%s', found: %s", "/usr/bin/go", paths[0].Value)
	}
	if paths[1].Value != "/usr/share/golang" {
		t.Errorf("expected '%s', found: %s", "/usr/share/golang", paths[0].Value)
	}
}

func TestPatternsUnmarshalMultiple(t *testing.T) {
	input := `name: golang
patterns:
    - /usr/bin/go
    - /usr/share/golang
    - devel:
        - /usr/include
        - /usr/share/golang/include
`
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	value := struct {
		Name     string        `yaml:"name"`
		Patterns array.ListMap `yaml:",omitempty"`
	}{
		Patterns: make(array.ListMap),
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if value.Name != "golang" {
		t.Fatalf("expected '%s', found: %s", "golang", value.Name)
	}
	if len(value.Patterns) != 2 {
		t.Fatalf("expected 2 package, found: %d", len(value.Patterns))
	}
	paths := value.Patterns[constant.DefaultPackage]
	if len(paths) != 2 {
		t.Fatalf("expected 2 paths, found: %d", len(paths))
	}
	if paths[0].Value != "/usr/bin/go" {
		t.Errorf("expected '%s', found: %s", "/usr/bin/go", paths[0].Value)
	}
	if paths[1].Value != "/usr/share/golang" {
		t.Errorf("expected '%s', found: %s", "/usr/share/golang", paths[0].Value)
	}
	paths = value.Patterns["devel"]
	if len(paths) != 2 {
		t.Fatalf("expected 2 paths, found: %d", len(paths))
	}
	if paths[0].Value != "/usr/include" {
		t.Errorf("expected '%s', found: %s", "/usr/include", paths[0].Value)
	}
	if paths[1].Value != "/usr/share/golang/include" {
		t.Errorf("expected '%s', found: %s", "/usr/share/golang/include", paths[0].Value)
	}
}
