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

type Item struct {
	Link  string
	Name  string
	Image string
}

type Data struct {
	Title      string
	Breadcrum  []Item
	Pictures   []string
	Values     map[string][]Item
	Months     []string
	MonthsName map[string]string
	Folder     string
	Tags       string
}

var monthsName = map[string]string{"01": "Janvier", "02": "Février", "03": "Mars", "04": "Avril", "05": "Mai", "06": "Juin", "07": "Juillet", "08": "Aout", "09": "Septembre", "10": "Octobre", "11": "Novembre", "12": "Décembre", "": "Dossier"}

var folderRE = regexp.MustCompile("([0-9]+)-([0-9]+).*")
var ignoreRE = regexp.MustCompile(`.git|.svn|.DS_Store|Thumbs.db|meta.properties`)
var urlFolderRE = regexp.MustCompile("(.*)/([^/]*)/?")

//var templates = template.Must(template.ParseFiles("template/img.tmpl"))

func RenderUI(w http.ResponseWriter, r *http.Request, conf Conf) {
	r.ParseForm()
	folderS := r.URL.Path[1:]
	folder := path.Join(conf.DataDir, folderS)
	log.Println("Recherche des images dans '", folder, "'")

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(folder)
	if err != nil && os.IsNotExist(err) || !info.IsDir() {
		http.NotFound(w, r)
		return
	}

	data := Data{}

	// title
	title := ""
	matches := urlFolderRE.FindStringSubmatch(folderS)
	if len(matches) > 0 {
		title = matches[2]
	} else {
		title = folderS
	}
	data.Title = title
	data.MonthsName = monthsName
	if strings.HasSuffix(folderS, "/") {
		data.Folder = folderS
	} else if len(folderS) > 0 {
		data.Folder = folderS + "/"
	}

	// breadcrum
	if len(folderS) > 0 {
		splits := strings.Split(folderS, "/")
		var breadcrum []Item
		for i, split := range splits {
			if len(split) > 0 {
				item := Item{}
				item.Name = split
				if i == 0 {
					item.Link = "/" + split
				} else {
					item.Link = breadcrum[i-1].Link + "/" + split
				}
				breadcrum = append(breadcrum, item)
			}
		}
		breadcrum[len(breadcrum)-1].Link = ""

		data.Breadcrum = breadcrum
	}
	files, err := ioutil.ReadDir(folder)
	check(err)

	// gestion des photos
	data.Values = make(map[string][]Item)
	for _, file := range files {
		if file.IsDir() {
			manageFolder(folder, file, data)
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

	// récupération des tags
	tags := tagList()
	tag := ""
	for i, t := range tags {
		if i == 0 {
			tag += `"` + t + `"`
		} else {
			tag += `,"` + t + `"`
		}
	}
	data.Tags = tag

	// generation finale
	var templates = template.Must(template.ParseFiles("template/img.tmpl"))
	err = templates.Execute(w, data)
	check(err)
}

func manageFolder(folder string, file os.FileInfo, data Data) {
	f := Item{}
	month := ""
	f.Link = file.Name()
	matches := folderRE.FindStringSubmatch(f.Link)
	if len(matches) > 0 {
		f.Name = matches[2]
		month = matches[1]
	} else {
		f.Name = f.Link
	}

	// thumb folder
	files, err := ioutil.ReadDir(path.Join(folder, file.Name()))
	check(err)

	for _, file := range files {
		if !file.IsDir() && !ignoreRE.MatchString(file.Name()) {
			f.Image = file.Name()
			break
		}
	}

	// month
	v, ok := data.Values[month]
	if !ok {
		data.Values[month] = make([]Item, 5)
	}
	data.Values[month] = append(v, f)

}
