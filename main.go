package main

import (
	"fmt"
	"github.com/alexpfx/go_common/str"
	"github.com/alexpfx/go_quickfix/internal/action"
	"github.com/atotto/clipboard"
	"log"
	"os"

	"github.com/alexpfx/go_common/exception"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "query",
				Aliases: []string{"q"},
			},
			&cli.BoolFlag{
				Name:    "execute",
				Usage:   "Aplica o quickfix ao primeiro elemento retornado",
				Aliases: []string{"x"},
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
			execute := context.Bool("execute")
			if query != "" {
				items := action.All().Query(query)

				json, err := str.FormatJson(items)
				exception.CheckThrow(err)
				if len(items) == 0 {
					log.Printf("nenhum action disponível para a query: %s", query)
					return nil
				}

				if execute {
					first := items[0]
					newValue := first.Replace(query)
					_ = clipboard.WriteAll(newValue)
				}
				fmt.Println(json)

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
