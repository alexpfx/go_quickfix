package main

import (
	"os"

	"github.com/alexpfx/go_common/exception"
	"github.com/urfave/cli/v2"
)


func main() {


	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "query",
				Aliases: []string{"q"},
			},
			&cli.BoolFlag{
				Name:    "daemon",
				Usage:   "Inicia o modo daemon se este não estiver sendo executado",
				Aliases: []string{"d"},
			},
			&cli.BoolFlag{
				Name: "kill",
				Aliases: []string{
					"k",
				},
				Usage: "Pára o modo daemon",
			},
		},
		Commands: []*cli.Command{
			{},
		},
		Action: func(context *cli.Context) error {

			query := context.String("query")
			if (query != ""){
				
				return nil
			}

			runDaemon := context.Bool("daemon")
			if runDaemon {

				return nil

			}

			killDaemon := context.Bool("kill")
			if killDaemon {

				return nil
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	exception.CheckThrow(err)
}
