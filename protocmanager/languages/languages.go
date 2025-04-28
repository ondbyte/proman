package languages

import (
	"fmt"
	"strings"
)

var Languages = map[string]Language{}

func LanguagesFromCommaSeparatedList(list string) ([]Language, error) {
	//split list
	splitted := strings.Split(list, ",")
	langs := []Language{}
	for _, lang := range splitted {
		if Languages[strings.Trim(lang, " ")] != nil {
			langs = append(langs, Languages[lang])
		} else {
			return nil, fmt.Errorf("Language %s not found", lang)
		}
	}
	return langs, nil
}

func RegisterLanguage(lang Language) bool {
	Languages[lang.Command()] = lang
	return true
}

type Plugin struct {
	Name            string
	InstallLocation string
}

type Language interface {
	IsInstalled() bool
	Name() string
	Command() string
	InstallPlugins() error
	CmdForGenSource(protocCmd, protoFolder, sourceOutputFolder string, grpc bool) ([]string, error)
}
