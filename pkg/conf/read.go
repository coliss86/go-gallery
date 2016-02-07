/*
This file is part of GO gallery.

GO gallery is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GO gallery is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with GO gallery.  If not, see <http://www.gnu.org/licenses/>.
*/

package conf

import (
	"log"
	"os"
	"strconv"
	"strings"
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
	err := Load(os.Args[1], meta)
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
