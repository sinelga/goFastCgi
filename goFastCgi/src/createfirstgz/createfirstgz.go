package createfirstgz

import (

	"compress/gzip"
	"io/ioutil"	
	"bytes"
	"log"
	"log/syslog"
)

func Creategzhtml(golog syslog.Writer,htmlfile string,webpagebytes []byte){

	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write(webpagebytes) 
	w.Close()
	log.Println("!!" +htmlfile+".gz")
//	err := ioutil.WriteFile(htmlfile+".gz", b.Bytes(), 0666)
	if err := ioutil.WriteFile(htmlfile+".gz", b.Bytes(), 0666); err != nil {
		golog.Crit(err.Error())
	}

}