package cmd

import (
	"fmt"
	"os"

	"github.com/luojilab/json2graphqlschema/inspect"
	"github.com/urfave/cli"
)

var inspectCmd = cli.Command{
	Name:  "inspect",
	Usage: "generate a graphql schema based on json",
	Flags: []cli.Flag{
		cli.BoolFlag{Name: "verbose, v", Usage: "Show logs"},
		cli.StringFlag{Name: "file, f", Usage: "The json filename, Cannot be used with -u "},
		cli.StringFlag{Name: "output, o", Usage: "The target filename to store generated schema, default is draft.graphql"},
		cli.StringFlag{Name: "url, u", Usage: "The json request url, Cannot be used with -u"},
		cli.StringFlag{Name: "token, t", Usage: "the token of json request url"},
	},
	Action: func(ctx *cli.Context) {
		var err error

		var output string
		if output = ctx.String("output"); output == "" {
			output = "draft.graphql"
		}
		if ctx.String("file") == "" && ctx.String("url") == "" {
			fmt.Println("param not found: -f or -i is needed.")
			os.Exit(2)
		} else if ctx.String("file") != "" && ctx.String("url") != "" {
			fmt.Println("param error: -f cannot be used with -u")
			os.Exit(2)
		} else if filename := ctx.String("file"); filename != "" {
			fmt.Println("schema create with json file")
			if err = inspect.InspectWithFile(filename, output); err != nil {
				os.Exit(2)
			}
		} else {
			url := ctx.String("url")
			token := ctx.String("token")
			if token == "" {
				fmt.Println("token is empty. you can input -f to request with token")
			}
			if err = inspect.InspectWithUrl(url, output, token); err != nil {
				os.Exit(2)
			}
		}
	},
}
