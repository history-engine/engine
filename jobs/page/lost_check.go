package page

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	entPage "history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/service/host"
	"history-engine/engine/service/page"
	"history-engine/engine/service/readability"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var LostCheck = &cli.Command{
	Name:    "lost-check",
	Aliases: []string{"lc"},
	Usage:   "Check html file exists and other data is lost",
	Action:  runLostCheck,
}

func runLostCheck(ctx *cli.Context) error {
	x := db.GetEngine()
	users, err := x.User.Query().All(ctx.Context)
	if err != nil {
		return err
	}

	for _, user := range users {
		logger.Zap().Info("scan user html", zap.Int64("user_id", user.ID))

		root := fmt.Sprintf("%s/%d", setting.SingleFile.HtmlPath, user.ID)
		if !utils.PathExist(root) {
			continue
		}

		files := make(map[string]int64, 0)
		filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() {
				files[path] = info.Size()
			}
			return err
		})

		for file, size := range files {
			split := strings.Split(file, "/")
			name := split[len(split)-1]
			split = strings.Split(name, ".")
			uniqueId := split[0]
			version, err := strconv.Atoi(split[1])
			if err != nil {
				continue
			}

			exist, _ := x.Page.Query().
				Where(entPage.UserID(user.ID), entPage.UniqueID(uniqueId), entPage.Version(version)).
				Exist(ctx.Context)
			if exist {
				continue
			}

			head := make([]byte, 2048)
			f, err := os.Open(file)
			if err != nil {
				continue
			}

			_, err = f.Read(head)
			if err != nil {
				continue
			}

			url := readability.Parser().ExtractSingleFileUrl(head)
			if len(url) == 0 || (!host.Include(user.ID, url) && host.Exclude(user.ID, url)) {
				continue
			}

			htmlPath := strings.Replace(file, setting.SingleFile.HtmlPath, "", 1)
			row, err := x.Page.Create().
				SetUserID(user.ID).
				SetUniqueID(uniqueId).
				SetVersion(version).
				SetURL(url).
				SetPath(htmlPath).
				SetSize(int(size)).
				Save(ctx.Context)
			if err != nil {
				continue
			}

			if err := page.ParserPageWithId(row.ID); err != nil {
				panic(err)
			}
			logger.Zap().Info("find lost data", zap.String("path", file), zap.Int64("new id", row.ID))
		}
	}

	return nil
}
