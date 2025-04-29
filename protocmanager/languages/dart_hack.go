package languages

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func FindProtocGenDart() (string, error) {
	binaryName := "protoc-gen-dart"
	if runtime.GOOS == "windows" {
		binaryName += ".bat"
	}

	// 1. Try PATH
	if path, err := exec.LookPath(binaryName); err == nil {
		return filepath.Abs(path)
	}

	// 2. Dart global install locations
	pathsToCheck := []string{}

	homeDir, err := os.UserHomeDir()
	if err == nil {
		if runtime.GOOS == "windows" {
			// %APPDATA%\Pub\Cache\bin
			if appData := os.Getenv("APPDATA"); appData != "" {
				pathsToCheck = append(pathsToCheck, filepath.Join(appData, "Pub", "Cache", "bin", binaryName))
			}
			// %LOCALAPPDATA%\Pub\Cache\bin
			if localAppData := os.Getenv("LOCALAPPDATA"); localAppData != "" {
				pathsToCheck = append(pathsToCheck, filepath.Join(localAppData, "Pub", "Cache", "bin", binaryName))
			}
		} else {
			// ~/.pub-cache/bin
			pathsToCheck = append(pathsToCheck, filepath.Join(homeDir, ".pub-cache", "bin", binaryName))
		}
	}

	// 3. Optional local bin folder
	pathsToCheck = append(pathsToCheck, filepath.Join(".", "bin", binaryName))

	for _, path := range pathsToCheck {
		if fileExists(path) {
			return filepath.Abs(path)
		}
	}

	return "", errors.New("protoc-gen-dart not found in PATH, pub-cache, or ./bin")
}
