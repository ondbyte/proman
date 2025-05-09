package languages

import (
	"fmt"
	"os/exec"
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

// CmdForGenSource implements Language.
func (d *Dart) CmdForGenSource(
	protocCmd,
	importsPath,
	inputFolder string,
	sourceOutputFolder string,
	grpc bool,
) ([]string, error) {
	protocGenDartPath, err := FindProtocGenDart()
	if err != nil {
		return nil, fmt.Errorf("failed to find protoc-gen-dart: %w", err)
	}
	dartOut := ""
	if grpc {
		dartOut += "grpc:"
	}
	dartOut += sourceOutputFolder
	files, err := listProtoFileNamesInFolder(inputFolder)
	if err != nil {
		return nil, fmt.Errorf("failed to list proto files: %w", err)
	}
	pre := []string{
		protocCmd,
		"--plugin=protoc-gen-dart=" + protocGenDartPath,
		"--proto_path=" + importsPath,
		"--proto_path=" + inputFolder,
		"--dart_out=" + dartOut,
	}
	return append(pre, files...), nil
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
