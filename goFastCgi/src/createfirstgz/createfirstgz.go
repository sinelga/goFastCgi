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
)

func Creategzhtml(htmlfile string, webpagebytes []byte) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}

	var index bool = false

	var b bytes.Buffer

	golog.Info("Start Creategzhtml: " + htmlfile)

	htmlfilesp := strings.Split(htmlfile, "/")

	lastsplit := htmlfilesp[len(htmlfilesp)-1]

	golog.Info("lastsplit " + lastsplit)

	var _htmlfile string
	if strings.HasSuffix(htmlfile, ".gz") {

		_htmlfile = htmlfile

	} else if strings.Index(lastsplit, ".") > 0 {
		_htmlfile = htmlfile + ".gz"

	} else {

		_htmlfile = htmlfile + "/index.html"
		golog.Info("CreateIndex: " + _htmlfile)

		index = true
		//		file, err := os.Create(_htmlfile)
		//		if err != nil {
		//			golog.Crit(err.Error())
		//		}

		//		file.Write(webpagebytes)
		//		file.Close()

//		writer := bufio.NewWriter(&b)
//		writer.Write(webpagebytes)
//		writer.Flush()
		//		writer.Flush()
		//		file.Close()
		if err := ioutil.WriteFile(_htmlfile, webpagebytes, 0666); err != nil {
			golog.Crit(err.Error())
		}

	}

	if !index {
		golog.Info("Creategzhtml: " + _htmlfile)

		w := gzip.NewWriter(&b)
		w.Write(webpagebytes)
		w.Close()

		if err := ioutil.WriteFile(_htmlfile, b.Bytes(), 0666); err != nil {
			golog.Crit(err.Error())
		}
	}

}
