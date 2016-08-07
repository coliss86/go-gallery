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

package router

import (
	"net/http"

	"gitlab.com/coliss86/go-gallery/pkg/handler"
	"github.com/gorilla/mux"
)

var BaseDir string

func NewRouter() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/ui", handler.RenderUI)
	router.HandleFunc("/ui/{folder:.*}", handler.RenderUI)
	router.HandleFunc("/tag/{action}/{tag}", handler.ManageTag)
	router.HandleFunc("/thumb/{img:.*}", handler.RenderThumb)
	router.HandleFunc("/download/{img:.*}", handler.RenderDownload)
	router.HandleFunc("/img/{img:.*}", handler.RenderImg)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//redirecting to /ui/ when / is called
		http.Redirect(w, r, "/ui/", http.StatusMovedPermanently)
	})
	http.Handle("/", router)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(BaseDir + "/static"))))

	return router
}
