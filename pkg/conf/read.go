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

func Read() Configuration {
	meta := make(map[string]string)
	err := cfg.Load(os.Args[1], meta)
	check(err)

	var configuration Configuration

	configuration.Export = meta["export"]
	configuration.Cache = meta["cache"]
	configuration.Images = meta["images"]
	port, ok := meta["port"]
	if ok {
		configuration.Port, err = strconv.Atoi(port)
		check(err)
	}

	if configuration.Export == "" {
		configuration.Export = configuration.Images + "/export/"
	}
	if configuration.Port == 0 {
		configuration.Port = DEFAULT_PORT
	}

	if !strings.HasSuffix(configuration.Images, "/") {
		configuration.Images += "/"
	}
	if !strings.HasSuffix(configuration.Export, "/") {
		configuration.Export += "/"
	}
	if !strings.HasSuffix(configuration.Cache, "/") {
		configuration.Cache += "/"
	}

	log.Println("Images directory", configuration.Images)
	log.Println("Export directory", configuration.Export)

	return configuration
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
