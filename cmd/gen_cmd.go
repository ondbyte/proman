package cmd

import (
	"context"
	"fmt"

	"github.com/ondbyte/proman/protocmanager"
	"github.com/urfave/cli/v3"
)

var genCmd = &cli.Command{
	Name:  "gen",
	Usage: "generates source files from proto files",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "lang",
			Aliases: []string{"l"},
			Usage:   "comma separated languages to generate",
		},
		&cli.StringFlag{
			Name:    "in",
			Aliases: []string{"i"},
			Usage:   "folder containing proto files",
		},
		&cli.StringFlag{
			Name:    "out",
			Aliases: []string{"o"},
			Usage:   "folder to output the generated source files",
		},
		&cli.BoolFlag{
			Name: "grpc",
		},
		&cli.StringFlag{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "additional commands to pass to protoc, should be in format \"ARG1=ARG_VALUE ARG2=ARG2_VALUE\"",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		cfg, err := readeCfg()
		if err != nil {
			return fmt.Errorf("error reading config file: %w", err)
		}
		if cfg != nil {
			if !c.IsSet("lang") && cfg.Language != "" {
				c.Set("lang", cfg.Language)
			}
			if !c.IsSet("in") && cfg.InputFolder != "" {
				c.Set("in", cfg.InputFolder)
			}
			if !c.IsSet("out") && cfg.OutputFolder != "" {
				c.Set("out", cfg.OutputFolder)
			}
			if !c.IsSet("grpc") {
				c.Set("grpc", fmt.Sprintf("%v", cfg.ShouldGenerateGrpcStubs))
			}
		}
		errS := ""
		if !c.IsSet("lang") {
			errS += "--lang is required\n"
		}
		if !c.IsSet("in") {
			errS += "--in is required\n"
		}
		if !c.IsSet("out") {
			errS += "--out is required"
		}
		if errS != "" {
			return fmt.Errorf(errS)
		}
		lang := c.String("lang")
		in := c.String("in")
		out := c.String("out")
		grpc := c.Bool("grpc")
		add := c.String("add")
		return protocmanager.Generate(lang, in, out, add, grpc)
	},
}
