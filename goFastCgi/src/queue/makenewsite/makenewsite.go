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

	if _, err := c.Do("SADD", queuename, bfirstpage); err != nil {
		log.Fatal(err)

	} 

	c.Flush()
	c.Close()
}
