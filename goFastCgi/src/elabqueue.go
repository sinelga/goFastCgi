package main

import (
	"domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"htmlfileexist"
	"io/ioutil"
	"log"
	"log/syslog"
	"makenewsite"
)

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())
	}


	golog.Info("elabque: Start firstpagebin")
	if qfirstpagesbin, err := redis.Strings(c.Do("HKEYS", "firstpagebin")); err != nil {
		golog.Crit(err.Error())
	} else {

		for _, keystr := range qfirstpagesbin {

			if webpagebytes, err := redis.Bytes(c.Do("HGET", "firstpagebin", keystr)); err != nil {

				golog.Crit(err.Error())

			} else {

				if err := ioutil.WriteFile(keystr, webpagebytes, 0666); err != nil {
					golog.Crit(err.Error())
				} else {

					if _, err := redis.Int(c.Do("HDEL", "firstpagebin", keystr)); err != nil {

						golog.Crit(err.Error())

					}

				}

			}

		}

	}

	golog.Info("elabque: Start firstpage")
	if qfirstpages, err := redis.Int(c.Do("SCARD", "firstpage")); err != nil {
		golog.Crit(err.Error())

	} else {

		for i := 0; i < qfirstpages; i++ {

			bfirstpage, _ := redis.Bytes(c.Do("SPOP", "firstpage"))
			var unmar domains.FirstPage
			err := json.Unmarshal(bfirstpage, &unmar)
			if err != nil {
				//				log.Fatal(err)
				golog.Crit(err.Error())

			}
			makenewsite.Makenew(*golog, unmar)

		}

	}

	golog.Info("elabque: Start pagetocreate")
	if qpages, err := redis.Int(c.Do("SCARD", "pagetocreate")); err != nil {
		log.Fatal(err)

	} else {

		for i := 0; i < qpages; i++ {

			msite, _ := redis.Bytes(c.Do("SPOP", "pagetocreate"))

			var unmar domains.Site
			err := json.Unmarshal(msite, &unmar)
			if err != nil {
				golog.Crit(err.Error())

			}

			htmlfile := string("www/" + unmar.Locale + "/" + unmar.Themes + "/" + unmar.Domain + unmar.Pathinfo)

			htmlfileexist.StartCheck(*golog, htmlfile, unmar.Domain, unmar.Pathinfo)

		}

	}

	golog.Info("elabque: END pagetocreate")
	c.Flush()
	c.Close()

}
