package main

import (
	"domains"
	"htmlfileexist"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
	"makenewsite"
)

func main() {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	
	if qfirstpages, err := redis.Int(c.Do("SCARD", "firstpage")); err != nil {
		log.Fatal(err)

	} else {
	
		for i := 0; i < qfirstpages; i++ {
			
			bfirstpage, _ := redis.Bytes(c.Do("SPOP", "firstpage"))
			var unmar domains.FirstPage
			err := json.Unmarshal(bfirstpage, &unmar)
			if err != nil {
				log.Fatal(err)

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
