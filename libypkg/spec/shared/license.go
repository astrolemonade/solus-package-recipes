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

// Licenses is allowed to contain one or more SPDX license identifiers, but may be a scalar or a list in YAML
type Licenses []yaml.Node

// ErrNotLicense indicates that a Licenses is being read from invalid YAML or written when empty
var ErrNotLicense = errors.New("Licenses must be a single string or an array of strings")

// MarshalYAML converts Licenses to a string (singular) or list (plural) when writing YAML
func (l Licenses) MarshalYAML() (out interface{}, err error) {
	switch len(l) {
	case 0:
		err = ErrNotLicense
		return
	case 1:
		if l[0].Kind != yaml.ScalarNode {
			err = ErrNotLicense
			return
		}
		out = l[0]
	case 2:
		for _, node := range l {
			if node.Kind != yaml.ScalarNode {
				err = ErrNotLicense
				return
			}
		}
		out = []yaml.Node(l)
	}
	return
}

// UnmarshalYAML converts a string (singular) or list (plural) to Licenses when reading YAML
func (l *Licenses) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		*l = append(*l, *value)
	case yaml.SequenceNode:
		if len(value.Content) == 0 {
			return ErrNotLicense
		}
		var ls Licenses
		for _, node := range value.Content {
			if node.Kind != yaml.ScalarNode {
				return ErrNotLicense
			}
			ls = append(ls, *node)
		}
		*l = ls
	default:
		return ErrNotLicense
	}
	return nil
}
