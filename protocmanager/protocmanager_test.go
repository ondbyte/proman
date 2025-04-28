package protocmanager_test

import (
	"testing"

	"github.com/ondbyte/proman/protocmanager"
)

func TestProtocInstalled(t *testing.T) {
	if protocmanager.IsProtocInstalled() == nil {
		t.Log("Protoc is installed")
	} else {
		t.Error("Protoc is not installed")
	}
}

func TestInstallProtoc(t *testing.T) {
	err := protocmanager.InstallProtoc()
	if err != nil {
		panic(err)
	}
	err = protocmanager.InstallLangPlugins()
	if err != nil {
		panic(err)
	}
}
