package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

func Execute() {
	app := cli.NewApp()
	app.Name = "json2graphql"
	app.Usage = inspectCmd.Usage
	app.Description = "inspect json and generate draft schema.graphql"
	app.HideVersion = true
	app.Flags = inspectCmd.Flags
	app.Version = "0.0.2"
	app.Before = func(context *cli.Context) error {
		if context.Bool("verbose") {
			log.SetFlags(0)
		} else {
			log.SetOutput(ioutil.Discard)
		}
		return nil
	}

	app.Action = inspectCmd.Action
	app.Commands = []cli.Command{
		inspectCmd,
		serverCmd,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
