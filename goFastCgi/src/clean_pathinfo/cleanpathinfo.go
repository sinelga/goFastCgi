package clean_pathinfo

import (
	"log/syslog"
	"net/url"
)

func CleanPath(golog syslog.Writer,path string) string {

	u, err := url.Parse(path)
	if err != nil {
//		log.Fatal(err)
		golog.Err(err.Error())
	}

//	log.Println("u.Path", u.Path)

	cleanpath := u.Path

	var retpath string
	sz := len(cleanpath)
	if sz > 0 && cleanpath[sz-1] == '/' {
		retpath = cleanpath + "index.html"
	} else {

		retpath = cleanpath
	}

	return retpath

}
