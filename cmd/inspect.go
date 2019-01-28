package main

import (
	"os"
	"fmt"
	"log"
	"io/ioutil"

	"github.com/urfave/cli"
	"github.luojilab.com/json2graphqlschema/inspect"
)

var inspectCmd = cli.Command{
	Name:  "inspect",
	Usage: "generate a graphql schema based on json",
	Flags: []cli.Flag{
		cli.BoolFlag{Name: "verbose, v", Usage: "show logs"},
		cli.StringFlag{Name: "input, i", Usage: "the json filename"},
		cli.StringFlag{Name: "output, o", Usage: "the target filename to store generated schema"},
	},
	Action: func(ctx *cli.Context) {
		var err error
		var input, output string
		if input = ctx.String("input"); input == "" {
			fmt.Println("param not found: input or -i is needed.")
			os.Exit(2)
		}

		if output = ctx.String("output"); output == "" {
			output = "draft.graphql"
		}

		err = inspect.Inspect(input, output)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(3)
		}
	},
}

func Execute() {
	app := cli.NewApp()
	app.Name = "inspect"
	app.Usage = inspectCmd.Usage
	app.Description = "inspect json and generate draft schema.graphql"
	app.HideVersion = true
	app.Flags = inspectCmd.Flags
	app.Version = "0.0.1"
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
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func main() {
	Execute()
}

