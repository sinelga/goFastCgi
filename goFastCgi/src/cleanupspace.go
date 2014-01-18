package main

import (
	"cleandb"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"flag"
	"log"
	"log/syslog"
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

var hitsflag = flag.Int("hits", 0, "hits must be >= 0 normal 10 and more")
var createdflag = flag.Int("created", 0, "created in days ago mast must be > 0 normal 20-30 and more")

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	flag.Parse()

	if *hitsflag >= 0 && *createdflag > 0 {

		var sitearr []Site

		db, err := sql.Open("sqlite3", "file:gofast.db?cache=shared&mode=rwc")
		if err != nil {
//			log.Fatal(err)
			golog.Err(err.Error())
		}
		defer db.Close()

		seccreated := *createdflag * 24 * 60 * 60
		//		sqlstr := "select rowid,Locale,Themes,Site,Pathinfo from sites where hits >=" + strconv.Itoa(*hitsflag) + " or Created < (strftime('%s','now') -" + strconv.Itoa(seccreated) + ")"
		sqlstr := "select rowid,Locale,Themes,Site,Pathinfo from sites where hits <" + strconv.Itoa(*hitsflag) + " and Created < (strftime('%s','now') -" + strconv.Itoa(seccreated) + ")"
		//		log.Println(sqlstr)
		golog.Info(sqlstr)

		rows, err := db.Query(sqlstr)
		if err != nil {
//			log.Fatal(err)
			golog.Err(err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var site Site

			rows.Scan(&site.Id, &site.Locale, &site.Themes, &site.Site, &site.Pathinfo)

			sitearr = append(sitearr, site)

		}
		rows.Close()

		sqlstr = "SELECT  rowid,Locale,Themes,Site,Pathinfo FROM sites group by Site,Pathinfo"

		rows, err = db.Query(sqlstr)
		
		if err != nil {
//			log.Fatal(err)
			golog.Err(err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var site Site

			rows.Scan(&site.Id, &site.Locale, &site.Themes, &site.Site, &site.Pathinfo)
			
			golog.Info("Bad -->"+site.Site+site.Pathinfo)
//			sitearr = append(sitearr, site)

		}
		rows.Close()

		if len(sitearr) > 0 {

			var htmlfile string

			for _, site := range sitearr {

				htmlfile = "www/" + site.Locale + "/" + site.Themes + "/" + site.Site + site.Pathinfo

				log.Println(htmlfile)

				if finfo, err := os.Stat(htmlfile); err != nil {

					if os.IsNotExist(err) {
						//						log.Fatall("file does not exist??? Cant be!!! but delete record from DB anyway")
						log.Println("file does not exist??? Cant be!!! but delete record from DB anyway!! id ->", site.Id)
						cleandb.Makeclean(db, site.Id)

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
