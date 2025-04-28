package protocmanager

import (
	"fmt"

	"github.com/ondbyte/proman/protocmanager/languages"
)

var plugins = languages.Languages

func InstallLangPlugins() error {
	fmt.Println("installing protoc plugins")
	for _, lang := range plugins {
		fmt.Println("---------------------***------------------------------------------------")
		if !lang.IsInstalled() {
			fmt.Printf("your system doesn't support language: '%v', skipping it.\n", lang.Name())
			continue
		}
		fmt.Printf("installing latest protoc plugins for language: '%v'\n", lang.Name())
		err := lang.InstallPlugins()
		if err != nil {
			fmt.Printf("failed to install protoc plugins for language: '%v', err: %v\n", lang.Name(), err)
			continue
		}
		fmt.Printf("done installing protoc plugins for language: '%v'\n", lang.Name())
	}
	return nil
}
