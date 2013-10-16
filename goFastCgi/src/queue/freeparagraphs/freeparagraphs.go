package freeparagraphs

import (
	"domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
)

func CreateParagraphs(golog syslog.Writer, locale string, themes string, paragraphs []domains.Paragraph) {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	queuename := locale + ":" + themes

	for _,paragraph   := range paragraphs {

		bparagraph, _ := json.Marshal(paragraph)

		if pgq, err := c.Do("LPUSH", queuename, bparagraph); err != nil {
			log.Fatal(err)

		} else {

			log.Println("in queue ", pgq)

		}

	}

//	if qpages, err := redis.Int(c.Do("LLEN", queuename)); err != nil {
//		log.Fatal(err)
//
//	} else {
//
//		for i := 0; i < qpages; i++ {
//
//			msite, _ := redis.Bytes(c.Do("LPOP", queuename))
//
//			var unmar domains.Paragraph
//			err := json.Unmarshal(msite, &unmar)
//			if err != nil {
//				log.Fatal(err)
//
//			}
//
//			log.Println(unmar.Pphrase)
//
//			//
//			//			htmlfile := string("www/" + unmar.Locale + "/" + unmar.Themes + "/" + unmar.Domain + unmar.Pathinfo)
//			//			log.Println(htmlfile)
//			//
//			//			htmlfileexist.StartCheck(htmlfile, unmar.Domain,unmar.Pathinfo)
//
//		}
//
//	}

}
