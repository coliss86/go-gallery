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

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Error : missing argument\n\nUsage : webcam <dataDir> <cachedir> [port]")
		os.Exit(1)
	}

	var err error
	dataDir := strings.Trim(os.Args[1], " ")
	cacheDir := strings.Trim(os.Args[2], " ")
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
	http.HandleFunc("/img/", makeHandler(RenderImg, dataDir))
	http.HandleFunc("/thumb/", makeHandlerArgs(RenderThumb, []string{dataDir, cacheDir}))
	http.HandleFunc("/", makeHandler(RenderUI, dataDir))
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	check(err)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string), s string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, s)
	}
}

func makeHandlerArgs(fn func(http.ResponseWriter, *http.Request, []string), s []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, s)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
