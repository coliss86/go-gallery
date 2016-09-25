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

package img

import (
	"bytes"
	"log"
	"os/exec"
)

func ConvertThumbnail(origName, newName string) {
	// convert -define jpeg:size=200x200 original.jpeg  -thumbnail 100x100^ -gravity center -extent 100x100  thumbnail.jpeg
	var args = []string{
		"-auto-orient",
		"-define", "jpeg:size=150x150",
		"-thumbnail", "100x100^",
		"-gravity", "center",
		"-extent", "100x100",
		origName, newName,
	}
	convert(args)
}

func ConvertSmall(origName, newName string) {
	// convert -resize x500 original.jpeg thumbnail.jpeg
	var args = []string{
		"-auto-orient",
		"-resize", "x500",
		origName, newName,
	}
	convert(args)
}

func convert(args []string) {
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
