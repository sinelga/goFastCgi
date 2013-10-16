package findfreeparagraph

import (
	"domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
)

func FindFromQ(locale string, themes string) domains.Paragraph  {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	queuename := locale + ":" + themes
	var unmarPar domains.Paragraph

	if quan_prs, err := redis.Int(c.Do("LLEN", queuename)); err != nil {
		log.Fatal(err)

	} else {

		if quan_prs > 1 {

			bparagraph, _ := redis.Bytes(c.Do("LPOP", queuename))

			
			err := json.Unmarshal(bparagraph, &unmarPar)
			if err != nil {
				log.Fatal(err)

			}

		} else {

			log.Fatalln("need more free paragraphs!!!!")

		}


	}
	return unmarPar
}
