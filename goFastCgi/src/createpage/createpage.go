package createpage

import (
	"bytes"
	"domains"
	"html/template"
	"log"
	"log/syslog"
	"os"
)

func CreatePg(golog syslog.Writer, htmlfile string, webcontents domains.WebContents) {

	var index = template.Must(template.ParseFiles(
		"templ/_base.html",
		"templ/index.html",
	))

	log.Println("createpage:")
	thishtmlfile := htmlfile 
	
	webpage := bytes.NewBuffer(nil)
	if err := index.Execute(webpage, webcontents); err != nil {

		panic(err)
	}
	webpagebytes := make([]byte, webpage.Len())
	webpagebytes = webpage.Bytes()

	if f, err := os.Open(htmlfile); err != nil {

		golog.Err(err.Error())
		//			return
	} else {

		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			//			fmt.Println(err)
			golog.Err(err.Error())
			//			return
		}
		switch mode := fi.Mode(); {

		case mode.IsDir():

			golog.Warning("directory ??? fix it add index.html so result " + htmlfile+"/index.html")
			thishtmlfile = htmlfile+"/index.html"
			
		case mode.IsRegular():

			golog.Info("IsRegular so remake " + htmlfile)
			//				os.Remove(htmlfile)
		}

	}

	file, err := os.Create(thishtmlfile)
	if err != nil {
		panic(err)
	}
	file.Write(webpagebytes)
	file.Close()

}
