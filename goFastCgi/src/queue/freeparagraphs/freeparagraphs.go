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



}
