package createfirstpage

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"bytes"
	"comutils"
	"createfirstpagedb"
	"domains"
	"math/rand"
	"time"
	"titlegen"
	"makenewsite"
)

var index = template.Must(template.ParseFiles(
	"templ/_base_first.html",
	"templ/index_first.html",
))


func CreatePage(locale string, themes string, host string, pathinfo string, keywords []string, phrases []string) {

	var paragraphid int64
	
	htmlfile := string("www/" + locale + "/" + themes + "/" + host + pathinfo)
	log.Println(htmlfile)

	path := filepath.Dir(htmlfile)
	log.Println(path)
	err := os.MkdirAll(path, 0777)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(htmlfile)
	if err != nil {
		panic(err)
	}

	ext := string(".html")
	paragraphid,sentences := createfirstpagedb.FindFreeSentences(locale, themes)


	destkey := make([]string, len(keywords))

	rand.Seed(time.Now().UTC().UnixNano())
	permkey := rand.Perm(len(keywords))

	for i, v := range permkey {
		destkey[v] = keywords[i]
	}

	upcasesomekeywords := make([]string, len(keywords))
	for i := 0; i < len(keywords); i++ {

		upcasesomekeywords[i] = comutils.UpcaseInitial(destkey[i])
	}

	title := comutils.UpcaseInitial(titlegen.GetTitle(locale, themes, host, pathinfo))

	destphr := make([]string, len(phrases))
	rand.Seed(time.Now().UTC().UnixNano())
	permphr := rand.Perm(len(phrases))

	for i, v := range permphr {
		destphr[v] = comutils.UpcaseInitial(phrases[i])
	}

	go makenewsite.Makenew(locale,themes,host, pathinfo,title,paragraphid) 

	webinfo := domains.WebInfo{
		Site:       host,
		Ext:        ext,
		Sentences:  sentences,
		Keywords:   destkey,
		UpKeywords: upcasesomekeywords,
		Title:      title,
		Phrases:    destphr,
	}

	webpage := bytes.NewBuffer(nil)

	if err := index.Execute(webpage, webinfo); err != nil {

		panic(err)
	}
	webpagebytes := make([]byte, webpage.Len())
	webpagebytes = webpage.Bytes()

	file.Write(webpagebytes)
	file.Close()

}
