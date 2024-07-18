package setting

import (
	"github.com/spf13/viper"
	"history-engine/engine/utils"
	"log"
	"os/exec"
)

var (
	Readability = struct {
		Parser   string
		ExecPath string
	}{
		Parser:   "mozilla",
		ExecPath: "",
	}
)

func loadReadability() {
	v := viper.Sub("readability")
	if v != nil {
		if v.IsSet("parser") {
			Readability.Parser = v.GetString("parser")
		}
		if v.IsSet("exe_path") {
			Readability.ExecPath = v.GetString("exec_path")
		}
	}

	checkReadabilityPath()
}

func checkReadabilityPath() {
	var err error
	var path string
	switch Readability.Parser {
	case "mozilla":
		path, err = exec.LookPath("readability-parse")
		if err != nil {
			log.Fatalf("check mozilla readability-parse path err:%v\n", err)
		}
	default:
		// todo
	}

	if !utils.FileExist(path) {
		log.Fatalf("readability-parse not exist\n")
	}

	log.Printf("readability parser: %s, path: %s\n", Readability.Parser, path)
}
