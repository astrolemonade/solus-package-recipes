package order

import (
	"testing"
)

func newFile() *File {
	return &File{
		Tiers: []*Tier{
			&Tier{
				Name: "0",
				Packages: []string{
					"a",
					"b",
					"c",
				},
			},
			&Tier{
				Name: "1",
				Packages: []string{
					"d",
					"e",
					"f",
				},
			},
		},
	}
}

func TestFileNext(t *testing.T) {
	f := newFile()
	tier, ok := f.Next()
	if !ok {
		t.Error("should have found a tier")
	}
	if tier == nil {
		t.Error("tier should not be nil")
		t.FailNow()
	}
	if tier.Name != "0" {
		t.Errorf("tier should be '%s', found '%s'", "0", tier.Name)
	}
	tier, ok = f.Next()
	if !ok {
		t.Error("should have found a tier")
	}
	if tier == nil {
		t.Error("tier should not be nil")
		t.FailNow()
	}
	if tier.Name != "1" {
		t.Errorf("tier should be '%s', found '%s'", "1", tier.Name)
	}
	tier, ok = f.Next()
	if ok {
		t.Error("should not have found a tier")
	}
	if tier != nil {
		t.Error("tier should be nil")
		t.FailNow()
	}
}

func TestTierNext(t *testing.T) {
	f := newFile()
	tier, ok := f.Next()
	if !ok {
		t.Error("should have found a tier")
	}
	if tier == nil {
		t.Error("tier should not be nil")
		t.FailNow()
	}
	dir, ok := tier.Next()
	if !ok {
		t.Error("should have found a dir")
	}
	if dir != "a" {
		t.Errorf("expected dir '%s', found '%s'", "a", dir)
	}
	dir, ok = tier.Next()
	if !ok {
		t.Error("should have found a dir")
	}
	if dir != "b" {
		t.Errorf("expected dir '%s', found '%s'", "b", dir)
	}
	dir, ok = tier.Next()
	if !ok {
		t.Error("should have found a dir")
	}
	if dir != "c" {
		t.Errorf("expected dir '%s', found '%s'", "c", dir)
	}
	dir, ok = tier.Next()
	if ok {
		t.Error("should not have found a dir")
	}
	if len(dir) > 0 {
		t.Errorf("expected no dir, found '%s'", dir)
	}
}
