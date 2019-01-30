package cmd

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/luojilab/json2graphqlschema/server"
)

var serverCmd = cli.Command{
	Name:  "server",
	Usage: "run a server that convert json to graphql schema",
	Flags: []cli.Flag{
		cli.BoolFlag{Name: "verbose, v", Usage: "show logs"},
		cli.StringFlag{Name: "port, p", Usage: "assign listening port(default is 8080)"},
	},
	Action: func(ctx *cli.Context) {
		var port string
		if port = ctx.String("port"); port == "" {
			port = ":8080"
		}
		fmt.Println("before")
		server.Run(port)
		fmt.Println("after")
	},
}
