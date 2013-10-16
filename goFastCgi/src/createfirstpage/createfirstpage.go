package createfirstpage

import (
	"bytes"
	"comutils"
	//	"createfirstpagedb"
	"domains"
	"html/template"
	"log"
	"queue/makenewsite"

	"math/rand"
	"os"
	"path/filepath"
	"queue/findfreeparagraph"
	"time"
	"titlegen"
)

func CreatePage(locale string, themes string, host string, pathinfo string, keywords []string, phrases []string) {

	var index = template.Must(template.ParseFiles(
		"templ/_base_first.html",
		"templ/index_first.html",
	))

	//	var paragraphid int64

	htmlfile := string("www/" + locale + "/" + themes + "/" + host + pathinfo)

	path := filepath.Dir(htmlfile)
	log.Println(path)
	err := os.MkdirAll(path, 0777)
	if err != nil {
		//		panic(err)
		log.Fatal(err)
	}

	file, err := os.Create(htmlfile)
	if err != nil {
		//		panic(err)
		log.Fatal(err)
	}

	ext := string(".html")
	//	paragraphid, sentences := createfirstpagedb.FindFreeSentences(locale, themes)

	paragraph := findfreeparagraph.FindFromQ(locale, themes)

	sentences := paragraph.Sentences

	//	paragraph := findfreeparagraph.FindFromQ(locale, themes)

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

	//	go makenewsite.Makenew(locale, themes, host, pathinfo, title, paragraphid)

	firstPage := domains.FirstPage{

		Locale:     locale,
		Themes:     themes,
		Domain:     host,
		Pathinfo:   pathinfo,
		Title:      title,
		Ptitle:     paragraph.Ptitle,
		Pphrase:    paragraph.Pphrase,
		Phost:		paragraph.Phost,
		Plocallink: paragraph.Plocallink,
		Sentences:  paragraph.Sentences,
	}

	go makenewsite.MakeNewByQ(firstPage)

	webinfo := domains.WebInfo{
		Site:       host,
		Ext:        ext,
		Sentences:  sentences,
		Keywords:   destkey,
		UpKeywords: upcasesomekeywords,
		Title:      title,
		Phrases:    destphr[:5],
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
