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
	workdir := t.TempDir()

	packages := map[string][]string{}
	// Declare err before it's used by filepath.Walk
	var err error
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
	for protoDir, sources := range packages {
		outputDirForProtos := filepath.Join(workdir, protoDir)
		if err := os.MkdirAll(outputDirForProtos, 0755); err != nil {
			t.Fatalf("Failed to create output dir %s: %v", outputDirForProtos, err)
		}

		protocArgs := []string{
			"-I/tmp/protoc/include", // For google.protobuf well-known types
			"-Itestdata",            // For local .proto includes
			"-Itestdata/oneof",      // For local .proto includes within oneof
			"--go_out=" + outputDirForProtos,
			"--cobra_out=" + outputDirForProtos,
		}
		protocArgs = append(protocArgs, sources...)

		t.Logf("Running protoc for package %s: %v", protoDir, protocArgs)
		runProtoc(t, protocArgs)

		// Log files generated in outputDirForProtos
		generatedFilesInDir := []string{}
		filepath.Walk(outputDirForProtos, func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() {
				generatedFilesInDir = append(generatedFilesInDir, filepath.Base(path))
			}
			return nil
		})
		t.Logf("Files generated in %s: %v", outputDirForProtos, generatedFilesInDir)

		// Process generated files for this package
		err := filepath.Walk(outputDirForProtos, func(generatedFilePath string, info os.FileInfo, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if info.IsDir() {
				return nil
			}

			generatedFileName := filepath.Base(generatedFilePath)
			// expectedGoldenPath is relative to project root, e.g., "testdata/pb/bank.pb.go"
			expectedGoldenPath := filepath.Join(protoDir, generatedFileName)

			// Run goimports -w on the generated file in workdir first
			if true { // Scope currentGobin and goimportsResolvedPath
				currentGobin := ""
				currentGopath := os.Getenv("GOPATH")
				if currentGopath != "" {
					currentGobin = filepath.Join(currentGopath, "bin")
				} else {
					currentHome := os.Getenv("HOME")
					if currentHome != "" {
						currentGobin = filepath.Join(currentHome, "go", "bin")
					}
				}
				if currentGobin != "" {
					goimportsResolvedPath := filepath.Join(currentGobin, "goimports")
					goimportsCmd := exec.Command(goimportsResolvedPath, "-w", generatedFilePath) // Operate on generatedFilePath
					if goimportsErr := goimportsCmd.Run(); goimportsErr != nil {
						t.Logf("Warning: goimports -w %s (using %s) failed: %v", generatedFilePath, goimportsResolvedPath, goimportsErr)
					}
				} else {
					t.Logf("Warning: GOBIN not found, cannot run goimports on %s", generatedFilePath)
				}
			}

			generatedContent, err := os.ReadFile(generatedFilePath) // Read after goimports
			if err != nil {
				t.Errorf("Failed to read generated file %s (after goimports): %v", generatedFilePath, err)
				return nil // Continue checking other files
			}

			if *regenerate {
				// Ensure parent directory of goldenPath exists
				goldenParentDir := filepath.Dir(expectedGoldenPath)
				if err := os.MkdirAll(goldenParentDir, 0755); err != nil {
					t.Errorf("Failed to create parent dir for golden file %s: %v", expectedGoldenPath, err)
					return nil
				}
				if err := os.WriteFile(expectedGoldenPath, generatedContent, 0o666); err != nil {
					t.Errorf("Failed to write golden file %s: %v", expectedGoldenPath, err)
					// Do not return yet, try goimports if write succeeded partially or for cleanup
				}
				// Run goimports -w on the newly written golden file
				// goimportsPath := filepath.Join(gobin, "goimports") // gobin is not in this scope
                                // For simplicity in this edit, re-determine gobin. In a real scenario, pass it or make it global to the test.
                                currentGobin := ""
                                currentGopath := os.Getenv("GOPATH")
                                if currentGopath != "" {
                                    currentGobin = filepath.Join(currentGopath, "bin")
                                } else {
                                    currentHome := os.Getenv("HOME")
                                    if currentHome != "" {
                                        currentGobin = filepath.Join(currentHome, "go", "bin")
                                    }
                                }
                                if currentGobin != "" {
                                   goimportsResolvedPath := filepath.Join(currentGobin, "goimports")
                                   goimportsCmd := exec.Command(goimportsResolvedPath, "-w", expectedGoldenPath)
                                   if goimportsErr := goimportsCmd.Run(); goimportsErr != nil {
                                       t.Logf("Warning: goimports -w %s (using %s) failed: %v", expectedGoldenPath, goimportsResolvedPath, goimportsErr)
                                   }
                                } else {
                                   t.Logf("Warning: GOBIN not found, cannot run goimports on %s", expectedGoldenPath)
                                }
				// } // THIS EXTRA BRACE IS REMOVED
				return nil
			}

			expectedContent, err := os.ReadFile(expectedGoldenPath)
			if err != nil {
				t.Errorf("Failed to read golden file %s: %v (run with -regenerate to create)", expectedGoldenPath, err)
				return nil
			}

			// Normalize content (e.g., remove dynamic parts like fileDescriptor)
			normalizedExpected := fdescRE.ReplaceAll(expectedContent, nil)
			normalizedGenerated := fdescRE.ReplaceAll(generatedContent, nil)

			if !bytes.Equal(normalizedGenerated, normalizedExpected) {
				cmd := exec.Command("diff", "-u", expectedGoldenPath, generatedFilePath)
				out, _ := cmd.CombinedOutput()
				t.Errorf("Golden file differs for %s:\n%s", expectedGoldenPath, string(out))
			}
			return nil
		})
		if err != nil {
			t.Fatalf("Error walking generated files for package %s: %v", protoDir, err)
		}
	}
}

var fdescRE = regexp.MustCompile(`(?ms)^var fileDescriptor.*}`) // Corrected regex

// runProtoc is the refactored helper function
func runProtoc(t *testing.T, args []string) {
	// Determine GOBIN (consistent with Makefiles)
	gobin := ""
	gopath := os.Getenv("GOPATH")
	if gopath != "" {
		gobin = filepath.Join(gopath, "bin")
	} else {
		home := os.Getenv("HOME")
		if home != "" {
			gobin = filepath.Join(home, "go", "bin")
		}
	}
	if gobin == "" {
		t.Fatal("GOBIN could not be determined. Set GOPATH or ensure $HOME/go/bin exists.")
	}

	protocGenGoPath := filepath.Join(gobin, "protoc-gen-go")
	if _, err := os.Stat(protocGenGoPath); os.IsNotExist(err) {
		t.Fatalf("protoc-gen-go not found at %s. Run 'make deps' in testdata or example/pb.", protocGenGoPath)
	}
	protocGenGoGRPCPath := filepath.Join(gobin, "protoc-gen-go-grpc")
	if _, err := os.Stat(protocGenGoGRPCPath); os.IsNotExist(err) {
		t.Fatalf("protoc-gen-go-grpc not found at %s. Run 'make deps' in testdata or example/pb.", protocGenGoGRPCPath)
	}
	protocGenCobraPath := os.Args[0] // The test binary itself

	// Use the specific protoc binary path
	protocBinPath := "/tmp/protoc/bin/protoc"
	if _, err := os.Stat(protocBinPath); os.IsNotExist(err) {
		t.Fatalf("protoc not found at %s. Ensure ci/install-protoc.sh has run.", protocBinPath)
	}

	// Separate --go_out, --cobra_out, --go-grpc_out from original args
	// and ensure they point to the same outputDirForProtos
	var outputDir string
	var otherArgs []string
	for _, arg := range args {
		if strings.HasPrefix(arg, "--go_out=") {
			outputDir = strings.TrimPrefix(arg, "--go_out=")
			continue
		}
		if strings.HasPrefix(arg, "--cobra_out=") {
			// Assume same as go_out, already captured
			continue
		}
		if strings.HasPrefix(arg, "--go-grpc_out=") {
			// Assume same as go_out, already captured
			continue
		}
		otherArgs = append(otherArgs, arg)
	}
	if outputDir == "" {
		t.Fatal("--go_out not found in protoc args")
	}

	cmdArgs := []string{
		"--plugin=protoc-gen-go=" + protocGenGoPath,
		"--plugin=protoc-gen-go-grpc=" + protocGenGoGRPCPath,
		"--plugin=protoc-gen-cobra=" + protocGenCobraPath,
		"--go_out=" + outputDir,
		"--go-grpc_out=" + outputDir,
		"--cobra_out=" + outputDir,
	}
	cmdArgs = append(cmdArgs, otherArgs...) // Add -I paths, proto files

	cmd := exec.Command(protocBinPath, cmdArgs...)
	cmd.Env = append(os.Environ(), "RUN_AS_PROTOC_GEN_COBRA=1")

	out, err := cmd.CombinedOutput()
	if len(out) > 0 || err != nil {
		// Using t.Logf for more structured logging
		t.Logf("RUNNING: %s %s", protocBinPath, strings.Join(cmdArgs, " "))
	}
	if len(out) > 0 {
		t.Logf("PROTOC OUTPUT:\n%s", string(out))
	}
	if err != nil {
		t.Fatalf("protoc execution failed: %v", err)
	}
}
