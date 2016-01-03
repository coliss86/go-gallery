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
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
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
	log.Println("copy", img, "to", tag)
	src := path.Join(conf.DataDir, img)
	dest := path.Join(conf.ExportDir, tag, path.Base(img))
	err := copyFile(src, dest)
	if err != nil {
		log.Println("Copy error", src, "to", dest, ":", err)
	}
}

func tagDeselect(img string, tag string) {
	log.Println("delete", img, "from", tag)
	dest := path.Join(conf.ExportDir, tag, path.Base(img))
	os.Remove(dest)
}

func tagAdd(img string, tag string) {
	log.Println("create folder", tag)
	os.MkdirAll(path.Join(conf.ExportDir, tag), os.ModePerm)
}

func tagDelete(img string, tag string) {
	log.Println("delete folder", tag, " NON IMPLEMENTED !!")
}

func tagList() (tags []string) {
	log.Println("listing tag")
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

// Copies file source to destination dest.
func copyFile(source string, dest string) (err error) {
	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()

	_, err = io.Copy(df, sf)
	if err != nil {
		return err
	}

	err = df.Close()
	if err != nil {
		return err
	}
	si, err := os.Stat(source)
	if err != nil {
		return err
	}
	err = os.Chmod(dest, si.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(dest, time.Now(), si.ModTime())
	return
}
