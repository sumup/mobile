// Copyright 2016 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var tests = []struct {
	name string
	lang string
	pkg  string
	goos string
}{
	{"ObjC-Testpkg", "objc", "github.com/sumup/mobile/bind/testdata/testpkg", ""},
	{"Java-Testpkg", "java", "github.com/sumup/mobile/bind/testdata/testpkg", ""},
	{"Go-Testpkg", "go", "github.com/sumup/mobile/bind/testdata/testpkg", ""},
	{"Java-Javapkg", "java", "github.com/sumup/mobile/bind/testdata/testpkg/javapkg", "android"},
	{"Go-Javapkg", "go", "github.com/sumup/mobile/bind/testdata/testpkg/javapkg", "android"},
	{"Go-Javapkg", "go,java,objc", "github.com/sumup/mobile/bind/testdata/cgopkg", "android"},
}

func installGobind() error {
	if out, err := exec.Command("go", "install", "github.com/sumup/mobile/cmd/gobind").CombinedOutput(); err != nil {
		return fmt.Errorf("gobind install failed: %v: %s", err, out)
	}
	return nil
}

func runGobind(lang, pkg, goos string) error {
	cmd := exec.Command("gobind", "-lang", lang, pkg)
	if goos != "" {
		cmd.Env = append(os.Environ(), "GOOS="+goos)
		cmd.Env = append(os.Environ(), "CGO_ENABLED=1")
	}
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("gobind -lang %s %s failed: %v: %s", lang, pkg, err, out)
	}
	return nil
}

func TestGobind(t *testing.T) {
	if err := installGobind(); err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := runGobind(test.lang, test.pkg, test.goos); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestDocs(t *testing.T) {
	if err := installGobind(); err != nil {
		t.Fatal(err)
	}
	// Create a fake package for doc.go
	tmpdir, err := ioutil.TempDir("", "gobind-test-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)
	docPkg := filepath.Join(tmpdir, "src", "doctest")
	if err := os.MkdirAll(docPkg, 0700); err != nil {
		t.Fatal(err)
	}
	const docsrc = `
package doctest

// This is a comment.
type Struct struct{
}`
	if err := ioutil.WriteFile(filepath.Join(docPkg, "doc.go"), []byte(docsrc), 0700); err != nil {
		t.Fatal(err)
	}

	const comment = "This is a comment."
	for _, lang := range []string{"java", "objc"} {
		cmd := exec.Command("gobind", "-lang", lang, "doctest")
		cmd.Env = append(os.Environ(), "GOROOT="+tmpdir)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Errorf("gobind -lang %s failed: %v: %s", lang, err, out)
			continue
		}
		if bytes.Index(out, []byte(comment)) == -1 {
			t.Errorf("gobind output for language %s did not contain the comment %q", lang, comment)
		}
	}
}

func BenchmarkGobind(b *testing.B) {
	if err := installGobind(); err != nil {
		b.Fatal(err)
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if err := runGobind(test.lang, test.pkg, test.goos); err != nil {
					b.Error(err)
				}
			}
		})
	}
}
