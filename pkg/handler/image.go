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
	"net/http"
	"os"
	"path"
	"regexp"

	"github.com/gmembre/go-gallery/pkg/conf"
	"github.com/gmembre/go-gallery/pkg/img"
	"github.com/gorilla/mux"
)

var urlImgRE = regexp.MustCompile("(.*)/([^/]*)/?")

func RenderImg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)
	serveFile(w, r, path.Join(conf.Config.Images, vars["img"]))
}

func RenderThumb(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)
	thumb := vars["img"]

	ir := path.Join(conf.Config.Images, thumb)
	ic := path.Join(conf.Config.Cache, thumb)
	// Return a 404 if the template doesn't exist
	infoir, err := os.Stat(ir)
	if err != nil && os.IsNotExist(err) || infoir.IsDir() {
		http.NotFound(w, r)
		return
	}

	dir := path.Dir(ic)
	os.MkdirAll(dir, os.ModePerm)

	infoic, err := os.Stat(ic)
	if err != nil && os.IsNotExist(err) || infoir.ModTime().After(infoic.ModTime()) {
		img.ConvertThumbnail(ir, ic)
	}

	serveFile(w, r, ic)
}

func RenderDownload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)
	img := vars["img"]
	matches := urlImgRE.FindStringSubmatch(img)
	title := ""
	if len(matches) > 0 {
		title = matches[2]
	} else {
		title = img
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+title)
	serveFile(w, r, path.Join(conf.Config.Images, img))
}

func serveFile(w http.ResponseWriter, r *http.Request, file string) {
	// Return a 404 if the template doesn't exist
	info, err := os.Stat(file)
	if err != nil && os.IsNotExist(err) || info.IsDir() {
		http.NotFound(w, r)
		return
	}

	fileOs, err := os.Open(file)
	defer fileOs.Close()
	http.ServeContent(w, r, info.Name(), info.ModTime(), fileOs)
}
