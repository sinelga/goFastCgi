package createpage

import (
	"bytes"
	"domains"
	"html/template"
//	"log"
	"log/syslog"
//	"os"
	"createfirstgz"
)

func CreatePg(golog syslog.Writer, htmlfile string, webcontents domains.WebContents) {

	var index = template.Must(template.ParseFiles(
		"templ/_base.html",
		"templ/index.html",
	))

	thishtmlfile := htmlfile
// 	log.Println("createpage: "+thishtmlfile)
// 	golog.Info("CreatePg: "+thishtmlfile)
	
	webpage := bytes.NewBuffer(nil)
	if err := index.Execute(webpage, webcontents); err != nil {

		panic(err)
	}
	webpagebytes := make([]byte, webpage.Len())
	webpagebytes = webpage.Bytes()


	createfirstgz.Creategzhtml(false,thishtmlfile,webpagebytes)


}
