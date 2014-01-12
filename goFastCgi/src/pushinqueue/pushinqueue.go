package pushinqueue

import (
	"github.com/garyburd/redigo/redis"
//	"log"
	"log/syslog"
	"domains"
	"encoding/json"
//	"strconv"

)

func PushInQueue(golog syslog.Writer,queuesys string,locale string,themes string,host string,pathinfo string) {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Err(err.Error())
	}

	site := domains.Site{

		Locale:   locale,
		Themes:   themes,
		Domain:   host,
		Pathinfo: pathinfo,
	}
	bsite, _ := json.Marshal(site)

	if _, err := c.Do("SADD", "pagetocreate", bsite); err != nil {
		golog.Err(err.Error())

	}
//	 else {

//		pgqint64 := pgq.(int64)
//		pgqint := int(pgqint64) 
//		golog.Info("pushinqueue: in queue pagetocreate "+ strconv.Itoa(pgqint))
	
//	}
	c.Flush()
	c.Close()
}