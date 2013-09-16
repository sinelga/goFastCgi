package createfirstpage

import (
	"log"
	"os"
	"path/filepath"
	"html/template"
//	"encoding/json"
//	"io"
	"bytes"
	d "domains"
)
var index = template.Must(template.ParseFiles(
	"templ/_base.html",
	"templ/index.html",
))


func CreatePage(locale string, themes string, host string, pathinfo string) {

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

//	b := []byte("hellø")
//	lst := make([]map[string]string, 0)
//	filej, err := os.Open("data.json")
//	if err != nil {
//		panic(err)
//	}
//
//	json := json.NewDecoder(filej)
//	json.Decode(&lst)
//	file.Close()
	
		site := string("porno.com")
	ext := string(".html")
	
	newsentslice :=[]string{"llslslslsls","222222222222222222","3333333333333333"}
	somekeywords :=[]string{"porno","seksi","pillu","dildo","fack","sex"}
	upcasesomekeywords :=[]string{"Porno","Seksi","Pillu","Dildo","Fack","Sex"}
	title := string("Seksi chat")
	somephrases :=[]string{"frase   111111","frase 2222","frase 3333","frase 444","frase 55","frase 6666","frase 777","frase 8888"}
	
				webinfo := d.WebInfo{
					Site:       site,
					Ext:        ext,
					Sentences:  newsentslice,
					Keywords:   somekeywords,
					UpKeywords: upcasesomekeywords,
					Title:      title,
					Phrases:    somephrases,
				}
	
//	var tmpl = template.Must(template.New("templateout").Parse("template.html"))
	webpage := bytes.NewBuffer(nil)
	
	if err := index.Execute(webpage, webinfo); err != nil {
					//					http.Error(w, err.Error(), http.StatusInternalServerError)
					panic(err)
				}
		webpagebytes := make([]byte, webpage.Len())
		webpagebytes = webpage.Bytes()
//	io.Copy(w, webpage)		
				
	//	buf := make([]byte, 1024)
	//	 w := bytes.NewBuffer(buf)
	//	 w.Write(b)
	file.Write(webpagebytes)
	file.Close()

	//
	//	b := []byte("hellø")

}
