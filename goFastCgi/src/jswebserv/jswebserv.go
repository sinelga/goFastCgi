package jswebserv

import (

"log/syslog"
"net/http"
"strings"
)

func JsServ(golog syslog.Writer,w http.ResponseWriter, r *http.Request,rootdir string){

	pathstr := r.URL.Path
	useragent :=r.UserAgent()
//	golog.Info("Start JS  pathstr "+pathstr)
	
	if strings.HasSuffix(pathstr, ".js") || strings.HasSuffix(pathstr, ".css") {

		if strings.Index(pathstr, "/packages/") == -1 {

			splitpatharr := strings.Split(pathstr, "/")

			if len(splitpatharr) > 2 {

				pathstr = "/" + splitpatharr[len(splitpatharr)-1]

			}

		} else {

			pathstr = pathstr[strings.Index(pathstr, "/packages/"):]

		}

		http.ServeFile(w, r, rootdir+pathstr)
	} else {
		golog.Info("JsServ: "+pathstr+" "+useragent)
//		golog.Info(useragent)
		http.ServeFile(w, r, rootdir+"dartapp.html")

	}
		
}