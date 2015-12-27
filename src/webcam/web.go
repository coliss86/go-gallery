package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"webcam/controler"
)

const PORT = 9090

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Error : missing argument\n\nUsage : webcam <dataDir> [port]")
		os.Exit(1)
	}

	var err error

	dataDir := strings.Trim(os.Args[1], " ")
	port := PORT

	if len(os.Args) == 3 {
		port, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error : invalid port number")
			os.Exit(1)
		}
	}

	log.Println("Listening on", port)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/img/", makeHandler(controler.RenderImg, dataDir))
	http.HandleFunc("/thumb/", makeHandler(controler.RenderThumb, dataDir))
	http.HandleFunc("/", makeHandler(controler.RenderUI, dataDir))
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	check(err)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string), s string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, s)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
