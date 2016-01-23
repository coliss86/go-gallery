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
	"path"

	"github.com/gmembre/go-gallery/pkg/conf"
	"github.com/gmembre/go-gallery/pkg/file"
	"github.com/gorilla/mux"
)

func ManageTag(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)
	action := vars["action"]
	tag := vars["tag"]
	img := r.FormValue("img")
	switch {
	case action == "select":
		tagSelect(img, tag)
	case action == "deselect":
		tagDeselect(img, tag)
	case action == "add":
		tagAdd(img, tag)
	case action == "delete":
		tagDelete(img, tag)
	}
}

func tagSelect(img string, tag string) {
	log.Println("Copy", img, "to", tag)
	src := path.Join(conf.Config.Images, img)
	dest := path.Join(conf.Config.Export, tag, path.Base(img))
	err := file.CopyFile(src, dest)
	if err != nil {
		log.Println("Copy error", src, "to", dest, ":", err)
	}
}

func tagDeselect(img string, tag string) {
	log.Println("Delete", img, "from", tag)
	dest := path.Join(conf.Config.Export, tag, path.Base(img))
	os.Remove(dest)
}

func tagAdd(img string, tag string) {
	log.Println("Create folder", tag)
	os.MkdirAll(path.Join(conf.Config.Export, tag), os.ModePerm)
}

func tagDelete(img string, tag string) {
	log.Println("Delete folder", tag)
	err := os.RemoveAll(path.Join(conf.Config.Export, tag))
	check(err)
}

func tagList() (tags []string) {
	files, err := ioutil.ReadDir(conf.Config.Export)
	check(err)

	tags = make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			tags = append(tags, file.Name())
		}
	}
	return
}

func tagListPictures(tag string) (files []string) {
	fs, err := ioutil.ReadDir(path.Join(conf.Config.Export, tag))
	check(err)

	files = make([]string, 0)
	for _, file := range fs {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
