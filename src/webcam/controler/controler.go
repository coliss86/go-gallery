package controler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"
)

type Folder struct {
	Link string
	Name string
}

type Data struct {
	Title    string
	Pictures []string
	Folders  []Folder
}

var re = regexp.MustCompile("[0-9]+-([0-9]+).*")

//var templates = template.Must(template.ParseFiles("template/img.tmpl"))

func RenderUI(w http.ResponseWriter, r *http.Request, dataDir string) {
	r.ParseForm()
	folder := path.Join(dataDir, r.URL.Path[1:])
	log.Println("Recherche des images dans '", folder, "'")

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(folder)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if !info.IsDir() {
		http.NotFound(w, r)
		return
	}

	var picturesSt string
	li := strings.LastIndex(folder, "/")
	if li != -1 {
		picturesSt = folder[li+1:]
	}
	data := Data{}
	data.Title = picturesSt

	files, err := ioutil.ReadDir(folder)
	check(err)

	for _, file := range files {
		if file.IsDir() {
			f := Folder{}
			f.Link = file.Name()
			matches := re.FindStringSubmatch(f.Link)
			if len(matches) > 0 {
				f.Name = matches[1]
			} else {
				f.Name = f.Link
			}

			data.Folders = append(data.Folders, f)
		} else {
			data.Pictures = append(data.Pictures, file.Name())
		}
	}

	var templates = template.Must(template.ParseFiles("template/img.tmpl"))
	err = templates.Execute(w, data)
	check(err)
}

func RenderImg(w http.ResponseWriter, r *http.Request, dataDir string) {
	r.ParseForm()
	img := r.URL.Path[5:]
	serveFile(w, r, dataDir, img)
}

func RenderThumb(w http.ResponseWriter, r *http.Request, dataDir string) {
	r.ParseForm()
	thumb := r.URL.Path[7:]
	serveFile(w, r, dataDir, thumb)
}

func serveFile(w http.ResponseWriter, r *http.Request, dataDir string, file string) {
	ip := path.Join(dataDir, file)

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

	fileOs, err := os.Open(ip)
	http.ServeContent(w, r, info.Name(), info.ModTime(), fileOs)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
