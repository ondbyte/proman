package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

type Config struct {
	Language                string `json:"language"`
	InputFolder             string `json:"inputFolder"`
	OutputFolder            string `json:"outputFolder"`
	ShouldGenerateGrpcStubs bool   `json:"shouldGenerateGrpcStubs"`
}

func readeCfg() (c *Config, err error) {
	bs, err := os.ReadFile("./.proman")
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}
	c = &Config{}
	err = json.Unmarshal(bs, &c)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %v", err)
	}
	return
}

var cfgCmd = &cli.Command{
	Name:  "cfg",
	Usage: "proman config related commands",
	Commands: []*cli.Command{
		{
			Name:  "init",
			Usage: "initializes proman config file",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				if _, err := os.Stat("./.proman"); err == nil {
					return fmt.Errorf("config file already exists")
				}
				bs, err := json.Marshal(&Config{})
				if err != nil {
					return fmt.Errorf("failed to marshal config: %v", err)
				}
				err = os.WriteFile("./.proman", bs, 0777)
				if err != nil {
					return fmt.Errorf("failed to write config file: %v", err)
				}
				return nil
			},
		},
	},
}
