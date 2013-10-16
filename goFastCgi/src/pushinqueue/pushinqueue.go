package pushinqueue

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"domains"
	"encoding/json"

)

func PushInQueue(queuesys string,locale string,themes string,host string,pathinfo string) {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	site := domains.Site{

		Locale:   locale,
		Themes:   themes,
		Domain:   host,
		Pathinfo: pathinfo,
	}
	bsite, _ := json.Marshal(site)
	
	if pgq, err := c.Do("LPUSH", "pagetocreate", bsite); err != nil {
		log.Fatal(err)

	} else {
	
		log.Println("in queue ",pgq)
	
	}
	c.Flush()
	c.Close()
}