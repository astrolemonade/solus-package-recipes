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

package shared

import (
	"gopkg.in/yaml.v3"
	"strings"
	"testing"
)

func TestLicensesMarshalEmpty(t *testing.T) {
	var ls Licenses
	var out strings.Builder
	enc := yaml.NewEncoder(&out)
	if err := enc.Encode(ls); err != ErrNotLicense {
		t.Fatalf("Expected ErrNotLicense, found: %s", err)
	}
	if result := out.String(); len(result) != 0 {
		t.Fatalf("Expected empty result, found: %s", result)
	}
}

func TestLicensesMarshalSingle(t *testing.T) {
	expected := "license: MIT\n"
	var ls Licenses
	l := yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: "MIT",
	}
	ls = append(ls, l)
	value := struct {
		License Licenses `yaml:"license"`
	}{
		License: ls,
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

func TestLicensesMarshalMultiple(t *testing.T) {
	expected := `license:
    - MIT
    - Apache-2.0
`
	var ls Licenses
	l := yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: "MIT",
	}
	ls = append(ls, l)
	l = yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: "Apache-2.0",
	}
	ls = append(ls, l)
	value := struct {
		License Licenses `yaml:"license"`
	}{
		License: ls,
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

func TestLicensesUnmarshalEmpty(t *testing.T) {
	input := "license:"
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		License Licenses `yaml:"license"`
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if l := len(value.License); l != 0 {
		t.Fatalf("expected exactly 0 entries, found: %d", l)
	}
}

func TestLicensesUnmarshalSingle(t *testing.T) {
	input := "license: MIT # hello"
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		License Licenses `yaml:"license"`
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if l := len(value.License); l != 1 {
		t.Fatalf("expected exactly one entry, found: %d", l)
	}
	node := value.License[0]
	if node.Kind != yaml.ScalarNode {
		t.Errorf("exected '%d', found: %d", yaml.ScalarNode, node.Kind)
	}
	if node.Value != "MIT" {
		t.Errorf("expected 'MIT', found: %s", node.Value)
	}
	if node.LineComment != "# hello" {
		t.Errorf("expected '# hello', found: %s", node.LineComment)
	}
}

func TestLicensesUnmarshalMultiple(t *testing.T) {
	input := `license:
     - MIT # hello
     - Apache-2.0
`
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		License Licenses `yaml:"license"`
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if l := len(value.License); l != 2 {
		t.Fatalf("expected exactly two entries, found: %d", l)
	}
	node := value.License[0]
	if node.Kind != yaml.ScalarNode {
		t.Errorf("exected '%d', found: %d", yaml.ScalarNode, node.Kind)
	}
	if node.Value != "MIT" {
		t.Errorf("expected 'MIT', found: %s", node.Value)
	}
	if node.LineComment != "# hello" {
		t.Errorf("expected '# hello', found: %s", node.LineComment)
	}
	node = value.License[1]
	if node.Kind != yaml.ScalarNode {
		t.Errorf("exected '%d', found: %d", yaml.ScalarNode, node.Kind)
	}
	if node.Value != "Apache-2.0" {
		t.Errorf("expected 'Apache-2.0', found: %s", node.Value)
	}
	if node.LineComment != "" {
		t.Errorf("expected '', found: %s", node.LineComment)
	}
}
