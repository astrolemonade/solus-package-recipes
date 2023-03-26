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
	"os"
	"strings"
	"testing"
)

func TestNewSourceInvalid(t *testing.T) {
	_, err := NewSource("bob")
	if err == nil {
		t.Error("expected error for invalid source")
	}
}

func TestNewSourceGit(t *testing.T) {
	git := "git|https://github.com/DataDrake/cuppa:v1.0.1"
	src, err := NewSource(git)
	if err != nil {
		t.Errorf("expected no error, found: %s", err)
	}
	hash, ok := src["git|https://github.com/DataDrake/cuppa"]
	if !ok {
		t.Fatal("missing source")
	}
	if hash != "v1.0.1" {
		t.Fatalf("expected '%s', found: %s", "v1.0.1", hash)
	}
}

func TestNewSourceFile(t *testing.T) {
	wd, _ := os.Getwd()
	file := strings.Join([]string{"file:/", wd, "TESTING", "file.md"}, "/")
	sum := "d17245c4f327262bb7c4d7571a95d71d452bb6073331d7866b289154be6396ba"
	src, err := NewSource(file)
	if err != nil {
		t.Errorf("expected no error, found: %s", err)
	}
	hash, ok := src[file]
	if !ok {
		t.Fatal("missing source")
	}
	if hash != sum {
		t.Fatalf("expected '%s', found: %s", sum, hash)
	}
}

func TestNewSourceHTTP(t *testing.T) {
	url := "https://github.com/DataDrake/cuppa/archive/v1.0.1.tar.gz"
	sum := "97bb4ca8003fcf36075a968bd6bf80f864acac5d26284fb92e3fe6899ad92fd5"
	src, err := NewSource(url)
	if err != nil {
		t.Errorf("expected no error, found: %s", err)
	}
	hash, ok := src[url]
	if !ok {
		t.Fatal("missing source")
	}
	if hash != sum {
		t.Fatalf("expected '%s', found: %s", sum, hash)
	}
}

func TestUpdateSources(t *testing.T) {
	urls := []string{
		"git|https://github.com/DataDrake/cuppa/archive:v1.0",
		"git|https://github.com/DataDrake/cuppa/archive:v1.0.1",
		"git|https://github.com/DataDrake/cuppa/archive:v1.0.2",
	}
	keys := []string{
		"git|https://github.com/DataDrake/cuppa/archive",
		"git|https://github.com/DataDrake/cuppa/archive",
		"git|https://github.com/DataDrake/cuppa/archive",
	}
	hashes := []string{
		"v1.0",
		"v1.0.1",
		"v1.0.2",
	}
	srcs, err := UpdateSources(urls)
	if err != nil {
		t.Errorf("expected no error, found: %s", err)
	}
	for i, src := range srcs {
		hash, ok := src[keys[i]]
		if !ok {
			t.Error("missing source")
			continue
		}
		if sum := hashes[i]; hash != sum {
			t.Errorf("expected '%s', found: %s", sum, hash)
		}
	}
}
