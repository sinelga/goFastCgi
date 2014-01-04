package main

import (
	"flag"
	"log"
	"log/syslog"
	"pushinqueue"
	"strings"
)

var htmlfileflag = flag.String("htmlfile", "", "must be ...html.gz file")
var eventflag = flag.String("event", "", "must be OPEN CLOSE etc.")

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")		
	}

	flag.Parse()
	
	htmlfileflag := *htmlfileflag

	if  strings.HasSuffix(htmlfileflag,".gz")  || strings.HasSuffix(htmlfileflag,"/index.html") {

		if *eventflag == "CLOSE_NOWRITE,CLOSE" {
		
			htmlfilesp := strings.Split(htmlfileflag, "/")
			locale := htmlfilesp[1]
			themes := htmlfilesp[2]
			host := htmlfilesp[3]

			var pathinfo string = "/"

			for i := 4; i < len(htmlfilesp); i++ {

				pathinfo = pathinfo + htmlfilesp[i]

				if i < (len(htmlfilesp) - 1) {
					pathinfo = pathinfo + "/"
				}

			}

			pushinqueue.PushInQueue(*golog, "redis", locale, themes, host, pathinfo)

		}
		
	}

}
