package main

import (
	"cleandb"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"log/syslog"
	"os"
	"strconv"
	"time"
)

type Site struct {
	Id       int
	Locale   string
	Themes   string
	Site     string
	Pathinfo string
}

var hitsflag = flag.Int("hits", 0, "hits must be >= 0 normal 10 and more")
var createdflag = flag.Int("created", 0, "created in hours ago mast must be > 0 normal 120-144 and more")
var deepflag = flag.Int("deep", 0, "param to make deep clean if 0 nothing will be done")
var updatedflagstr string
var oldfile bool

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	flag.Parse()

	hitsflagint := *hitsflag
	createdflagint := *createdflag
	deepflagint := *deepflag

	golog.Info("Start cleanupspace")

	oldfile = false

	if hitsflagint >= 0 && createdflagint > 0 {

		cleaner(*golog, hitsflagint, createdflagint, deepflagint)

	} else {

		if _, err := os.Stat("cleanupspace.csv"); err != nil {
			if os.IsNotExist(err) {

				golog.Info("cleanupspace.csv dont exist create default")
				csvFile, err := os.Create("cleanupspace.csv")
				defer csvFile.Close()
				if err != nil {

					golog.Crit("cleanupspace: " + err.Error())
				}
				writer := csv.NewWriter(csvFile)

				csvstring := []string{"1", "720", "0", "800"}

				err = writer.Write(csvstring)
				if err != nil {

					golog.Crit("cleanupspace: " + err.Error())
				}
				writer.Flush()
				csvFile.Close()

			}

		}

		finfo, err := os.Stat("cleanupspace.csv")
		if err != nil {
			// TODO: handle errors (e.g. file not found)
			golog.Crit("cleanupspace: " + err.Error())
		}

		lasmod := finfo.ModTime().Unix()

		if (time.Now().Unix() - lasmod) > 84400 {

			golog.Info("old don't change createdflagint in cleanupspace.csv")
			oldfile = true
		} else {

			golog.Info("resently updated (less 1 day)  cleanupspace.csv modify !!")
		}

		csvFile, err := os.Open("cleanupspace.csv")
		defer csvFile.Close()
		if err != nil {

			golog.Crit("cleanupspace: " + err.Error())
		}
		csvReader := csv.NewReader(csvFile)

		for {
			fields, err := csvReader.Read()

			if err == io.EOF {
				break
			} else if err != nil {
				golog.Crit(err.Error())

			}

			hitsflagint, err = strconv.Atoi(fields[0])
			if err != nil {

				golog.Crit("cleanupspace: " + err.Error())
			}
			createdflagint, err = strconv.Atoi(fields[1])
			if err != nil {

				golog.Crit("cleanupspace: " + err.Error())
			}
			deepflagint, err = strconv.Atoi(fields[2])
			if err != nil {

				golog.Crit("cleanupspace: " + err.Error())
			}
			updatedflagstr = fields[3]

		}
		csvFile.Close()

		golog.Info("hitsflagint " + strconv.Itoa(hitsflagint) + " createdflagint " + strconv.Itoa(createdflagint))

		csvFile, err = os.Create("cleanupspace.csv")
		defer csvFile.Close()
		if err != nil {

			golog.Crit("cleanupspace: " + err.Error())
		}
		var newcreatedflagint int

		if !oldfile {
			newcreatedflagint = createdflagint - 5

		} else {

			newcreatedflagint = createdflagint
		}

		writer := csv.NewWriter(csvFile)

		csvstring := []string{strconv.Itoa(hitsflagint), strconv.Itoa(newcreatedflagint), strconv.Itoa(deepflagint), updatedflagstr}

		err = writer.Write(csvstring)
		if err != nil {

			golog.Crit("cleanupspace: " + err.Error())
		}
		writer.Flush()
		csvFile.Close()

		cleaner(*golog, hitsflagint, createdflagint, deepflagint)

	}
	golog.Info("Stop cleanupspace")

}

func cleaner(golog syslog.Writer, hitsflagint int, createdflagint int, deepflagint int) {

	var sitearr []Site

	db, err := sql.Open("sqlite3", "file:gofast.db?cache=shared&mode=rwc")
	if err != nil {
		golog.Err(err.Error())
	}
	defer db.Close()

	seccreated := createdflagint * 60 * 60

	//		sqlstr := "select rowid,Locale,Themes,Site,Pathinfo from sites where hits >=" + strconv.Itoa(*hitsflag) + " or Created < (strftime('%s','now') -" + strconv.Itoa(seccreated) + ")"
	sqlstr := "select rowid,Locale,Themes,Site,Pathinfo from sites where hits <" + strconv.Itoa(hitsflagint) + " and Created < (strftime('%s','now') -" + strconv.Itoa(seccreated) + ")"

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

	if deepflagint > 0 {

		golog.Info("cleanupspace: deep clean ") //may be add deep to request
		sqlstr = "SELECT rowid,Locale,Themes,Site,Pathinfo,count(*) FROM sites group by Site,Pathinfo"

		rows, err = db.Query(sqlstr)
		if err != nil {
			golog.Err("cleanupspace: " + err.Error())
			os.Exit(-1)
		}
		defer rows.Close()

		for rows.Next() {
			var site Site
			var count int
			
			rows.Scan(&site.Id, &site.Locale, &site.Themes, &site.Site, &site.Pathinfo, &count)

			if count > 1 {
				golog.Info("!!!cleanupspace: Bad ?? delete all -->" + site.Site + site.Pathinfo + " count " + strconv.Itoa(count))
				sitearr = append(sitearr, site)

			}
		}
		rows.Close()

	} else {

		golog.Info("cleanupspace: NOT --> deep clean ")

	}

	if len(sitearr) > 0 {

		var htmlfile string

		for _, site := range sitearr {

			htmlfile = "www/" + site.Locale + "/" + site.Themes + "/" + site.Site + site.Pathinfo

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

//	golog.Info("cleanupspace: END cleanupspace ")
