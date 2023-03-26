package eopkg

import (
	"testing"
)

func TestIndexFetch(t *testing.T) {
	i, err := FetchIndex()
	if err != nil {
		t.Errorf("Failed to fetch index, reason: %s", err)
	}
	if i == nil {
		t.Errorf("nil index")
		t.FailNow()
	}
	if len(i.Packages) == 0 {
		t.Error("Index has no packages")
	}
}

func BenchmarkIndexFetch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := FetchIndex()
		if err != nil {
			b.Errorf("Failed to fetch index, reason: %s", err)
		}
	}
}

func TestHasReleasae(t *testing.T) {
	i, err := FetchIndex()
	if err != nil {
		t.Errorf("Failed to fetch index, reason: %s", err)
	}
	if i == nil {
		t.Errorf("nil index")
		t.FailNow()
	}
	if len(i.Packages) == 0 {
		t.Error("Index has no packages")
	}
	if !i.HasRelease("nano", 123) {
		t.Errorf("Missing %s release %d", "nano", 123)
	}
	if i.HasRelease("nano", 5000) {
		t.Errorf("Should not have %s release %d", "nano", 130)
	}
	if i.HasRelease("nano", 100) {
		t.Errorf("Should not have %s release %d", "nano", 100)
	}
}
