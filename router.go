package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/ui", RenderUI)
	router.HandleFunc("/ui/{folder:.*}", RenderUI)
	router.HandleFunc("/tag/{action}/{tag}", ManageTag)
	router.HandleFunc("/thumb/{img:.*}", RenderThumb)
	router.HandleFunc("/download/{img:.*}", RenderDownload)
	router.HandleFunc("/img/{img:.*}", RenderImg)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//redirecting to /ui/ when / is called
		http.Redirect(w, r, "/ui/", http.StatusMovedPermanently)
	})
	http.Handle("/", router)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	return router
}
