package main

import (
	"fmt"
	"github.com/alexpfx/go_common/str"
	"github.com/alexpfx/go_quickfix/internal/action"
	"log"
	"os"

	"github.com/alexpfx/go_common/exception"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "query",
				Usage: "busca na lista de actions e retorna à que atende ao pattern",
			},
			&cli.BoolFlag{
				Name:  "auto",
				Usage: "Se o comando query retornar apenas um resultado, executa a action",
			},
			&cli.IntFlag{
				Name:  "execute",
				Usage: "Executa o item de índice n. Usado em conjunto com --query",
				Value: -1,
			},
		},
		Commands: []*cli.Command{
			{},
		},
		Action: func(context *cli.Context) error {

			query := context.String("query")
			if query != "" {
				res := action.All().Query(query)

				jsonRes, err := str.FormatJson(res)
				exception.CheckThrow(err)

				lenRes := len(res)
				if lenRes == 0 {
					log.Printf("nenhum action disponível para a query: %s", query)
					return nil
				}

				if lenRes == 1 && context.Bool("auto") {
					execute(res, 0, query)
				}

				index := context.Int("execute")
				if index == -1 {
					fmt.Println(jsonRes)
					return nil
				}

				if index > lenRes {
					exception.CheckThrow(fmt.Errorf("elemento não disponível %d", index))
				}

				execute(res, index, query)

				return nil
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	exception.CheckThrow(err)
}

func execute(items []action.Item, index int, query string) {
	item := items[index]
	rStr := item.Replace(query)
	fmt.Println(rStr)
}
