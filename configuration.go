package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jimlawless/cfg"
)

const DEFAULT_PORT = 9090

type Config struct {
	Images string
	Cache  string
	Export string
	Port   int
}

func NewConfiguration() Config {
	meta := make(map[string]string)
	err := cfg.Load(os.Args[1], meta)
	check(err)
	config.Export = meta["export"]
	config.Cache = meta["cache"]
	config.Images = meta["images"]
	port, ok := meta["port"]
	if ok {
		config.Port, err = strconv.Atoi(port)
		check(err)
	}

	if config.Export == "" {
		config.Export = config.Images + "/export/"
	}
	if config.Port == 0 {
		config.Port = DEFAULT_PORT
	}

	if !strings.HasSuffix(config.Images, "/") {
		config.Images += "/"
	}
	if !strings.HasSuffix(config.Export, "/") {
		config.Export += "/"
	}
	if !strings.HasSuffix(config.Cache, "/") {
		config.Cache += "/"
	}

	log.Println("Images directory", config.Images)
	log.Println("Export directory", config.Export)

	return config
}
