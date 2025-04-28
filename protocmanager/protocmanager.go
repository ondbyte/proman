package protocmanager

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
			&cli.Command{
				Name: "version",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println(version)
					return nil
				},
			},
			&cli.Command{
				Name:        "rm",
				Description: "removes any installed protoc, so it can be reinstalled in the next generation",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return RemoveProtoc()
				},
			},
			&cli.Command{
				Name: "gen",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "lang",
						Aliases:  []string{"l"},
						Usage:    "comma separated languages to generate",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "in",
						Aliases:  []string{"i"},
						Usage:    "folder containing proto files",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "out",
						Aliases:  []string{"o"},
						Usage:    "folder to output the generated source files",
						Required: true,
					},
					&cli.BoolWithInverseFlag{
						Name: "grpc",
					},
					&cli.StringFlag{
						Name:    "add",
						Aliases: []string{"a"},
						Usage:   "additional commands to pass to protoc, should be in format \"ARG1=ARG_VALUE ARG2=ARG2_VALUE\"",
					},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					lang := c.String("lang")
					in := c.String("in")
					out := c.String("out")
					grpc := c.Bool("grpc")
					add := c.String("add")
					return Generate(lang, in, out, add, grpc)
				},
			},
		},
	}
	err := app.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Generate(langs, in, out, add string, grpc bool) (err error) {
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
