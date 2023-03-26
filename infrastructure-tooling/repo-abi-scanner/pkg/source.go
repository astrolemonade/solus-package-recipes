package pkg

import (
	"archive/tar"
	"bytes"
	"errors"
	"fmt"
	"github.com/DataDrake/abi-wizard/abi"
	log "github.com/DataDrake/waterlog"
	"github.com/getsolus/libeopkg/archive"
	git "github.com/go-git/go-git/v5"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Source represents a single Source Repo and its packages
type Source struct {
	Name     string
	Version  string
	Release  int
	Packages []string
	dir      string
	report   abi.Report
}

var (
	// ErrNoHistory missing history for this source repo
	ErrNoHistory = errors.New("missing History in source PSpec")
	// ErrNoPackages missing pacakges for this source repo
	ErrNoPackages = errors.New("missing Packages in source PSpec")
)

// NewSource creates a Source from a PSpec on disk
func NewSource(path string) (s *Source, err error) {
	spec, err := LoadPSpec(filepath.Join(path, "pspec_x86_64.xml"))
	if err != nil {
		if os.IsNotExist(err) {
			spec, err = LoadPSpec(filepath.Join(path, "pspec.xml"))
		}
		if err != nil {
			return
		}
	}
	if len(spec.History) == 0 {
		err = ErrNoHistory
		return
	}
	if len(spec.Packages) == 0 {
		err = ErrNoPackages
	}
	latest := spec.History[0]
	s = &Source{
		Name:    spec.Source.Name,
		Version: latest.Version,
		Release: latest.Release,
		dir:     filepath.Dir(path),
		report:  make(abi.Report),
	}
	for _, p := range spec.Packages {
		s.Packages = append(s.Packages, p.Name)
	}
	return
}

func getPathComponent(name string) string {
	nom := strings.ToLower(name)
	letter := nom[0:1]
	if strings.HasPrefix(nom, "lib") && len(nom) > 3 {
		return filepath.Join(nom[0:4], nom)
	}
	return filepath.Join(letter, nom)
}

// Scan reads each of the packages in this Source and adds them to the ABI report
func (s *Source) Scan(repo string) error {
	for _, p := range s.Packages {
		if err := s.ScanPackage(p, repo); err != nil {
			return err
		}
	}
	missing, err := s.report.Resolve()
	if err != nil {
		log.Fatalf("Failed to resolve symbols, reason: %s\n", err)
	}
	if len(missing) > 0 {
		log.Errorln("Failed to find libraries:")
		for _, lib := range missing {
			log.Errorln("\t" + lib)
		}
	}
	return err
}

// ScanPackages adds a single Package to the ABI report
func (s *Source) ScanPackage(name, repo string) error {
	// calculate the path
	pkgName := fmt.Sprintf("%s-%s-%d-1-x86_64.eopkg", name, s.Version, s.Release)
	path := filepath.Join(repo, getPathComponent(s.Name), pkgName)
	// Open archive
	log.Infof("Scanning '%s'...\n", path)
	p, err := archive.Open(path)
	if err != nil {
		log.Fatalf("Failed to open package '%s', reason: %s\n", path, err)
	}
	defer p.Close()
	// Unpack tarball
	temp, err := ioutil.TempDir(os.TempDir(), "repo-abi")
	if err != nil {
		log.Fatalf("Failed to create temp directory, reason: %s\n", err)
	}
	defer removeTemp(temp)
	log.Debugf("Created temp directory '%s'\n", temp)
	if err = p.ExtractTarball(temp); err != nil {
		log.Fatalf("Failed to unpack tarball, reason: %s\n", err)
	}
	log.Debugf("Unpacked install.tar.xz to temp dir '%s'\n", temp)
	name = filepath.Join(temp, "install.tar")
	tarFile, err := os.Open(name)
	if err != nil {
		log.Fatalf("Failed to open tar '%s', reason: %s\n", name, err)
	}
	defer tarFile.Close()
	// Read the tarball
	log.Debugf("Opened tar '%s'\n", name)
	files := tar.NewReader(tarFile)
	h, err := files.Next()
	for err == nil {
		if h.Size == 0 {
			h, err = files.Next()
			continue
		}
		if h.Typeflag == tar.TypeDir {
			h, err = files.Next()
			continue
		}
		// ignore statically linked archives or debug symbols
		if strings.HasSuffix(h.Name, ".o") || strings.HasSuffix(h.Name, ".la") || strings.HasSuffix(h.Name, ".a") || strings.HasSuffix(h.Name, ".debug") || strings.HasSuffix(h.Name, ".debuginfo") {
			h, err = files.Next()
			continue
		}
		var raw []byte
		if raw, err = ioutil.ReadAll(files); err != nil {
			log.Fatalf("Failed to read file '%s', reason: %s\n", h.Name, err)
		}
		if err = s.report.AddFile(bytes.NewReader(raw), filepath.Base(h.Name)); err != nil && err != io.EOF {
			log.Fatalf("Failed to add file '%s', reason: %s\n", h.Name, err)
		}
		log.Debugf("Processed tar entry '%s'\n", h.Name)
		h, err = files.Next()
	}
	if err != nil && err != io.EOF {
		log.Fatalf("Failed to read from tar '%s', reason: %s\n", name, err)
	}
	log.Goodln("Done")
	return nil
}

func removeTemp(temp string) {
	if err := os.RemoveAll(temp); err != nil {
		log.Fatalf("Failed to remove temp dir '%s', reason: %s\n", temp, err)
	}
}

// Save writes the report out to the source directory
func (s *Source) Save() error {
	return s.report.Save(s.dir)
}

// Commit adds all ABI reports to the Git index and then commits them with a generic message
func (s *Source) Commit() error {
	log.Debugf("Opening repo '%s'...\n", s.Name)
	repo, err := git.PlainOpen(s.dir)
	if err != nil {
		return err
	}
	work, err := repo.Worktree()
	if err != nil {
		return err
	}
	log.Infoln("Adding ABI reports to Git")
	if err = work.AddGlob("abi_*"); err != nil {
		return err
	}
	status, err := work.Status()
	if err != nil {
		return err
	}
	if status.IsClean() {
		log.Debugf("No changes found, skipping commit for '%s'\n", s.Name)
		return nil
	}
	_, err = work.Commit("Updated ABI reports using repo-abi-scanner", &git.CommitOptions{})
	return err
}
