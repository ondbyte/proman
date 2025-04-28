package cmd

import (
	"context"
	"fmt"

	"github.com/ondbyte/proman/protocmanager"
	"github.com/urfave/cli/v3"
)

var versionCmd = &cli.Command{
	Name: "version",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		fmt.Println(version)
		return nil
	},
}
var rmCmd = &cli.Command{
	Name:  "rm",
	Usage: "removes any installed protoc, so it can be reinstalled in the next generation",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		return protocmanager.RemoveProtoc()
	},
}
