package createfirstgz

import (
	"bytes"
	"checkpathinfo"
	"compress/gzip"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log"
	"log/syslog"
	"strings"
)

func Creategzhtml(byqueue bool, htmlfile string, webpagebytes []byte) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}

	var queuename string
	var c redis.Conn

	if byqueue {
		c, err = redis.Dial("tcp", ":6379")
		if err != nil {
		
			golog.Err(err.Error())
			
		} else {

			queuename = "firstpagebin"
		}
	}

	var index bool = false

	_htmlfile := checkpathinfo.Check(htmlfile)

	if strings.HasSuffix(_htmlfile, "/index.html") {

		index = true
	}

	if !index {

		var b bytes.Buffer

		w := gzip.NewWriter(&b)
		w.Write(webpagebytes)
		w.Close()

		if byqueue {
			if _, err := c.Do("HSET", queuename, _htmlfile, b.Bytes()); err != nil {

				golog.Err(err.Error())
			}
		} else {
			if err := ioutil.WriteFile(_htmlfile, b.Bytes(), 0666); err != nil {
				golog.Crit(err.Error())
			}
		}
	} else {

		if byqueue {
			if _, err := c.Do("HSET", queuename, _htmlfile, webpagebytes); err != nil {

				golog.Err(err.Error())
			}

		} else {
			if err := ioutil.WriteFile(_htmlfile, webpagebytes, 0666); err != nil {
				golog.Crit(err.Error())
			}
		}
	}

	if byqueue {
		c.Flush()
		c.Close()
	}

}
