package createfirstpage

import (
	"bytes"
	"comutils"
	"domains"
	"html/template"
//	"log"
	"log/syslog"
	"math/rand"
	"os"
	"path/filepath"
	"queue/findfreeparagraph"
	"queue/makenewsite"
	"checkpathinfo"
//	"strings"
	"time"
	"titlegen"
//	"compress/gzip"
//	"io/ioutil"
)

func CreatePage(golog syslog.Writer, locale string, themes string, host string, pathinfo string, keywords []string, phrases []string) []byte{

	var index = template.Must(template.ParseFiles(
		"templ/_base_first.html",
		"templ/index_first.html",
	))

	thispathinfo := checkpathinfo.Check(pathinfo)

	htmlfile := string("www/" + locale + "/" + themes + "/" + host + thispathinfo)

	path := filepath.Dir(htmlfile)
//	log.Println(path)
	err := os.MkdirAll(path, 0777)
	if err != nil {

		golog.Err(err.Error())
	}

//	file, err := os.Create(htmlfile)
//	if err != nil {
//		//		panic(err)
//		log.Fatal(err)
//	}

	ext := string(".html")

	paragraph, newdomain := findfreeparagraph.FindFromQ(locale, themes)

	if newdomain != "" {
		//		log.Println("new domain", newdomain)
		golog.Info("new domain " + newdomain)
	}

	sentences := paragraph.Sentences

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

	title := comutils.UpcaseInitial(titlegen.GetTitle(locale, themes, host, thispathinfo))

	destphr := make([]string, len(phrases))
	rand.Seed(time.Now().UTC().UnixNano())
	permphr := rand.Perm(len(phrases))

	for i, v := range permphr {

		destphr[v] = comutils.UpcaseInitial(phrases[i])
	}

	firstPage := domains.FirstPage{

		Locale:     locale,
		Themes:     themes,
		Domain:     host,
//		Pathinfo:   thispathinfo+".gz",
		Pathinfo:   thispathinfo,
		Title:      title,
		Ptitle:     paragraph.Ptitle,
		Pphrase:    paragraph.Pphrase,
		Phost:      paragraph.Phost,
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
		Newdomain:  newdomain,
	}

	webpage := bytes.NewBuffer(nil)

	if err := index.Execute(webpage, webinfo); err != nil {

//		panic(err)
		golog.Err(err.Error())
	}
	webpagebytes := make([]byte, webpage.Len())
	webpagebytes = webpage.Bytes()
	

return webpagebytes

	
}
