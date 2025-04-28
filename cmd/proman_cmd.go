package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/ondbyte/proman/protocmanager/languages"
	"github.com/urfave/cli/v3"
)

var version = "v0.0.1"

func description() (s string) {
	s = `protocmanager is cli tool to simplify the process of 
	* setting up your machine with protobuff compiler(protoc)
	* installing compiler's language specific plugins
	* generating source files from proto files

protocmanager supports the following languages:
`
	for _, lang := range languages.Languages {
		s += fmt.Sprintf("%s, ", lang.Command())
	}
	s, _ = strings.CutSuffix(s, ",")
	return
}

func Main() {
	app := cli.Command{
		Name: "proman",
		//Usage:       "proman gen --in=<folder> --out=<folder> --grpc",
		Usage: description(),
		Commands: []*cli.Command{
			cfgCmd,
			genCmd,
			versionCmd,
			rmCmd,
		},
	}
	err := app.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
