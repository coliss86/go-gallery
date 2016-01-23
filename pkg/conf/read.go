package conf

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jimlawless/cfg"
)

const DEFAULT_PORT = 9090

type Configuration struct {
	Images string
	Cache  string
	Export string
	Port   int
}

var Config Configuration

func Read() Configuration {
	meta := make(map[string]string)
	err := cfg.Load(os.Args[1], meta)
	check(err)

	Config.Export = meta["export"]
	Config.Cache = meta["cache"]
	Config.Images = meta["images"]
	port, ok := meta["port"]
	if ok {
		Config.Port, err = strconv.Atoi(port)
		check(err)
	}

	if Config.Export == "" {
		Config.Export = Config.Images + "/export/"
	}
	if Config.Port == 0 {
		Config.Port = DEFAULT_PORT
	}

	if !strings.HasSuffix(Config.Images, "/") {
		Config.Images += "/"
	}
	if !strings.HasSuffix(Config.Export, "/") {
		Config.Export += "/"
	}
	if !strings.HasSuffix(Config.Cache, "/") {
		Config.Cache += "/"
	}

	log.Println("Images directory", Config.Images)
	log.Println("Export directory", Config.Export)

	return Config
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
