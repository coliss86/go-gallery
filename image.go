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
	"bytes"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"
)

var urlImgRE = regexp.MustCompile("(.*)/([^/]*)/?")

func RenderImg(w http.ResponseWriter, r *http.Request, conf Conf) {
	r.ParseForm()
	img := r.URL.Path[5:]
	serveFile(w, r, path.Join(conf.DataDir, img))
}

func RenderThumb(w http.ResponseWriter, r *http.Request, conf Conf) {
	r.ParseForm()
	thumb := r.URL.Path[7:]

	ip := path.Join(conf.DataDir, thumb)
	it := path.Join(conf.CacheDir, thumb)
	// Return a 404 if the template doesn't exist
	infoip, err := os.Stat(ip)
	if err != nil && os.IsNotExist(err) || infoip.IsDir() {
		http.NotFound(w, r)
		return
	}

	dir := path.Dir(it)
	os.MkdirAll(dir, os.ModePerm)

	infoit, err := os.Stat(it)
	if err != nil && os.IsNotExist(err) || infoip.ModTime().After(infoit.ModTime()) {
		imageMagickThumbnail(ip, it)
	}

	serveFile(w, r, it)
}

func RenderDownload(w http.ResponseWriter, r *http.Request, conf Conf) {
	r.ParseForm()
	img := r.URL.Path[10:]
	matches := urlImgRE.FindStringSubmatch(img)
	title := ""
	if len(matches) > 0 {
		title = matches[2]
	} else {
		title = img
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+title)
	serveFile(w, r, path.Join(conf.DataDir, img))
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

func imageMagickThumbnail(origName, newName string) {
	var args = []string{
		"-define", "jpeg:size=150x150",
		"-thumbnail", "100x100!",
		origName, newName,
	}

	var cmd *exec.Cmd
	path, _ := exec.LookPath("convert")
	cmd = exec.Command(path, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println("Erreur de génération : ", err, ". stdout :", stdout.String(), ". stderr :", stderr.String())
	}
}