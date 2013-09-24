package createpage

import (
	"bytes"
	"domains"
	"html/template"
	"log"
	"os"
)

func CreatePg(htmlfile string, webcontents domains.WebContents) {

//	pwd := os.Get("templ/_base.html")	
//	log.Println("pwd",pwd)

	var index = template.Must(template.ParseFiles(
		"templ/_base.html",
		"templ/index.html",
	))

	log.Println("createpage:")
	webpage := bytes.NewBuffer(nil)
	if err := index.Execute(webpage, webcontents); err != nil {

		panic(err)
	}
	webpagebytes := make([]byte, webpage.Len())
	webpagebytes = webpage.Bytes()

	file, err := os.Create(htmlfile)
	if err != nil {
		panic(err)
	}
	file.Write(webpagebytes)
	file.Close()

}
