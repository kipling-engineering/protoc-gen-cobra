package main

import (
	"bytes"
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

var regenerate = flag.Bool("regenerate", false, "regenerate golden files")

// When the environment variable RUN_AS_PROTOC_GEN_COBRA is set, we skip running tests and instead
// act as protoc-gen-cobra. This allows the test binary to pass itself to protoc.
func init() {
	if os.Getenv("RUN_AS_PROTOC_GEN_COBRA") != "" {
		main()
		os.Exit(0)
	}
}

func TestGolden(t *testing.T) {
	workdir, err := os.MkdirTemp("", "proto-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(workdir)

	packages := map[string][]string{}
	err = filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".proto") {
			return nil
		}
		dir := filepath.Dir(path)
		packages[dir] = append(packages[dir], path)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	// Compile each package, using this binary as protoc-gen-cobra.
	for dir, sources := range packages {
		dir := filepath.Join(workdir, dir)
		if err := os.MkdirAll(dir, 0o666); err != nil {
			t.Fatal(err)
		}
		args := []string{"-Itestdata", "-Itestdata/oneof", "--cobra_out=" + dir}
		args = append(args, sources...)
		t.Log(args)
		protoc(t, args)
	}

	filepath.Walk(workdir, func(genPath string, info os.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(workdir, genPath)
		if err != nil {
			t.Errorf("filepath.Rel(%q, %q): %v", workdir, genPath, err)
			return nil
		}
		if filepath.SplitList(relPath)[0] == ".." {
			t.Errorf("generated file %q is not relative to %q", genPath, workdir)
		}

		goldenPath := relPath

		got, err := os.ReadFile(genPath)
		if err != nil {
			t.Error(err)
			return nil
		}
		if *regenerate {
			// If --regenerate set, just rewrite the golden files.
			err := os.WriteFile(goldenPath, got, 0o666)
			if err != nil {
				t.Error(err)
			}
			return nil
		}

		want, err := os.ReadFile(goldenPath)
		if err != nil {
			t.Error(err)
			return nil
		}

		want = fdescRE.ReplaceAll(want, nil)
		got = fdescRE.ReplaceAll(got, nil)
		if bytes.Equal(got, want) {
			return nil
		}

		cmd := exec.Command("diff", "-u", goldenPath, genPath)
		out, _ := cmd.CombinedOutput()
		t.Errorf("golden file differs: %v\n%v", relPath, string(out))
		return nil
	})
}

var fdescRE = regexp.MustCompile(`(?ms)^var fileDescriptor.*}`)

func protoc(t *testing.T, args []string) {
	cmd := exec.Command("protoc", "--plugin=protoc-gen-cobra="+os.Args[0])
	cmd.Args = append(cmd.Args, args...)
	// We set the RUN_AS_PROTOC_GEN_COBRA environment variable to indicate that the subprocess
	// should act as a proto compiler rather than a test.
	cmd.Env = append(os.Environ(), "RUN_AS_PROTOC_GEN_COBRA=1")
	out, err := cmd.CombinedOutput()
	if len(out) > 0 || err != nil {
		t.Log("RUNNING: ", strings.Join(cmd.Args, " "))
	}
	if len(out) > 0 {
		t.Log(string(out))
	}
	if err != nil {
		t.Fatalf("protoc: %v", err)
	}
}
