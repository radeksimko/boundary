package config

import (
	"strings"

	"github.com/hashicorp/boundary/internal/cmd/base"
	"github.com/mitchellh/cli"
)

var _ cli.Command = (*Command)(nil)

type Command struct {
	*base.Command
}

func (c *Command) Synopsis() string {
	return "Manage sensitive values in Boundary's configuration files"
}

func (c *Command) Help() string {
	helpText := `
Usage: boundary config <subcommand> [options] [args]

  This command groups subcommands for operators interacting with Boundary's
  config files. Here are a few examples of config commands:

    Encrypt sensitive values in a config file:

    $ boundary config encrypt config.hcl

    Decrypt sensitive values in a config file:

    $ boundary config decrypt config.hcl

  Please see the individual subcommand help for detailed usage information.`

	return strings.TrimSpace(helpText)
}

func (c *Command) Run(args []string) int {
	return cli.RunResultHelp
}
