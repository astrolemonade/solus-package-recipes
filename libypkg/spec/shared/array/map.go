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
	"errors"
	"gopkg.in/yaml.v3"
	"sort"
)

// Map is a YAML list or single value that gets read in as a map
//
// Examples:
// component: system.devel
//
// component:
//     - system.devel
//     - docs: programming.tools
type Map map[string]*yaml.Node

// NewMap creates an empty Map
func NewMap() Map {
	return make(Map)
}

// ErrInvalidMap indicates that an Map is either invalid or being filled by invalid YAML
var ErrInvalidMap = errors.New("Map must be a single string or an array of key value pairs")

// MarshalYAML handles custom marshaling for Map
func (am Map) MarshalYAML() (out interface{}, err error) {
	switch len(am) {
	case 0:
		return
	case 1:
		main, ok := am[constant.DefaultPackage]
		if !ok {
			err = ErrInvalidMap
			return
		}
		out = main
	default:
		main, ok := am[constant.DefaultPackage]
		if !ok {
			err = ErrInvalidMap
			return
		}
		nodes := make([]*yaml.Node, 0)
		nodes = append(nodes, main)
		var names []string
		for name := range am {
			if name != constant.DefaultPackage {
				names = append(names, name)
			}
		}
		sort.Strings(names)
		for _, name := range names {
			node := yaml.Node{
				Kind: yaml.MappingNode,
			}
			key := yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: name,
			}
			node.Content = append(node.Content, &key, am[name])
			nodes = append(nodes, &node)
		}
		out = nodes
	}
	return
}

// UnmarshalYAML handles custom unmarshaling for Map
func (am *Map) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		if len(value.Value) == 0 {
			return ErrInvalidMap
		}
		*am = make(Map)
		(*am)[constant.DefaultPackage] = value
	case yaml.SequenceNode:
		if len(value.Content) == 0 {
			return ErrInvalidMap
		}
		main := value.Content[0]
		if main.Kind != yaml.ScalarNode || len(main.Value) == 0 {
			return ErrInvalidMap
		}
		m := make(Map)
		m[constant.DefaultPackage] = main
		for _, node := range value.Content[1:] {
			if node.Kind != yaml.MappingNode || len(node.Content) != 2 {
				return ErrInvalidMap
			}
			k := node.Content[0]
			if k.Kind != yaml.ScalarNode || len(k.Value) == 0 || k.Value == constant.DefaultPackage {
				return ErrInvalidMap
			}
			v := node.Content[1]
			if v.Kind != yaml.ScalarNode || len(v.Value) == 0 {
				return ErrInvalidMap
			}
			m[k.Value] = v
		}
		*am = m
	default:
		return ErrInvalidMap
	}
	return nil
}
