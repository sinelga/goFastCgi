package jswebserv

import (
//"net"
"log/syslog"
"net/http"
"strings"
)

func JsServ(golog syslog.Writer,w http.ResponseWriter, r *http.Request,rootdir string){

	golog.Info("Start JS rootdir "+rootdir)
	pathstr := r.URL.Path

	if strings.HasSuffix(pathstr, ".js") || strings.HasSuffix(pathstr, ".css") {

		if strings.Index(pathstr, "/packages/") == -1 {

			splitpatharr := strings.Split(pathstr, "/")

			if len(splitpatharr) > 2 {

				pathstr = "/" + splitpatharr[len(splitpatharr)-1]
//				log.Println("css  !!" + pathstr)

			}

		} else {

//			log.Println("cut pathstr !!" + pathstr)
			pathstr = pathstr[strings.Index(pathstr, "/packages/"):]
//			log.Println("resalt !!" + pathstr)

		}

		http.ServeFile(w, r, rootdir+pathstr)
	} else {
		http.ServeFile(w, r, rootdir+"dartapp.html")

	}
		
}