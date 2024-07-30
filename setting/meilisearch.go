package setting

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

var MeiliSearch = struct {
	Host      string
	MasterKey string
}{
	Host:      "http://localhost:7070",
	MasterKey: "",
}

func loadMeiliSearch() {
	v := viper.Sub("meilisearch")
	if v != nil {
		if v.IsSet("host") {
			MeiliSearch.Host = v.GetString("host")
		}
		if v.IsSet("master_key") {
			MeiliSearch.MasterKey = v.GetString("master_key")
		}
	}
	if Search.Engine == "meili" {
		checkMeiliSearchVersion()
	}
}

func checkMeiliSearchVersion() {
	if MeiliSearch.Host == "" {
		log.Fatalln("MeiliSearch host is required")
	}

	if MeiliSearch.MasterKey == "" {
		log.Fatalln("MeiliSearch master key is required")
	}

	req, err := http.NewRequest(http.MethodGet, MeiliSearch.Host+"/version", nil)
	if err != nil {
		log.Fatalf("check MeiliSearch version err:%v\n", err)
	}

	req.Header.Set("Authorization", "Bearer "+MeiliSearch.MasterKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("check MeiliSearch version err: %v\n", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("check MeiliSearch version err: %d\n", resp.StatusCode)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("check MeiliSearch version err: %v\n", err)
	}

	var mv struct {
		CommitSha  string `json:"commitSha"`
		CommitDate string `json:"commitDate"`
		PkgVersion string `json:"pkgVersion"`
	}
	err = json.Unmarshal(content, &mv)
	if err != nil {
		log.Fatalf("check MeiliSearch version err: %v\n", err)
	}

	log.Printf("meilisearch version: %s, sha: %s, date: %s\n", mv.PkgVersion, mv.CommitSha, mv.CommitDate)
}
