package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"flag"
	"github.com/garyburd/redigo/redis"
	"log"
)

var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")
var domainFlag = flag.String("domain", "", "any like test.com")
var expireFlag = flag.Int("expire", 0, "expire must be > 0 in seconds")

var hostsarr []string
var newsite string

func main() {
	flag.Parse()
	queuename := "newdomains:" + *localeFlag + ":" + *themesFlag
	log.Println(queuename)

	db, err := sql.Open("sqlite3", "singo.db")
	if err != nil {
		log.Fatal(err)
	}
	sqlstr := "select host from hosts where locale='" + *localeFlag + "' and themes='" + *themesFlag + "'"

	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var host string
		rows.Scan(&host)
		hostsarr = append(hostsarr, host)
	}
	rows.Close()

	db.Close()

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	for _, site := range hostsarr {

		newsite = site + "." + *domainFlag
		log.Println("hosts", newsite)
		if _, err := c.Do("SADD",queuename, newsite); err != nil {

			log.Fatal(err)

		}
	}

	if expset, err := c.Do("EXPIRE",queuename, *expireFlag); err != nil {

	} else {

		log.Println("EXPIRE ", expset)
	}

}
