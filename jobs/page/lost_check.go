package page

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	entPage "history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
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

		root := fmt.Sprintf("%s/%d", setting.Common.HtmlPath, user.ID)
		if !utils.PathExist(root) {
			logger.Zap().Info("user html root not exist：" + root)
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
				logger.Zap().Warn("version conv err", zap.String("version", split[1]))
				continue
			}

			exist, _ := x.Page.Query().
				Where(entPage.UserID(user.ID), entPage.UniqueID(uniqueId), entPage.Version(version)).
				Exist(ctx.Context)
			if exist {
				continue
			}

			logger.Zap().Info("find lost data", zap.String("path", file))

			head := make([]byte, 2048)
			f, err := os.Open(file)
			if err != nil {
				logger.Zap().Warn("open html file err", zap.Error(err), zap.String("file", file))
				continue
			}

			_, err = f.Read(head)
			if err != nil {
				logger.Zap().Warn("read html file err", zap.Error(err), zap.String("file", file))
				continue
			}

			url := readability.Parser().ExtractSingleFileUrl(head)
			if len(url) == 0 {
				continue
			}

			htmlPath := strings.Replace(file, setting.Common.HtmlPath, "", 1)
			_, err = x.Page.Create().
				SetUserID(user.ID).
				SetUniqueID(uniqueId).
				SetVersion(version).
				SetURL(url).
				SetPath(htmlPath).
				SetSize(int(size)).
				Save(ctx.Context)
			if err != nil {
				logger.Zap().Warn("create page err", zap.Error(err), zap.String("file", file))
				continue
			}
		}
	}

	return nil
}
