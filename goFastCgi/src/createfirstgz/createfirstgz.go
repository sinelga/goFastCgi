package createfirstgz

import (
	//	"bufio"
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
	"log/syslog"
	//	"os"
	"strings"
	"checkpathinfo"
)

func Creategzhtml(htmlfile string, webpagebytes []byte) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}

	var index bool = false

	_htmlfile := checkpathinfo.Check(htmlfile)
	
	golog.Info("Creategzhtml: thispathinfo " + _htmlfile)

	if strings.HasSuffix(_htmlfile, "/index.html") {
	
	  index = true
	} 


//	htmlfilesp := strings.Split(htmlfile, "/")
//
//	lastsplit := htmlfilesp[len(htmlfilesp)-1]
//
//	golog.Info("lastsplit " + lastsplit)
//
//	var _htmlfile string
//	if strings.HasSuffix(htmlfile, ".gz") {
//
//		_htmlfile = htmlfile
//
//	} else if strings.Index(lastsplit, ".") > 0 && lastsplit != "index.html" {
//		_htmlfile = htmlfile + ".gz"
//
//	} else if lastsplit == "index.html" {
//
//		index = true
//		_htmlfile = htmlfile
//
//	} else {
//
//		_htmlfile = htmlfile + "/index.html"
//		golog.Info("CreateIndex: " + _htmlfile)
//
//		index = true
//
//	}

	if !index {
		golog.Info("Creategzhtml: gz file " + _htmlfile)
		
		var b bytes.Buffer

		w := gzip.NewWriter(&b)
		w.Write(webpagebytes)
		w.Close()

		if err := ioutil.WriteFile(_htmlfile, b.Bytes(), 0666); err != nil {
			golog.Crit(err.Error())
		}
	} else {
         
         golog.Info("Creategzhtml: index.html " + _htmlfile)
		
		if err := ioutil.WriteFile(_htmlfile, webpagebytes, 0666); err != nil {
			golog.Crit(err.Error())
		}
	}

}
