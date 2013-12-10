package checkdbexist

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"domains"
	"log"
	"log/syslog"
	"time"
//	"strconv"
)

func Checkdb(golog syslog.Writer, db *sql.DB, host string, pathinfo string) domains.WebContents {
	var locale string
	var themes string
	var title string
	var site string
	var rowid int64
	var created int64
	var updated int64
	var hits int
	var allhits int

	var webcontents domains.WebContents
	//	var deltamin int
//	golog.Info("Start  Checkdb")
	

	now := time.Now().Unix()

	sqlstr := "select rowid,Created,Updated,Hits,Locale,Themes,Title,Site,Allhits from sites where Site='" + host + "' and Pathinfo='" + pathinfo + "'"

	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&rowid, &created, &updated, &hits, &locale, &themes, &title, &site, &allhits)
		webcontents = domains.WebContents{

			Rowid:    rowid,
			Locale:   locale,
			Themes:   themes,
			Title:    title,
			Site:     site,
			PathInfo: pathinfo,
			Created:  time.Unix(created, 0),
			Updated:  time.Unix(updated, 0),
			Hits:     hits,
			AllHits:  allhits,
		}

	}
	rows.Close()
	
	
//	golog.Info("rowid " + strconv.FormatInt(rowid,2) +" "+host+" "+pathinfo)
//	golog.Info("webcontents.rowid " + strconv.FormatInt(webcontents.Rowid,2)) 
	
	if rowid > 0 {

		tx, err := db.Begin()
		if err != nil {

			golog.Crit(err.Error())
		}

		newhits := hits + 1

		sqlstr = "update sites set Hits = ?,Updated=? where rowid = ?"

		stmt, err := tx.Prepare(sqlstr)
		if err != nil {

			golog.Crit(err.Error())
		}
		defer stmt.Close()
		if _, err = stmt.Exec(newhits, now, rowid); err != nil {

			golog.Crit(err.Error())
		}

		tx.Commit()

	} else {

		golog.Warning("Record for " + host + "/" + pathinfo + " dont exist in sites!!!")

	}

	return webcontents

}
