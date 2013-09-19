package createfirstpage

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	//	"encoding/json"
	//	"io"
	"bytes"
	"comutils"
	d "domains"
	"math/rand"
	"time"
	"titlegen"
)

var index = template.Must(template.ParseFiles(
	"templ/_base.html",
	"templ/index.html",
))

func CreatePage(locale string, themes string, host string, pathinfo string, keywords []string, phrases []string) {

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

//	site := string("porno.com")
	ext := string(".html")

	newsentslice := []string{"llslslslsls", "222222222222222222", "3333333333333333"}
	destkey := make([]string, len(keywords))

	rand.Seed(time.Now().UTC().UnixNano())
	permkey := rand.Perm(len(keywords))

	for i, v := range permkey {
		destkey[v] = keywords[i]
	}

	//	somekeywords := []string{"porno", "seksi", "pillu", "dildo", "fack", "sex"}
	upcasesomekeywords := make([]string, len(keywords))
	for i := 0; i < len(keywords); i++ {
		
		upcasesomekeywords[i] = comutils.UpcaseInitial(destkey[i])
	}

	title := comutils.UpcaseInitial(titlegen.GetTitle(locale, themes, host, pathinfo))
//	somephrases := []string{"frase   111111", "frase 2222", "frase 3333", "frase 444", "frase 55", "frase 6666", "frase 777", "frase 8888"}
	
	destphr := make([]string, len(phrases))
	rand.Seed(time.Now().UTC().UnixNano())
	permphr := rand.Perm(len(phrases))
	
	for i, v := range permphr {
		destphr[v] = comutils.UpcaseInitial(phrases[i])
	}

	webinfo := d.WebInfo{
		Site:       host,
		Ext:        ext,
		Sentences:  newsentslice,
		Keywords:   destkey,
		UpKeywords: upcasesomekeywords,
		Title:      title,
		Phrases:    destphr,
	}

	webpage := bytes.NewBuffer(nil)

	if err := index.Execute(webpage, webinfo); err != nil {
		//					http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
	webpagebytes := make([]byte, webpage.Len())
	webpagebytes = webpage.Bytes()

	file.Write(webpagebytes)
	file.Close()

}
