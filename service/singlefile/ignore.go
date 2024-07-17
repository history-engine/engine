package singlefile

import (
	"github.com/tidwall/match"
	"go.uber.org/zap"
	"history-engine/engine/library/logger"
	"history-engine/engine/setting"
	"net/url"
	"regexp"
	"time"
)

func CheckIgnore(address string) bool {
	if address == "" {
		return true
	}

	parse, err := url.Parse(address)
	if err != nil {
		logger.Zap().Warn("parse ignore host err", zap.Error(err))
		return false
	}

	var ignore = false
	for _, item := range setting.SingleFile.IgnoreHost {
		if len(item) > 7 && item[0:7] == "regexp:" {
			compile, err := regexp.Compile(item[7:])
			if err != nil {
				continue
			}
			ignore = compile.MatchString(parse.Host)
		} else {
			ignore = match.Match(parse.Host, item)
		}

		if ignore {
			break
		}
	}

	return ignore
}

func CheckVersionInterval(lastTime time.Time) bool {
	end := lastTime.Add(time.Duration(setting.SingleFile.MinVersionInterval) * time.Second)
	comp := time.Now().Compare(end)
	return comp == -1
}
