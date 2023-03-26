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
	"errors"
	"gopkg.in/yaml.v3"
)

var (
	// ErrNotABool indicates a string value that cannot be converted to a bool
	ErrNotABool = errors.New("Not a valid boolean string")
)

// DefaultTrue is a boolean which should omitted if true when marshaling
type DefaultTrue struct {
	Valid bool
	Bool  bool
}

// MarshalYAML writes a DefaultTrue as "no" if false and omits it if empty or invalid
func (dt DefaultTrue) MarshalYAML() (out interface{}, err error) {
	node := yaml.Node{
		Kind: yaml.ScalarNode,
	}
	if dt.Valid && !dt.Bool {
		node.Value = "no"
	}
	out = node
	return
}

// UnmarshalYAML converts a string to a DefaultTrue based on its value
func (dt *DefaultTrue) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return ErrNotABool
	}
	switch value.Value {
	case "no", "NO", "No", "False", "false":
		(*dt).Valid = true
		(*dt).Bool = false
	case "yes", "YES", "Yes", "True", "true":
		(*dt).Valid = true
		(*dt).Bool = true
	default:
		return ErrNotABool
	}
	return nil
}

// DefaultFalse is a boolean which should omitted if false when marshaling
type DefaultFalse struct {
	Valid bool
	Bool  bool
}

// MarshalYAML writes a DefaultFalse as "yes" if true and omits it if empty or invalid
func (df DefaultFalse) MarshalYAML() (out interface{}, err error) {
	node := yaml.Node{
		Kind: yaml.ScalarNode,
	}
	if df.Valid && df.Bool {
		node.Value = "yes"
	}
	out = node
	return
}

// UnmarshalYAML converts a string to a DefaultFalse based on its value
func (df *DefaultFalse) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return ErrNotABool
	}
	switch value.Value {
	case "no", "NO", "No", "False", "false":
		(*df).Valid = true
		(*df).Bool = false
	case "yes", "YES", "Yes", "True", "true":
		(*df).Valid = true
		(*df).Bool = true
	default:
		return ErrNotABool
	}
	return nil
}
