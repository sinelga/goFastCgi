package pushinqueue

import (
	"github.com/garyburd/redigo/redis"
//	"log"
	"log/syslog"
	"domains"
	"encoding/json"
	"strconv"

)

func PushInQueue(golog syslog.Writer,queuesys string,locale string,themes string,host string,pathinfo string) {

//	log.Println("Start PushInQueue")
//	golog.Info("Start PushInQueue: "+host+pathinfo)
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
//		log.Fatal(err)
		golog.Err(err.Error())
	}

	site := domains.Site{

		Locale:   locale,
		Themes:   themes,
		Domain:   host,
		Pathinfo: pathinfo,
	}
	bsite, _ := json.Marshal(site)

	if pgq, err := c.Do("SADD", "pagetocreate", bsite); err != nil {
	
//		log.Fatal(err)
		golog.Err(err.Error())

	} else {
	
//		log.Println("in queue ",pgq)
		pgqint64 := pgq.(int64)
		pgqint := int(pgqint64) 
		golog.Info("pushinqueue: in queue pagetocreate "+ strconv.Itoa(pgqint))
	
	}
	c.Flush()
	c.Close()
}