package jswebserv

import (
	"log/syslog"
	"net/http"
	"strings"
)

func JsServ(golog syslog.Writer, w http.ResponseWriter, r *http.Request, rootdir string, host string) {

	pathstr := r.URL.Path
	useragent := r.UserAgent()
	referer := r.Referer()

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
		golog.Info("JsServ: " + host + pathstr + " " + useragent)
		urlstr := "http://"+host
		if referer != "" &&  !strings.HasPrefix(referer,urlstr) {
			golog.Info("Referer: " + referer)
		}
		http.ServeFile(w, r, rootdir+"dartapp.html")

	}

}
