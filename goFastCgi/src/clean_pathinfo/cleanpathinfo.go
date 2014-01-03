package clean_pathinfo

import (
	"log/syslog"
	"net/url"
	"strings"
)

func CleanPath(golog syslog.Writer,path string) string {

	u, err := url.Parse(path)
	if err != nil {

		golog.Err(err.Error())
	}

	cleanpath := u.Path

	var retpath string
	sz := len(cleanpath)
	if sz > 0 && cleanpath[sz-1] == '/' {
		retpath = cleanpath + "index.html"
	} else if strings.Index(cleanpath, ".") == -1 {
		retpath = cleanpath + "/index.html"
		
	} else {
	
		retpath = cleanpath
	
	}

	return retpath

}
