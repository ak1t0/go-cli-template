package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"go-cli-template/commands"
)

var (
	Version  = "0"
	CommitId = "0"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli"
	app.Usage = "boilerplate"
	app.Version = fmt.Sprintf("%s - %s", Version, CommitId)

	app.Commands = commands.Commands

	app.Run(os.Args)
}
