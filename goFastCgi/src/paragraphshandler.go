package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"createparagraphs"
	"database/sql"
	"flag"
	"log"
	"log/syslog"
	"os"
	"strconv"
)

var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance...")
var quantFlag = flag.Int("quant", 0, "quant must be > 0")

var keywordsarr []string
var phrasesarr []string
var hostsarr []string

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}
	flag.Parse()

	var locale string
	var themes string
	var quant int

	if *quantFlag > 0 {

		quant = *quantFlag

		if *localeFlag == "fi_FI" && *themesFlag == "finance" {

			locale = "fi_FI"
			themes = "finance"

		} else if *localeFlag == "fi_FI" && *themesFlag == "porno" {

			locale = "fi_FI"
			themes = "porno"

		} else {

			log.Fatalln("exit check parametrs by paragraphshandler -h")
			os.Exit(0)

		}

		dbst, err := sql.Open("sqlite3", "gofast.db")
		if err != nil {
			log.Fatal(err)
		}

		sqlstr := "select count() as count from paragraphs where locale='" + locale + "' and themes='" + themes + "' and Siteid is null"

		rows, err := dbst.Query(sqlstr)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var count string
		for rows.Next() {

			rows.Scan(&count)
		}
		rows.Close()
		log.Println("count", count)
		dbst.Close()

		golog.Info("Free parargraphs-> " + locale + " " + themes + " " + count)

		countint, err := strconv.Atoi(count)
		if err != nil {
			os.Exit(2)
		}

		if countint < quant {
		
			golog.Info("Time make job add "+strconv.Itoa(quant)+" "+locale + " " + themes)

			db, err := sql.Open("sqlite3", "singo.db")
			if err != nil {
				log.Fatal(err)
			}

			sqlstr = "select keyword from keywords where locale='" + locale + "' and themes='" + themes + "'"

			rows, err = db.Query(sqlstr)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			for rows.Next() {
				var keyword string
				rows.Scan(&keyword)
				keywordsarr = append(keywordsarr, keyword)

			}
			rows.Close()
			log.Println("keywords", len(keywordsarr))

			sqlstr = "select phrase from phrases where locale='" + locale + "' and themes='" + themes + "'"

			rows, err = db.Query(sqlstr)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			for rows.Next() {
				var phrase string
				rows.Scan(&phrase)
				phrasesarr = append(phrasesarr, phrase)

			}
			rows.Close()
			log.Println("Phrases", len(phrasesarr))

			sqlstr = "select host from hosts where locale='" + locale + "' and themes='" + themes + "'"

			rows, err = db.Query(sqlstr)
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
			log.Println("hosts", len(hostsarr))

			if db.Close(); err != nil {

				log.Fatal(err)

			} else {

				createparagraphs.CreatePr(locale, themes, keywordsarr, phrasesarr, hostsarr, quant)

			}
		} else {
			
			golog.Info("Not job for "+locale + " " + themes)
			
		
		} //make job
	} //end quant
}
