package main

import (
	"domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"htmlfileexist"
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
		//		log.Fatal(err)
		golog.Crit(err.Error())
	}

	if qfirstpages, err := redis.Int(c.Do("SCARD", "firstpage")); err != nil {
		//		log.Fatal(err)
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
			makenewsite.Makenew(unmar)

		}

	}

	if qpages, err := redis.Int(c.Do("SCARD", "pagetocreate")); err != nil {
		log.Fatal(err)

	} else {

		for i := 0; i < qpages; i++ {

			msite, _ := redis.Bytes(c.Do("SPOP", "pagetocreate"))

			var unmar domains.Site
			err := json.Unmarshal(msite, &unmar)
			if err != nil {
//				log.Fatal(err)
				golog.Crit(err.Error())

			}

			htmlfile := string("www/" + unmar.Locale + "/" + unmar.Themes + "/" + unmar.Domain + unmar.Pathinfo)
			log.Println(htmlfile)
			golog.Info("elabqueue: check existense "+htmlfile) 

			htmlfileexist.StartCheck(*golog, htmlfile, unmar.Domain, unmar.Pathinfo)

		}

	}

	c.Flush()
	c.Close()

}
