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
