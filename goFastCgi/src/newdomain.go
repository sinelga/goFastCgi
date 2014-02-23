package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"flag"
	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
	"pushdomain"
)

var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")
var domainFlag = flag.String("domain", "", "any like test.com")
var expireFlag = flag.Int("expire", 0, "expire must be > 0 in seconds")

var hostsarr []string
var newsite string

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	flag.Parse()

	locale := *localeFlag
	themes := *themesFlag
	domain := *domainFlag
	expire := *expireFlag

//	queuename := "newdomains:" + locale + ":" + themes
	queuename := "newdomains"
	
	if domain == "" {
	
		domain = pushdomain.SelectDomain(*golog,locale,themes)
	
	} 
		
	db, err := sql.Open("sqlite3", "singo.db")
	defer db.Close()
	if err != nil {

		golog.Err("newdomain: " + err.Error())
	} else {
		sqlstr := "select host from hosts where locale='" + locale + "' and themes='" + themes + "'"

		rows, err := db.Query(sqlstr)
		if err != nil {

			golog.Err("newdomain: " + err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var host string
			rows.Scan(&host)
			hostsarr = append(hostsarr, host)
		}
		rows.Close()
		db.Close()
	}

	c, err := redis.Dial("tcp", ":6379")
	defer c.Close()
	if err != nil {

		golog.Err("newdomain: " + err.Error())
		
	} else {

		for _, site := range hostsarr {

			newsite = site + "." + domain		
			golog.Info("newdomain: " +newsite)
			if _, err := c.Do("SADD", queuename, newsite); err != nil {
			
				golog.Err("newdomain: " + err.Error())

			}
		}

		if _, err := c.Do("EXPIRE", queuename, expire); err != nil {

			golog.Err("newdomain: " + err.Error())

		}

	}
	c.Flush()
	c.Close()

}
