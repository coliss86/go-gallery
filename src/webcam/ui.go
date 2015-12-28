package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

type Folder struct {
	Link string
	Name string
}

type Data struct {
	Title      string
	Breadcrum  []string
	Pictures   []string
	Values     map[string][]Folder
	Months     []string
	MonthsName map[string]string
}

var monthsName = map[string]string{"01": "Janvier", "02": "Février", "03": "Mars", "04": "Avril", "05": "Mai", "06": "Juin", "07": "Juillet", "08": "Aout", "09": "Septembre", "10": "Octobre", "11": "Novembre", "12": "Décembre"}

var folderRE = regexp.MustCompile("([0-9]+)-([0-9]+).*")

var ignoreRE = regexp.MustCompile(`.git|.svn|.DS_Store|Thumbs.db`)

//var templates = template.Must(template.ParseFiles("template/img.tmpl"))

func RenderUI(w http.ResponseWriter, r *http.Request, dataDir string) {
	r.ParseForm()
	folderS := r.URL.Path[1:]
	folder := path.Join(dataDir, folderS)
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

	title := ""
	li := strings.LastIndex(folderS, "/")
	if li != -1 {
		title = folderS[li+1:]
	} else {
		title = folderS
	}
	data := Data{}
	data.Title = title
	data.Values = make(map[string][]Folder)
	data.MonthsName = monthsName
	if len(folderS) > 0 {
		data.Breadcrum = strings.Split(folderS, "/")
	}
	files, err := ioutil.ReadDir(folder)
	check(err)

	for _, file := range files {
		if file.IsDir() {
			f := Folder{}
			month := ""
			f.Link = file.Name()
			matches := folderRE.FindStringSubmatch(f.Link)
			if len(matches) > 0 {
				f.Name = matches[2]
				month = matches[1]
			} else {
				f.Name = f.Link
			}

			v, ok := data.Values[month]
			if !ok {
				data.Values[month] = make([]Folder, 5)
			}
			data.Values[month] = append(v, f)
		} else if !ignoreRE.MatchString(file.Name()) {
			data.Pictures = append(data.Pictures, file.Name())
		}
	}

	// extraction des mois
	data.Months = make([]string, len(data.Values))
	i := 0
	for k, _ := range data.Values {
		data.Months[i] = k
		i++
	}
	sort.Strings(data.Months)

	// generation finale
	var templates = template.Must(template.ParseFiles("template/img.tmpl"))
	err = templates.Execute(w, data)
	check(err)
}
