package handler

import (
	"github.com/urfave/cli/v2"
	"history-engine/engine/service/page"
	"log"
	"net/url"
	"time"
)

var Analyse = &cli.Command{
	Name:    "analyse",
	Aliases: []string{"a"},
	Usage:   "Analyze the number of times the host appears in the database",
	Action:  runAnalyse,
}

func runAnalyse(ctx *cli.Context) error {
	host := make(map[string]int)
	start := 0
	limit := 100
	for {
		list, err := page.Page(ctx.Context, start, limit)
		if err != nil {
			panic(err)
		}

		if len(list) == 0 {
			break
		}

		start += limit
		time.Sleep(time.Millisecond * 100)

		for _, item := range list {
			parse, err := url.Parse(item.Url)
			if err != nil {
				panic(err)
			}
			if parse.Host == "" {
				continue
			}

			if _, ok := host[parse.Host]; ok {
				host[parse.Host] += 1
			} else {
				host[parse.Host] = 1
			}
		}
	}

	for item, count := range host {
		log.Println(count, "\t", item)
	}

	return nil
}
