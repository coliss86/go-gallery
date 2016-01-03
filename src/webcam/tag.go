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
	"strings"
)

func ManageTag(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	action := r.URL.Path[5:]
	img := r.FormValue("img")
	actions := strings.Split(action, "/")
	switch {
	case actions[0] == "select":
		tagSelect(img, actions[1])
	case actions[0] == "deselect":
		tagDeselect(img, actions[1])
	case actions[0] == "add":
		tagAdd(img, actions[1])
	case actions[0] == "delete":
		tagDelete(img, actions[1])
	}
}

func tagSelect(img string, tag string) {
	log.Println("Copy", img, "to", tag)
	src := path.Join(conf.DataDir, img)
	dest := path.Join(conf.ExportDir, tag, path.Base(img))
	err := CopyFile(src, dest)
	if err != nil {
		log.Println("Copy error", src, "to", dest, ":", err)
	}
}

func tagDeselect(img string, tag string) {
	log.Println("Delete", img, "from", tag)
	dest := path.Join(conf.ExportDir, tag, path.Base(img))
	os.Remove(dest)
}

func tagAdd(img string, tag string) {
	log.Println("Create folder", tag)
	os.MkdirAll(path.Join(conf.ExportDir, tag), os.ModePerm)
}

func tagDelete(img string, tag string) {
	log.Println("Delete folder", tag, " NON IMPLEMENTED !!")
}

func tagList() (tags []string) {
	log.Println("Listing tag")
	files, err := ioutil.ReadDir(conf.ExportDir)
	check(err)

	tags = make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			tags = append(tags, file.Name())
		}
	}
	return
}
