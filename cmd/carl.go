package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/xslook/carl/module/timestamp"
)

func appHandler(c *cli.Context) error {
	fmt.Println("Hello carl, enjoy yourself!")
	return nil
}

// Run command
func Run(ver, time, commit string) error {
	var app = &cli.App{
		Name:    "carl",
		Version: ver,
		Usage:   "A toolbox for development!",
		Action:  appHandler,
	}
	app.Commands = []*cli.Command{
		timestamp.Command(),
	}
	err := app.Run(os.Args)
	return err
}
