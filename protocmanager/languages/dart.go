package languages

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var _ = RegisterLanguage(&Dart{})

type Dart struct {
}

// InstallPlugins implements Language.
func (d *Dart) InstallPlugins() error {
	plugins := []string{
		"protoc_plugin",
	}
	for _, v := range plugins {
		if err := exec.Command("dart", "pub", "global", "activate", v).Run(); err != nil {
			return fmt.Errorf("failed to install %v: %w", v, err)
		}
	}
	return nil
}

func readAllProtoFilesInDir(dir string) ([]string, error) {
	es, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %w", err)
	}
	protoFiles := []string{}
	for _, v := range es {
		if v.IsDir() {
			continue
		}
		if !strings.HasSuffix(v.Name(), ".proto") {
			continue
		}
		protoFiles = append(protoFiles, v.Name())
	}
	return protoFiles, nil
}

// CmdForGenSource implements Language.
func (d *Dart) CmdForGenSource(protocCmd, protoFolder string, sourceOutputFolder string, grpc bool) ([]string, error) {
	protocGenDartPath, err := Where("protoc-gen-dart")
	if err != nil {
		return nil, fmt.Errorf("failed to find protoc-gen-dart: %w", err)
	}
	dartOut := ""
	if grpc {
		dartOut += "grpc:"
	}
	dartOut += sourceOutputFolder
	protofiles, err := readAllProtoFilesInDir(protoFolder)
	if err != nil {
		return nil, fmt.Errorf("failed to read proto files: %w", err)
	}
	files := strings.Join(protofiles, " ")
	return []string{
		protocCmd,
		"--plugin=protoc-gen-dart=" + protocGenDartPath,
		"--proto_path=" + protoFolder,
		"--dart_out=" + dartOut,
		files,
	}, nil
}

// Command implements Language.
func (d *Dart) Command() string {
	return "dart"
}

// IsInstalled implements Language.
func (d *Dart) IsInstalled() bool {
	return exec.Command("dart").Run() == nil
}

// Name implements Language.
func (d *Dart) Name() string {
	return "dart"
}

// Plugins implements Language.
func (d *Dart) Plugins() []string {
	return []string{
		"dart pub global activate protoc_plugin",
	}
}

var _ Language = (*Dart)(nil)
