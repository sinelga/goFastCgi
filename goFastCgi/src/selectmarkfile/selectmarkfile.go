package selectmarkfile

import (
	"io/ioutil"
	"log"
	"log/syslog"
	"math/rand"
	"strings"
	"time"
)

func SelectFile(golog syslog.Writer, locale string, themes string) string {

	log.Println(locale, themes)

	var rootdir string
	var filenameret string

	rootdir = "markresources/" + locale + "_" + themes + "/"

	if dr, err := ioutil.ReadDir(rootdir); err != nil {

		golog.Crit(err.Error())

	} else {

		var filesarr []string

		for _, fl := range dr {

			flname := fl.Name()
			if strings.HasPrefix(flname, locale) {
				filesarr = append(filesarr, rootdir+flname)
			}

		}
		rand.Seed(time.Now().UnixNano())

		filenameret = filesarr[rand.Intn(len(filesarr))]

	}
	return filenameret
}
