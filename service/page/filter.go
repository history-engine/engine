package page

import (
	"context"
	"errors"
	"history-engine/engine/model"
	"history-engine/engine/service/filetype"
	"history-engine/engine/service/host"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
	"strconv"
	"strings"
)

func Filter(hi *model.HtmlInfo) (bool, error) {
	if hi.Size > 0 && (hi.Size < 2048 || hi.Size > setting.SingleFile.MaxSize) {
		return false, errors.New("ignore by size: " + strconv.Itoa(hi.Size))
	}

	if hi.Path != "" && !utils.FileExist(setting.SingleFile.HtmlPath+hi.Path) {
		return false, errors.New("ignore by file not exist： " + hi.Path)
	}

	if hi.Host != "" && !host.Include(hi.UserId, hi.Host) && host.Exclude(hi.UserId, hi.Host) {
		return false, errors.New("ignore by rule: " + hi.Host)
	}

	if hi.Url != "" && !host.Include(hi.UserId, hi.Url) && host.Exclude(hi.UserId, hi.Url) {
		return false, errors.New("ignore by rule: " + hi.Url)
	}

	if hi.Suffix != "" && !filetype.Include(hi.UserId, hi.Suffix) && filetype.Exclude(hi.UserId, hi.Suffix) {
		return false, errors.New("ignore by suffix: " + hi.Suffix)
	}

	if hi.Sha1 != "" {
		_, created := NextVersion(context.Background(), hi.Sha1)
		if utils.CheckVersionInterval(setting.SingleFile.MinVersionInterval, created) {
			return false, errors.New("ignore by interval: " + hi.Sha1)
		}
	}

	return true, nil
}

func ParseHtmlInfo(filename string) *model.HtmlInfo {
	if len(filename) == 0 {
		return nil
	}

	sff := &model.HtmlInfo{}
	kvs := strings.Split(filename, "_")
	for _, item := range kvs {
		kv := strings.SplitN(item, "-", 2)
		if len(kv) != 2 {
			continue
		}

		switch kv[0] {
		case "host":
			sff.Host = kv[1]
		case "suffix":
			sff.Suffix = utils.FileSuffix(kv[1])
		case "sha1":
			sff.Sha1 = kv[1]
		}
	}

	if sff.Sha1 == "" && sff.Url != "" {
		sff.Sha1 = utils.Sha1Str(sff.Sha1)
	}

	if sff.Sha1 == "" {
		return nil
	}

	return sff
}
