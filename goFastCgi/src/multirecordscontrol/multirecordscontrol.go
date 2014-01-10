package multirecordscontrol

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log/syslog"
)

func CheckMulti(golog syslog.Writer, db *sql.DB, locale string, themes string, site string, pathinfo string) int {

	sqlstr := "select rowid,Created,Updated,Hits,Locale,Themes,Title,Site,Allhits from sites where Locale='" + locale + "' and Themes='" + themes + "' and Site='" + site + "' and Pathinfo='" + pathinfo + "'"

//	golog.Info("CheckMulti: " + sqlstr)

	rows, err := db.Query(sqlstr)
	if err != nil {

		golog.Crit(err.Error())
	}
	defer rows.Close()
	var recordsQuant int
	recordsQuant = 0

	for rows.Next() {

		recordsQuant = recordsQuant + 1

	}

	return recordsQuant
}
