package makenewsite

import (
	"domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
)

func MakeNewByQ(firstpage  domains.FirstPage) {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	queuename := "firstpage"

	bfirstpage, _ := json.Marshal(firstpage)

	if pgq, err := c.Do("LPUSH", queuename, bfirstpage); err != nil {
		log.Fatal(err)

	} else {

		log.Println("in queue ",queuename, pgq)

	}

	c.Flush()
	c.Close()
}
