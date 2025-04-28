package languages

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var goBinDir = filepath.Join(os.Getenv("GOPATH"), "bin")

func init() {
	RegisterLanguage(&Go{})
}

type Go struct {
}

// InstallPlugins implements Language.
func (d *Go) InstallPlugins() error {
	plugins := []string{
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
	}
	for _, v := range plugins {
		if err := exec.Command("go", "install", v).Run(); err != nil {
			return fmt.Errorf("failed to install %s: %w", v, err)
		}
	}
	return nil
}

// IsInstalled implements Language.
func (g *Go) IsInstalled() bool {
	return exec.Command("go", "version").Run() == nil
}

// CmdForGenSource implements Language.
func (g *Go) CmdForGenSource(protocCmd, protoFolder, sourceOutputFolder string, grpc bool) ([]string, error) {
	protocGenGoPath, err := Where("protoc-gen-go")
	if err != nil {
		return nil, fmt.Errorf("failed to find protoc-gen-go: %w", err)
	}
	pre := []string{
		protocCmd,
		"--plugin=protoc-gen-go=" + protocGenGoPath,
		"--proto_path=" + protoFolder,
		"--go_out=" + sourceOutputFolder,
		"--go_opt=paths=source_relative",
	}
	if grpc {
		pre = append(pre, "--go-grpc_out="+sourceOutputFolder, "--go-grpc_opt=paths=source_relative")
	}
	files, err := filepath.Glob(filepath.Join(protoFolder, "*.proto"))
	if err != nil {
		return nil, fmt.Errorf("failed to glob proto files: %w", err)
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no proto files found in %s", protoFolder)
	}

	return append(pre, files...), nil
}

// Command implements Language.
func (g *Go) Command() string {
	return "go"
}

// Name implements Language.
func (g *Go) Name() string {
	return "go"
}

// Plugins implements Language.
func (g *Go) Plugins() []string {
	return []string{
		"go install google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
	}
}

// PostPluginInstallMsg implements Language.
func (g *Go) PostPluginInstallMsg() string {
	p := filepath.Join(os.Getenv("$GOPATH"), "bin")
	return "make sure you have " + p + " is in your PATH"
}

var _ Language = (*Go)(nil)
