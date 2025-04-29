package languages

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// FindProtocGenGo tries to find the protoc-gen-go binary across all common locations.
func FindProtocGenGo() (string, error) {
	binaryName := "protoc-gen-go"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	// 1. Try to find in PATH
	if path, err := exec.LookPath(binaryName); err == nil {
		return filepath.Abs(path)
	}

	// 2. Try GOBIN
	if gobin := os.Getenv("GOBIN"); gobin != "" {
		path := filepath.Join(gobin, binaryName)
		if fileExists(path) {
			return filepath.Abs(path)
		}
	}

	// 3. Try GOPATH/bin
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		path := filepath.Join(gopath, "bin", binaryName)
		if fileExists(path) {
			return filepath.Abs(path)
		}
	} else {
		// If GOPATH not set, use default
		homeDir, err := os.UserHomeDir()
		if err == nil {
			defaultGoPath := filepath.Join(homeDir, "go")
			path := filepath.Join(defaultGoPath, "bin", binaryName)
			if fileExists(path) {
				return filepath.Abs(path)
			}
		}
	}

	// 4. Try local ./bin directory (common in projects)
	localBin := filepath.Join(".", "bin", binaryName)
	if fileExists(localBin) {
		return filepath.Abs(localBin)
	}

	// 5. Could not find
	return "", errors.New("protoc-gen-go not found in PATH, GOBIN, GOPATH/bin, or ./bin")
}

// fileExists checks if a file exists and is not a directory
func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
