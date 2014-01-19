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
var deepflag = flag.Int("deep", 0, "param to make deep clean if 0 nothing will be done")

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
			golog.Err(err.Error())
		}
		defer db.Close()

		seccreated := *createdflag * 24 * 60 * 60
		
		//		sqlstr := "select rowid,Locale,Themes,Site,Pathinfo from sites where hits >=" + strconv.Itoa(*hitsflag) + " or Created < (strftime('%s','now') -" + strconv.Itoa(seccreated) + ")"
		sqlstr := "select rowid,Locale,Themes,Site,Pathinfo from sites where hits <" + strconv.Itoa(*hitsflag) + " and Created < (strftime('%s','now') -" + strconv.Itoa(seccreated) + ")"
		//		log.Println(sqlstr)
		golog.Info("cleanupspace: "+sqlstr)


		rows, err := db.Query(sqlstr)
		if err != nil {

			golog.Err(err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var site Site

			rows.Scan(&site.Id, &site.Locale, &site.Themes, &site.Site, &site.Pathinfo)

			sitearr = append(sitearr, site)

		}
		rows.Close()


		if *deepflag > 0 { 

		golog.Info("cleanupspace: deep clean ") //may be add deep to request
		sqlstr = "SELECT rowid,Locale,Themes,Site,Pathinfo,count(*) FROM sites group by Site,Pathinfo"

		rows, err = db.Query(sqlstr)
		
		if err != nil {
			golog.Err(err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var site Site
			var count int

			rows.Scan(&site.Id, &site.Locale, &site.Themes, &site.Site, &site.Pathinfo,&count)
			
			if count > 1 {
				golog.Info("cleanupspace:!!!Bad delete all -->"+site.Site+site.Pathinfo + " count "+strconv.Itoa(count))
				sitearr = append(sitearr, site)
			
			}
		}
		rows.Close()
		
		} else {
		
			golog.Info("cleanupspace:NOT deep clean ")
		
		}

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
	golog.Info("cleanupspace: END cleanupspace ")
	
}
