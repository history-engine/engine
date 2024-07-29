package page

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	entPage "history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/service/page"
	"history-engine/engine/service/zincsearch"
	"time"
)

var timeZero = time.Time{}

var IndexCheck = &cli.Command{
	Name:    "index-check",
	Aliases: []string{"ic"},
	Usage:   "Check ZincSearch Document consistency",
	Action:  runIndexCheck,
}

func runIndexCheck(ctx *cli.Context) error {
	start := 0
	limit := 100
	time.Now().IsZero()
	x := db.GetEngine().Page
	for {
		list, err := x.Query().
			Where(entPage.And(entPage.ContentNEQ(""), entPage.IndexedAtEQ(timeZero))).
			Order(ent.Desc(entPage.FieldID)).
			Offset(start).
			Limit(limit).
			All(ctx.Context)
		if err != nil {
			panic(err)
		}

		if len(list) == 0 {
			break
		}

		start += limit
		time.Sleep(time.Millisecond * 100)

		for _, item := range list {
			doc, err := zincsearch.GetDocument(item.UserID, item.UniqueID, item.Version)
			if err != nil && err.Error() != "id not found" {
				continue
			}

			if doc != nil {
				x.Update().SetIndexedAt(time.Now()).Where(entPage.ID(item.ID)).Save(ctx.Context)
				continue
			}
			if err := page.PutIndexWithId(item.ID); err != nil {
				logger.Zap().Warn("put index err", zap.Error(err), zap.Int64("id", item.ID))
			}
		}
	}
	return nil
}
