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

// ListMap is a YAML list which gets read into a map
//
// Example:
//
// - one
// - two
// - red:
//     - three
//     - four
type ListMap map[string][]*yaml.Node

// NewListMap returns an empty ListMap
func NewListMap() ListMap {
	return make(ListMap)
}

// ErrInvalidListMap indicates that the specified YAML is invalid for this type
var ErrInvalidListMap = errors.New("ListMap must be a list of strings or a map of lists of strings")

// MarshalYAML is a custom marshaler to handle this type
func (am ListMap) MarshalYAML() (out interface{}, err error) {
	if len(am) == 0 {
		return
	}
	nodes := make([]*yaml.Node, 0)
	main := am[constant.DefaultPackage]
	if len(main) > 0 {
		nodes = append(nodes, main...)
	}
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
		value := yaml.Node{
			Kind: yaml.SequenceNode,
		}
		for _, e := range am[name] {
			value.Content = append(value.Content, e)
		}
		node.Content = append(node.Content, &key, &value)
		nodes = append(nodes, &node)
	}
	out = nodes
	return
}

// UnmarshalYAML is a custom unmarshaler to handle this type
func (am *ListMap) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.SequenceNode {
		return ErrInvalidListMap
	}
	if len(value.Content) == 0 {
		return ErrInvalidListMap
	}
	for _, node := range value.Content {
		switch node.Kind {
		case yaml.ScalarNode:
			(*am)[constant.DefaultPackage] = append((*am)[constant.DefaultPackage], node)
		case yaml.MappingNode:
			if len(node.Content) != 2 {
				return ErrInvalidListMap
			}
			k := node.Content[0]
			if k.Kind != yaml.ScalarNode || len(k.Value) == 0 || k.Value == constant.DefaultPackage {
				return ErrInvalidListMap
			}
			v := node.Content[1]
			if v.Kind != yaml.SequenceNode || len(v.Content) == 0 {
				return ErrInvalidListMap
			}
			for _, n := range v.Content {
				if n.Kind != yaml.ScalarNode || len(n.Value) == 0 {
					return ErrInvalidListMap
				}
				(*am)[k.Value] = append((*am)[k.Value], n)
			}
		default:
			return ErrInvalidListMap
		}
	}
	return nil
}
