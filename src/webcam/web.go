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
	"strconv"
	"strings"
)

const PORT = 9090

type Conf struct {
	DataDir  string
	CacheDir string
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Error : missing argument\n\nUsage : webcam <dataDir> <cachedir> [port]")
		os.Exit(1)
	}

	var err error
	conf := Conf{strings.Trim(os.Args[1], " "), strings.Trim(os.Args[2], " ")}
	port := PORT

	if len(os.Args) == 4 {
		port, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error : invalid port number")
			os.Exit(1)
		}
	}

	log.Println("Listening on", port)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/img/", makeHandler(RenderImg, conf))
	http.HandleFunc("/download/", makeHandler(RenderDownload, conf))
	http.HandleFunc("/thumb/", makeHandler(RenderThumb, conf))
	http.HandleFunc("/", makeHandler(RenderUI, conf))
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	check(err)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, Conf), conf Conf) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, conf)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
