package page

import (
	"context"
	"errors"
	"history-engine/engine/model"
	"history-engine/engine/service/filetype"
	"history-engine/engine/service/host"
	serviceSetting "history-engine/engine/service/setting"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
	"strconv"
	"strings"
)

func Filter(hi *model.HtmlInfo) (bool, error) {
	storageSetting, err := serviceSetting.GetSetting(context.Background(), hi.UserId)
	if err != nil {
		return false, err
	}

	if hi.Size > 0 && (hi.Size < 2048 || hi.Size > storageSetting.MaxSize) {
		return false, errors.New("ignore " + hi.Url + " by size: " + strconv.Itoa(hi.Size))
	}

	if hi.Path != "" && !utils.FileExist(setting.Common.HtmlPath+hi.Path) {
		return false, errors.New("ignore " + hi.Url + " by file not exist： " + hi.Path)
	}

	if hi.Host != "" && !host.Include(hi.UserId, hi.Host) && host.Exclude(hi.UserId, hi.Host) {
		return false, errors.New("ignore " + hi.Url + " by host rule: " + hi.Host)
	}

	if hi.Url != "" && !host.Include(hi.UserId, hi.Url) && host.Exclude(hi.UserId, hi.Url) {
		return false, errors.New("ignore " + hi.Url + " by url rule: " + hi.Url)
	}

	if hi.Suffix != "" && !filetype.Include(hi.UserId, hi.Suffix) && filetype.Exclude(hi.UserId, hi.Suffix) {
		return false, errors.New("ignore " + hi.Url + " by suffix: " + hi.Suffix)
	}

	if hi.Sha1 != "" {
		_, created := NextVersion(context.Background(), hi.Sha1)
		if utils.CheckVersionInterval(storageSetting.MinVersionInterval, created) {
			return false, errors.New("ignore " + hi.Url + " by interval: " + hi.Sha1)
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
