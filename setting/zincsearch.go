package setting

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
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
		IndexPrefix: "history_engine_index_",
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
}

func checkZincSearchVersion() {
	if ZincSearch.Host == "" {
		log.Fatalln("ZincSearch host is required")
	}

	type zincVersion struct {
		Version    string `json:"version"`
		Build      string `json:"build"`
		CommitHash string `json:"commit_hash"`
		Branch     string `json:"branch"`
		BuildDate  string `json:"build_date"`
	}

	resp, err := http.Get(ZincSearch.Host + "/version")
	if err != nil {
		log.Fatalf("check ZincSearch version err:%v\n", err)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("check ZincSearch version err:%v\n", err)
	}

	zv := zincVersion{}
	err = json.Unmarshal(content, &zv)
	if err != nil {
		log.Fatalf("check ZincSearch version err:%v\n", err)
	}

	log.Printf("zincSearch version: %s, build: %s, build_date: %s\n", zv.Version, zv.Build, zv.BuildDate)
}
