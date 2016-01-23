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
)

var config Config

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Error : missing argument\n\nUsage : ", os.Args[0], " <config file>")
		os.Exit(1)
	}

	config := NewConfiguration()

	dirExport := path.Dir(config.Export)
	os.MkdirAll(dirExport, os.ModePerm)

	log.Println("Listening on", config.Port)

	NewRouter()
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
