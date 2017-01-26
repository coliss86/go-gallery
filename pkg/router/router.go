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

	"github.com/gorilla/mux"
	"gitlab.com/coliss86/go-gallery/pkg/conf"
	"gitlab.com/coliss86/go-gallery/pkg/handler"
)

var BaseDir string

func NewRouter() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc(conf.Config.ContextRoot+"/ui", handler.RenderUI)
	router.HandleFunc(conf.Config.ContextRoot+"/ui/{folder:.*}", handler.RenderUI)
	router.HandleFunc(conf.Config.ContextRoot+"/tag/{action}/{tag}", handler.ManageTag)
	router.HandleFunc(conf.Config.ContextRoot+"/thumb/{img:.*}", handler.RenderThumb)
	router.HandleFunc(conf.Config.ContextRoot+"/small/{img:.*}", handler.RenderSmall)
	router.HandleFunc(conf.Config.ContextRoot+"/download/{img:.*}", handler.RenderDownload)
	router.HandleFunc(conf.Config.ContextRoot+"/img/{img:.*}", handler.RenderImg)
	router.HandleFunc(conf.Config.ContextRoot+"/", func(w http.ResponseWriter, r *http.Request) {
		//redirecting to /ui/ when / is called
		http.Redirect(w, r, conf.Config.ContextRoot+"/ui/", http.StatusMovedPermanently)
	})
	http.Handle(conf.Config.ContextRoot+"/", router)

	http.Handle(conf.Config.ContextRoot+"/static/", http.StripPrefix(conf.Config.ContextRoot+"/static/", http.FileServer(http.Dir(BaseDir+"/static"))))

	return router
}
