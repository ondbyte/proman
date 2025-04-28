package languages

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func Where(command string) (string, error) {
	where := "which"
	if runtime.GOOS == "windows" {
		where = "where"
	}
	cmd := exec.Command(where, command)
	op, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("%v not found: %v", command, err)
	}
	op, _ = bytes.CutSuffix(op, []byte("\n"))
	// windows specific
	op, _ = bytes.CutSuffix(op, []byte("\r"))
	return string(op), nil
}
