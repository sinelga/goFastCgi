package clean_pathinfo

import (
	"log"
	"net/url"
)

func CleanPath(path string) string {

	u, err := url.Parse(path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("u.Path", u.Path)

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
