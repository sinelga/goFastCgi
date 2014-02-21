package main

import (
	"cleandb"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"flag"
//	"fmt"
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


var updatedflag = flag.Int("updated", 0, "created in hours ago mast must be > 0 normal 120-144 and more")

func main() {
	flag.Parse() // Scan the arguments list

	if *updatedflag > 0 {

		golog, err := syslog.New(syslog.LOG_ERR, "golog")

		defer golog.Close()
		if err != nil {
			log.Fatal("error writing syslog!!")
		}

		db, err := sql.Open("sqlite3", "file:gofast.db?cache=shared&mode=rwc")
		if err != nil {
			golog.Err(err.Error())
		}
		defer db.Close()

		var sitearr []Site
		secupdated := *updatedflag * 60 * 60

		sqlstr := "select rowid,Locale,Themes,Site,Pathinfo from sites where hits > 0 and Updated < (strftime('%s','now') -" + strconv.Itoa(secupdated) + ")"

		golog.Info(sqlstr)

		rows, err := db.Query(sqlstr)
		if err != nil {

			golog.Err(err.Error())
			os.Exit(-1)
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

				htmlfile = "www/" + site.Locale + "/" + site.Themes + "/" + site.Site + site.Pathinfo
				
//				golog.Info("delete-> "+htmlfile) 

				if finfo, err := os.Stat(htmlfile); err != nil {

					if os.IsNotExist(err) {

						golog.Info("cleanupspace: file does not exist??? Cant be!!! but delete record from DB anyway!! id -> " + strconv.Itoa(site.Id))
						cleandb.Makeclean(db, site.Id)

					}

				} else {

					if !finfo.IsDir() {

						if err := os.Remove(htmlfile); err != nil {

							golog.Err(err.Error())

						} else {

							cleandb.Makeclean(db, site.Id)

						}

					}

				}

			}

		}

	}

}
