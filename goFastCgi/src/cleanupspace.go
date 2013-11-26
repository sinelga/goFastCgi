package main

import (
	"cleandb"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"flag"
	"log"
	"os"
	"strconv"
)

type Site struct {
	Id       int
	Locale   string
	Themes   string
	Site     string
	Pathinfo string
}

var hitsflag = flag.Int("hits", 0, "hits must be > 0 normal 10 and more")

func main() {

	flag.Parse()

	if *hitsflag > 0 {
		
		var sitearr []Site

		db, err := sql.Open("sqlite3", "file:gofast.db?cache=shared&mode=rwc")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		sqlstr := "select rowid,Locale,Themes,Site,Pathinfo from sites where hits >"+ strconv.Itoa(*hitsflag)

		rows, err := db.Query(sqlstr)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var site Site

			rows.Scan(&site.Id, &site.Locale, &site.Themes, &site.Site, &site.Pathinfo)

			sitearr = append(sitearr, site)

		}
		rows.Close()

		if len(sitearr) > 0 {

			var htmlfile string

			for _, site := range sitearr {

				//	  	log.Println(site.Site,site.Pathinfo)
				htmlfile = "www/" + site.Locale + "/" + site.Themes + "/" + site.Site + site.Pathinfo

				log.Println(htmlfile)

				if finfo, err := os.Stat(htmlfile); err != nil {

					if os.IsNotExist(err) {
						log.Fatalln("file does not exist??? Cant be!!!")

					}

				} else {

					if !finfo.IsDir() {

						log.Println("Exist not dir")
						if err := os.Remove(htmlfile); err != nil {

							log.Println(err.Error())

						} else {

							log.Println("file deleted ", htmlfile)

							cleandb.Makeclean(db, site.Id)

						}

					}

				}

			}

		}

	} else {

		log.Fatalln("check params try -h")

	}
}
