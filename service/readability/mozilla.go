package readability

import (
	"encoding/json"
	"go.uber.org/zap"
	"history-engine/engine/library/logger"
	"os/exec"
	"regexp"
	"strings"
)

var (
	regex *regexp.Regexp
	_     Readability = Mozilla{}
)

type Mozilla struct{}

func NewMozilla() Readability {
	_, err := exec.LookPath("readability-parse")
	if err != nil {
		panic("readability-parse not exist")
	}

	regex, _ = regexp.Compile(`(?s)<!--.*?(htt.+://\S+).*?saved\sdate.*?-->`)

	return Mozilla{}
}

func (m Mozilla) Parse(path string) *Article {
	cmd := exec.Command("readability-parse", path)
	data, err := cmd.Output()
	if err != nil {
		logger.Zap().Error("exec readability-parse err", zap.Error(err))
	}

	article := &Article{}
	err = json.Unmarshal(data, article)
	if err != nil {
		panic(err) // todo 使用等级日志
	}

	//article.Url = ExtractSingleFileUrl()

	article.TextContent = strings.ReplaceAll(article.TextContent, "\n", "")

	return article
}

func (m Mozilla) ParseContent(content []byte) *Article {
	//TODO implement me
	panic("implement me")
}

func (m Mozilla) ExtractSingleFileUrl(content []byte) string {
	matches := regex.FindStringSubmatch(string(content))
	if len(matches) == 2 {
		return matches[1]
	}
	return ""
}
