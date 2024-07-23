package readability

import (
	"bytes"
	"encoding/json"
	"errors"
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

func (m Mozilla) Parse(path string) (*Article, error) {
	var stdErr bytes.Buffer
	cmd := exec.Command("readability-parse", path)
	cmd.Stderr = &stdErr
	data, err := cmd.Output()
	if err != nil {
		return nil, errors.Join(err, errors.New(stdErr.String()))
	}

	article := &Article{}
	err = json.Unmarshal(data, article)
	if err != nil {
		return nil, err
	}

	article.TextContent = strings.ReplaceAll(article.TextContent, "\n", "")

	return article, nil
}

func (m Mozilla) ParseContent(content []byte) (*Article, error) {
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
