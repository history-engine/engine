package setting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	ZincSearch = struct {
		Host        string
		IndexPrefix string
		SharedNum   int
		User        string
		Password    string
	}{
		Host:        "http://localhost:4080",
		IndexPrefix: "history_engine_index",
		SharedNum:   3,
		User:        "admin",
		Password:    "",
	}
)

func loadZincSearch() {
	v := viper.Sub("zincsearch")
	if v != nil {
		if v.IsSet("host") {
			ZincSearch.Host = v.GetString("host")
		}
		if v.IsSet("index_prefix") {
			ZincSearch.IndexPrefix = v.GetString("index_prefix")
		}
		if v.IsSet("shared_num") {
			ZincSearch.SharedNum = v.GetInt("shared_num")
		}
		if v.IsSet("user") {
			ZincSearch.User = v.GetString("user")
		}
		if v.IsSet("password") {
			ZincSearch.Password = v.GetString("password")
		}
	}

	checkZincSearchVersion()
	checkIndexTemplate()
}

func checkZincSearchVersion() {
	if ZincSearch.Host == "" {
		log.Fatalln("ZincSearch host is required")
	}

	resp, err := http.Get(ZincSearch.Host + "/version")
	if err != nil {
		log.Fatalf("check ZincSearch version err:%v\n", err)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("check ZincSearch version err:%v\n", err)
	}

	var zv struct {
		Version    string `json:"version"`
		Build      string `json:"build"`
		CommitHash string `json:"commit_hash"`
		Branch     string `json:"branch"`
		BuildDate  string `json:"build_date"`
	}
	err = json.Unmarshal(content, &zv)
	if err != nil {
		log.Fatalf("check ZincSearch version err:%v\n", err)
	}

	log.Printf("zincSearch version: %s, build: %s, build_date: %s\n", zv.Version, zv.Build, zv.BuildDate)
}

// https://github.com/zincsearch/zincsearch/blob/main/pkg/meta/template.go
// https://github.com/zincsearch/zincsearch/blob/main/pkg/meta/index.go
func checkIndexTemplate() {
	content, err := os.ReadFile(Common.DataPath + "/zinc_template.json")
	if err != nil {
		log.Fatalf("read zinc template err: %v\n", err)
	}

	name := fmt.Sprintf("%s_%s_%s", ZincSearch.IndexPrefix, Common.Env, "template")
	pattern := fmt.Sprintf("%s_%s_%s", ZincSearch.IndexPrefix, Common.Env, "*")
	content = bytes.Replace(content, []byte("__NAME__"), []byte(name), 1)
	content = bytes.Replace(content, []byte("__PATTERN__"), []byte(pattern), 1)

	req, err := http.NewRequest(http.MethodPost, ZincSearch.Host+"/es/_index_template", bytes.NewReader(content))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(ZincSearch.User, ZincSearch.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("create or update template err: %v\n", err)
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("create or update template err: %v\n", err)
	}

	log.Printf("zincSearch index template: %s, pattern: %s\n", name, pattern)
}
