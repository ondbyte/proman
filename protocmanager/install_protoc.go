package protocmanager

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/go-github/v71/github"
)

func RemoveProtoc() error {
	return os.Remove(protocCmdPath)
}

func IsProtocInstalled() (err error) {
	err = exec.Command(protocCmdPath, "--version").Run()
	if err != nil {
		return fmt.Errorf("error running protoc: %w", err)
	}
	return nil
}

func InstallProtoc() (err error) {
	fmt.Println("installing latest protoc")
	gc := github.NewClient(http.DefaultClient)
	releases, _, err := gc.Repositories.ListReleases(context.Background(), "protocolbuffers", "protobuf", nil)
	if err != nil {
		return err
	}
	if len(releases) == 0 {
		return fmt.Errorf("no releases found")
	}
	latestRelease := releases[0]
	latestVersion := *latestRelease.TagName
	fmt.Println(latestVersion)
	latestVersion, _ = strings.CutPrefix(latestVersion, "v")
	filename, err := protocFilename(latestVersion, runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return err
	}
	var ast *github.ReleaseAsset

	for _, asset := range latestRelease.Assets {
		name := *asset.Name
		if name == filename {
			ast = asset
		}
	}
	if ast == nil {
		return fmt.Errorf("no release asset found")
	}
	assetURL := *ast.BrowserDownloadURL
	//download zip
	fmt.Println("downloading")
	cd, err := os.UserCacheDir()
	if err != nil {
		return fmt.Errorf("error getting user cache dir: %w", err)
	}
	filename = filepath.Join(cd, filename)
	err = downloadFile(assetURL, filename)
	if err != nil {
		return err
	}
	defer os.Remove(filename)
	folder := strings.TrimSuffix(filename, ".zip")
	defer os.RemoveAll(folder)
	fmt.Println("unzipping")
	err = unzip(filename, folder)
	if err != nil {
		return err
	}
	//move protoc to bin
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("error getting user config dir: %w", err)
	}
	promanDir := filepath.Join(cfgDir, "protocmanager")
	extractedProtoc := filepath.Join(folder, "bin", "protoc"+ext)
	err = os.Mkdir(promanDir, 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("error creating protocmanager dir: %w", err)
	}
	err = os.Rename(extractedProtoc, filepath.Join(promanDir, "protoc"+ext))
	if err != nil {
		return fmt.Errorf("error moving protoc: %w", err)
	}
	return IsProtocInstalled()
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("error opening zip file: %w", err)
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			// Create directories
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				return fmt.Errorf("error creating directory: %w", err)
			}
			continue
		}

		// Create parent directories if needed
		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return fmt.Errorf("error creating parent directories: %w", err)
		}

		// Create destination file
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return fmt.Errorf("error creating file: %w", err)
		}

		rc, err := f.Open()
		if err != nil {
			return fmt.Errorf("error opening zipped file: %w", err)
		}

		_, err = io.Copy(outFile, rc)

		// Close file handles
		outFile.Close()
		rc.Close()

		if err != nil {
			return fmt.Errorf("error writing file: %w", err)
		}
	}
	return nil
}

func downloadFile(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error downloading: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error saving file: %w", err)
	}

	return nil
}
func protocFilename(version string, goos string, goarch string) (string, error) {
	switch goos {
	case "linux":
		switch goarch {
		case "amd64":
			return fmt.Sprintf("protoc-%s-linux-x86_64.zip", version), nil
		case "386":
			return fmt.Sprintf("protoc-%s-linux-x86_32.zip", version), nil
		case "arm64":
			return fmt.Sprintf("protoc-%s-linux-aarch_64.zip", version), nil
		case "ppc64le":
			return fmt.Sprintf("protoc-%s-linux-ppcle_64.zip", version), nil
		case "s390x":
			return fmt.Sprintf("protoc-%s-linux-s390_64.zip", version), nil
		}
	case "darwin":
		switch goarch {
		case "amd64":
			return fmt.Sprintf("protoc-%s-osx-x86_64.zip", version), nil
		case "arm64":
			return fmt.Sprintf("protoc-%s-osx-aarch_64.zip", version), nil
		default:
			return fmt.Sprintf("protoc-%s-osx-universal_binary.zip", version), nil
		}
	case "windows":
		switch goarch {
		case "amd64":
			return fmt.Sprintf("protoc-%s-win64.zip", version), nil
		case "386":
			return fmt.Sprintf("protoc-%s-win32.zip", version), nil
		}
	}
	return "", fmt.Errorf("unsupported GOOS/GOARCH combo: %s/%s", goos, goarch)
}
