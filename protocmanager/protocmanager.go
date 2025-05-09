package protocmanager

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ondbyte/proman/protocmanager/languages"
)

var (
	ext                = ""
	protocManagerDir   = ""
	protocCmdPath      = ""
	googleProtosSrcDir = ""
)

func init() {
	if runtime.GOOS == "windows" {
		ext = ".exe"
	}
	userCfgDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	protocManagerDir = filepath.Join(userCfgDir, "protocmanager")
	protocCmdPath = filepath.Join(protocManagerDir, "protoc"+ext)
	googleProtosSrcDir = filepath.Join(protocManagerDir, "google-protos-src")
}

func Generate(langs, in, out, add string, grpc bool) (err error) {
	if err := os.Chdir(protocManagerDir); err != nil {
		return fmt.Errorf("error changing directory to protoc manager directory: %v", err)
	}
	if IsProtocInstalled() != nil {
		defer func() {
			if err != nil {
				err := RemoveProtoc()
				if err != nil {
					fmt.Println("error removing protoc")
				}
			}
		}()
		fmt.Println("protoc not found locally")
		err := InstallProtoc()
		if err != nil {
			return fmt.Errorf("error installing protoc: %v", err)
		}
		err = InstallLangPlugins()
		if err != nil {
			return fmt.Errorf("error installing language plugins: %v", err)
		}
	}

	if !isGoogleProtosInstalled() {
		if err := installGoogleProtos(); err != nil {
			return fmt.Errorf("error installing google protos: %v", err)
		}
	}

	langsToGen, err := languages.LanguagesFromCommaSeparatedList(langs)
	if err != nil {
		return fmt.Errorf("error getting languages to generate: %v", err)
	}

	in, err = filepath.Abs(in)
	if err != nil {
		return fmt.Errorf("failed to get absolute path of in folder: %w", err)
	}
	out, err = filepath.Abs(out)
	if err != nil {
		return fmt.Errorf("failed to get absolute path of out folder: %w", err)
	}
	for _, language := range langsToGen {
		cmdToExec, err := language.CmdForGenSource(protocCmdPath, in, out, grpc)
		if err != nil {
			return fmt.Errorf("error getting command to execute: %v", err)
		}
		cmd := exec.Command(cmdToExec[0], cmdToExec[1:]...)
		op, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error running command %v\n%v:\n %v", strings.Join(cmdToExec, " "), err, string(op))
		}
	}
	fmt.Println("generated succesfully")
	return nil
}
