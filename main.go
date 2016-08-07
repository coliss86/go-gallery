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
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
    "path/filepath"

	"gitlab.com/coliss86/go-gallery/pkg/conf"
	"gitlab.com/coliss86/go-gallery/pkg/handler"
	"gitlab.com/coliss86/go-gallery/pkg/router"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Error : missing argument\n\nUsage : ", os.Args[0], " <config file>")
		os.Exit(1)
	}

	conf.Read()

	dirExport := path.Dir(conf.Config.Export)
	os.MkdirAll(dirExport, os.ModePerm)

    // BaseDir = folder where the binary is stored, all ressources are accessed relativ to this folder
    baseDir, errb := filepath.Abs(filepath.Dir(os.Args[0]))
    check(errb)
    router.BaseDir = baseDir
    handler.BaseDir = baseDir

	router.NewRouter()

	log.Println("Listening on", conf.Config.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Config.Port), nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
