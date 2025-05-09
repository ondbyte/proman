package protocmanager

import (
	"fmt"
	"os"
	"os/exec"
)

func isGoogleProtosInstalled() bool {
	_, err := os.Stat("protobuf")
	return err != nil
}

func installGoogleProtos() error {
	installCmd := exec.Command("git", "clone", "https://github.com/protocolbuffers/protobuf.git")
	err := installCmd.Run()
	if err != nil {
		return fmt.Errorf("error downloading google protos: %w", err)
	}
	return nil
}
