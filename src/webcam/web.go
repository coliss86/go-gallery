package main

import (
	"log"
	"net/http"
	"webcam/controler"
)

func main() {
	log.Println("Listening...")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/img/", makeHandler(controler.RenderImg))
	http.HandleFunc("/", makeHandler(controler.RenderUI))
	err := http.ListenAndServe(":9090", nil)
	check(err)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//log.Println(r.URL.Path)
		fn(w, r)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
