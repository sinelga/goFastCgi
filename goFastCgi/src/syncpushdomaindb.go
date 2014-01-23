package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"domains"
	"flag"
	"fmt"
	"log"
	"log/syslog"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)

	}

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	var extdomainsarr []domains.Domaincsv
	var extdomainsarrtoinsert []domains.Domaincsv
	var domainstopush []string

	db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/singo.db")
	if err != nil {

		golog.Err(err.Error())
	}
	sqlstr := "select Locale,Themes,Domain,Ip from extdomains"

	golog.Info("syncpushdomain: Start " + sqlstr)

	rows, err := db.Query(sqlstr)
	if err != nil {

		golog.Err(err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		var extdomain domains.Domaincsv
		rows.Scan(&extdomain.Locale, &extdomain.Themes, &extdomain.Domain, &extdomain.Ip)
		extdomainsarr = append(extdomainsarr, extdomain)

	}
	rows.Close()

	db.Close()

	db, err = sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/pushdomains.db")
	if err != nil {

		golog.Err(err.Error())
	}
	sqlstr = "select Domain from pushdomains"
	rows, err = db.Query(sqlstr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var domaintopush string
		rows.Scan(&domaintopush)
		domainstopush = append(domainstopush, domaintopush)
	}

	rows.Close()

	var newpushdomain bool

	for _, extdomain := range extdomainsarr {
	
		newpushdomain = false
		
		for _, pushdomain := range domainstopush {
			
			if extdomain.Domain == pushdomain {

				newpushdomain = true
			}

		}

		if !newpushdomain {

			golog.Info("Add " + extdomain.Domain)
			extdomainsarrtoinsert = append(extdomainsarrtoinsert, extdomain)
			
		}

	}

	if len(extdomainsarrtoinsert) > 0 {

		sqlstr = "insert into pushdomains(Locale,Themes,Domain,Ip) values(?,?,?,?)"

		tx, err := db.Begin()
		if err != nil {
			golog.Err(err.Error())
		}
		stmt, err := tx.Prepare(sqlstr)
		if err != nil {
			log.Fatal(err)
			golog.Err(err.Error())
		}
		defer stmt.Close()

		for _, domain := range extdomainsarrtoinsert {

			if _, err = stmt.Exec(domain.Locale, domain.Themes, domain.Domain, domain.Ip); err != nil {
				
				golog.Err(err.Error())

			}

		}
		stmt.Close()
		tx.Commit()

	}

	db.Close()

}
