package main

import (
	"domains"
	"htmlfileexist"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
)

func main() {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	if qpages, err := redis.Int(c.Do("LLEN", "pagetocreate")); err != nil {
		log.Fatal(err)

	} else {

		for i := 0; i < qpages; i++ {

			log.Println("make")
			msite, _ := redis.Bytes(c.Do("LPOP", "pagetocreate"))

			var unmar domains.Site
			err := json.Unmarshal(msite, &unmar)
			if err != nil {
				log.Fatal(err)

			}
	
			htmlfile := string("www/" + unmar.Locale + "/" + unmar.Themes + "/" + unmar.Domain + unmar.Pathinfo)
			log.Println(htmlfile)
			
			htmlfileexist.StartCheck(htmlfile, unmar.Domain,unmar.Pathinfo)

		}

	}
	c.Flush()
	c.Close()

}