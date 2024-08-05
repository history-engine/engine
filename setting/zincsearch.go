package setting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"history-engine/engine/data"
	"io"
	"log"
	"net/http"
)

var ZincSearch = struct {
	Host      string
	SharedNum int
	User      string
	Password  string
}{
	Host:      "http://localhost:4080",
	SharedNum: 3,
	User:      "admin",
	Password:  "",
}

func loadZincSearch() {
	v := viper.Sub("zincsearch")
	if v != nil {
		if v.IsSet("host") {
			ZincSearch.Host = v.GetString("host")
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

	if Search.Engine == "zinc" {
		checkZincSearchVersion()
		checkIndexTemplate()
	}
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
	name := fmt.Sprintf("%s_%s_%s", Search.Prefix, Common.Env, "template")
	pattern := fmt.Sprintf("%s_%s_%s", Search.Prefix, Common.Env, "*")
	content := bytes.Replace(data.ZincTemplate, []byte("__NAME__"), []byte(name), 1)
	content = bytes.Replace(content, []byte("__PATTERN__"), []byte(pattern), 1)

	req, err := http.NewRequest(http.MethodPost, ZincSearch.Host+"/es/_index_template", bytes.NewReader(content))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(ZincSearch.User, ZincSearch.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("create or update template err: %v\n", err)
	}
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("create or update template err: %v\n", err)
	}

	var tr struct {
		Error    string `json:"error"`
		Message  string `json:"message"`
		Template string `json:"template"`
	}
	if err = json.Unmarshal(res, &tr); err != nil || tr.Error != "" {
		log.Fatalf("create or update template reason: %s, err: %v\n", tr.Error, err)
	}

	log.Printf("zincSearch index template: %s, pattern: %s\n", name, pattern)
}
