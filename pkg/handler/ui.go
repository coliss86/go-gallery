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

package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
	"gitlab.com/coliss86/go-gallery/pkg/conf"
	"gitlab.com/coliss86/go-gallery/pkg/file"
)

type Item struct {
	Link     string
	Name     string
	LongName string
	Thumb    string
	Class    string
}

type Data struct {
	Title        string
	Breadcrum    []Item
	Pictures     []string
	Videos       []string
	Folders      map[string][]Item
	Months       []string
	MonthsName   map[string]string
	Folder       string
	TagsPictures map[string][]string
	Total        int
}

var monthsName = map[string]string{"01": "Janvier", "02": "Février", "03": "Mars", "04": "Avril", "05": "Mai", "06": "Juin", "07": "Juillet", "08": "Août", "09": "Septembre", "10": "Octobre", "11": "Novembre", "12": "Décembre", "": "Dossiers"}

var folderRE = regexp.MustCompile("^([0-9]+)-([0-9]+)(_(.*))?$")
var folderExcludeRE = regexp.MustCompile("(?i)\\.Trash.*")
var pictureRE = regexp.MustCompile("(?i).*\\.(jpeg|jpg|gif|png|bmp)$")
var urlFolderRE = regexp.MustCompile("(.*)/([^/]*)/?")
var videoRE = regexp.MustCompile("(?i).*(mp4|m4v|mpeg|mpg|avi)$")

//var templates = template.Must(template.ParseFiles("template/img.tmpl"))

func RenderUI(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)
	folderS := vars["folder"]
	folder := file.PathJoin(conf.Config.Images, folderS)
	log.Println("Listing pictures in '", folder, "'")

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
					item.Link = split
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

	// pictures
	data.Folders = make(map[string][]Item)
	for _, file := range files {
		if file.IsDir() && !folderExcludeRE.MatchString(file.Name()) {
			manageFolder(folder, file, data)
		} else if videoRE.MatchString(file.Name()) {
			data.Videos = append(data.Videos, file.Name())
		} else if pictureRE.MatchString(file.Name()) {
			data.Pictures = append(data.Pictures, file.Name())
		}
	}

	// months
	data.Months = make([]string, len(data.Folders))
	i := 0
	for k, _ := range data.Folders {
		data.Months[i] = k
		i++
	}
	sort.Strings(data.Months)

	// tags
	data.TagsPictures = make(map[string][]string)
	tags := tagList()
	for _, t := range tags {
		data.TagsPictures[t] = tagListPictures(t)
	}

	data.Total = len(data.Videos) + len(data.Pictures)

	// final generation
	var templates = template.Must(template.ParseFiles("template/gallery.tmpl"))
	err = templates.Execute(w, data)
	check(err)
}

func manageFolder(parent string, folder os.FileInfo, data Data) {
	f := Item{}
	month := ""
	f.Link = folder.Name()
	matches := folderRE.FindStringSubmatch(f.Link)
	if len(matches) > 0 {
		f.Name = matches[2]
		month = matches[1]
		f.Class = "folder-short"
		if len(matches) == 5 {
			f.LongName = matches[4]
			if f.LongName == "" {
				f.LongName = f.Link
			}
		} else {
			f.LongName = f.Link
		}
		meta := readMeta(file.PathJoin(parent, f.Link))
		if meta != nil {
			f.LongName = meta["title"]
			f.LongName = strings.Replace(f.LongName, " ", "&nbsp;", -1)
		}
	} else {
		f.Name = f.Link
		f.LongName = f.Name
		if len(f.Name) > 7 {
			f.Class = "folder-long"
		} else {
			f.Class = "folder-normal"
		}
	}

	// thumb folder
	files, err := ioutil.ReadDir(file.PathJoin(parent, folder.Name()))
	check(err)

	for _, file := range files {
		if !file.IsDir() && pictureRE.MatchString(file.Name()) {
			f.Thumb = file.Name()
			f.Class += " folder-thumb"
			break
		}
	}

	// month
	v, ok := data.Folders[month]
	if !ok {
		data.Folders[month] = make([]Item, 5)
	}
	data.Folders[month] = append(v, f)
}

func readMeta(folder string) map[string]string {
	path := file.PathJoin(folder, "meta.properties")
	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		return nil
	}
	meta := make(map[string]string)
	conf.Load(path, meta)
	return meta
}
