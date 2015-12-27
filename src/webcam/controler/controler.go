package controler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
)

type Data struct {
	Title    string
	Pictures []string
	Folders  []string
}

//var templates = template.Must(template.ParseFiles("template/img.tmpl"))

func RenderUI(w http.ResponseWriter, r *http.Request) {
	picturesDir := strings.Trim(os.Args[1], " ")

	li := strings.LastIndex(picturesDir, "/")

	var picturesSt string
	if li != -1 {
		picturesSt = picturesDir[li+1:]
	}

	log.Println("Recherche des images dans '", picturesSt, "'")
	files, err := ioutil.ReadDir(picturesDir)
	check(err)

	data := Data{}
	data.Title = picturesSt

	for _, f := range files {
		if f.IsDir() {
			data.Folders = append(data.Folders, f.Name())
		} else {
			data.Pictures = append(data.Pictures, f.Name())
		}
	}

	var templates = template.Must(template.ParseFiles("template/img.tmpl"))
	err = templates.Execute(w, data)
	check(err)
}

func RenderImg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	img := r.URL.Path[5:]
	ip := path.Join("data", img)

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(ip)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	file, err := os.Open(ip)
	http.ServeContent(w, r, info.Name(), info.ModTime(), file)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
