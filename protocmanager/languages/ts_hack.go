package languages

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func FindProtocGenTS() (string, error) {
	binaryName := "protoc-gen-ts"
	if runtime.GOOS == "windows" {
		binaryName += ".cmd"
	}

	// 1. Try PATH
	if path, err := exec.LookPath(binaryName); err == nil {
		return filepath.Abs(path)
	}

	pathsToCheck := []string{}

	// 2. npm prefix -g
	if npmPrefix, err := exec.Command("npm", "prefix", "-g").Output(); err == nil {
		prefix := strings.TrimSpace(string(npmPrefix))
		binPath := filepath.Join(prefix, "bin", binaryName)
		pathsToCheck = append(pathsToCheck, binPath)
	}

	// 3. OS-specific fallback
	homeDir, err := os.UserHomeDir()
	if err == nil {
		if runtime.GOOS == "windows" {
			if appData := os.Getenv("APPDATA"); appData != "" {
				pathsToCheck = append(pathsToCheck, filepath.Join(appData, "npm", binaryName))
			}
			if localAppData := os.Getenv("LOCALAPPDATA"); localAppData != "" {
				pathsToCheck = append(pathsToCheck, filepath.Join(localAppData, "npm", binaryName))
			}
		} else {
			pathsToCheck = append(pathsToCheck, filepath.Join(homeDir, ".npm-global", "bin", binaryName))
		}
	}

	// 4. Local project node_modules/.bin
	local := filepath.Join(".", "node_modules", ".bin", binaryName)
	pathsToCheck = append(pathsToCheck, local)

	// Final loop
	for _, path := range pathsToCheck {
		if fileExists(path) {
			return filepath.Abs(path)
		}
	}

	return "", errors.New("protoc-gen-ts not found in PATH, npm global locations, or local node_modules")
}
