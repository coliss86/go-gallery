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

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jimlawless/cfg"
)

const PORT = 9090

type Config struct {
	Images string
	Cache  string
	Export string
	Port   int
}

var config Config

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Error : missing argument\n\nUsage : ", os.Args[0], " <config file>")
		os.Exit(1)
	}

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
		config.Port = PORT
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

	dirExport := path.Dir(config.Export)
	os.MkdirAll(dirExport, os.ModePerm)

	log.Println("Listening on", config.Port)

	r := mux.NewRouter()
	r.HandleFunc("/ui", RenderUI)
	r.HandleFunc("/ui/{folder:.*}", RenderUI)
	r.HandleFunc("/tag/{action}/{tag}", ManageTag)
	r.HandleFunc("/thumb/{img:.*}", RenderThumb)
	r.HandleFunc("/download/{img:.*}", RenderDownload)
	r.HandleFunc("/img/{img:.*}", RenderImg)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//redirecting to /ui/ when / is called
		http.Redirect(w, r, "/ui/", http.StatusMovedPermanently)
	})
	http.Handle("/", r)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
