package handler

import (
	"github.com/urfave/cli/v2"
	"history-engine/engine/service/page"
	"log"
	"net/url"
	"sort"
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
			parse, err := url.Parse(item.URL)
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

	type count struct {
		host  string
		count int
	}
	countList := make([]count, 0)

	for item, number := range host {
		countList = append(countList, count{item, number})
	}

	sort.Slice(countList, func(i, j int) bool {
		return countList[i].count > countList[j].count
	})

	for _, v := range countList {
		log.Println(v.count, "\t", v.host)
	}

	return nil
}
