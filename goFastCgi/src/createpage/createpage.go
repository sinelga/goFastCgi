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
 	golog.Info("CreatePg: "+thishtmlfile)
	
	webpage := bytes.NewBuffer(nil)
	if err := index.Execute(webpage, webcontents); err != nil {

		panic(err)
	}
	webpagebytes := make([]byte, webpage.Len())
	webpagebytes = webpage.Bytes()

//	if f, err := os.Open(htmlfile); err != nil {
//
//		golog.Err(err.Error())
//		//			return
//	} else {
//
//		defer f.Close()
//
//		fi, err := f.Stat()
//		if err != nil {
//		
//			golog.Err(err.Error())
//			
//		}
//		switch mode := fi.Mode(); {
//
//		case mode.IsDir():
//
//			golog.Warning("directory ??? fix it add index.html so result " + htmlfile+"/index.html")
//			thishtmlfile = htmlfile+"/index.html"
//			
//		case mode.IsRegular():
//
//		}
//
//	}

	createfirstgz.Creategzhtml(thishtmlfile,webpagebytes)


//	file, err := os.Create(thishtmlfile)
//	if err != nil {
//		panic(err)
//	}
//	file.Write(webpagebytes)
//	file.Close()

}
