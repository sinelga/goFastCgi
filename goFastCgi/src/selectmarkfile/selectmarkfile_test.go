package selectmarkfile

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func Test_Selectmarkfile(t *testing.T) {
	//	markresources/fi_FI_finance.txt
	locale := "fi_FI"
	themes := "finance"
	var rootdir string

	if locale == "fi_FI" && themes == "finance" {

		rootdir = "markresources/" + locale + "_" + themes + "/"

	}

	t.Log(locale, themes)
	log.Println(locale, themes)

	dr, _ := ioutil.ReadDir("../../markresources/fi_FI_finance")
	//	log.Println(len(dr))
	var filesarr []string

	for _, fl := range dr {

		//		log.Println(fl.Name())
		flname := fl.Name()
		if strings.HasPrefix(flname, locale) {
			filesarr = append(filesarr, rootdir+flname)
		}
		//		log.Println(fl.	me

	}

	//	log.Println(len(filesarr))

	rand.Seed(time.Now().UnixNano())
	log.Println(rand.Intn(len(filesarr)))

	log.Println(filesarr[rand.Intn(len(filesarr))])

}
