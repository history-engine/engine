package readability

import (
	"bytes"
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
	regex, _ = regexp.Compile(`(?s)<!--.*?(htt.+://\S+).*?saved\sdate.*?-->`)
	return Mozilla{}
}

func (m Mozilla) Parse(path string) *Article {
	var stdErr bytes.Buffer
	cmd := exec.Command("readability-parse", path)
	cmd.Stderr = &stdErr
	data, err := cmd.Output()
	if err != nil {
		logger.Zap().Error("exec readability-parse err", zap.String("stderr", stdErr.String()))
		return nil
	}

	article := &Article{}
	err = json.Unmarshal(data, article)
	if err != nil {
		logger.Zap().Error("unmarshal readability output err", zap.Error(err))
		return nil
	}

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
