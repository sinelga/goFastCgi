package pushinqueue

import (
	"domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
	"os"
)

func PushInQueue(golog syslog.Writer, queuesys string, locale string, themes string, host string, pathinfo string) {

	c, err := redis.Dial("tcp", ":6379")
	defer c.Close()
	if err != nil {

		golog.Err("PushInQueue: " + err.Error())
		os.Exit(1)
	} else {

		site := domains.Site{

			Locale:   locale,
			Themes:   themes,
			Domain:   host,
			Pathinfo: pathinfo,
		}
		bsite, _ := json.Marshal(site)

		if _, err := c.Do("SADD", "pagetocreate", bsite); err != nil {
			golog.Err("PushInQueue: " + err.Error())

		}
		c.Flush()
		c.Close()

	}

}
