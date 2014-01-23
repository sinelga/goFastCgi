package pushdomain

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log/syslog"
)

func SelectDomain(golog syslog.Writer, locale string, themes string) string {

	var domaintopush string

	db, err := sql.Open("sqlite3", "pushdomains.db")
	defer db.Close()
	if err != nil {

		golog.Err("SelectDomain: " + err.Error())
	} else {
		sqlstr := "select Domain from pushdomains where locale='" + locale + "' and themes='" + themes + "' order by Push limit 1"

		golog.Info(sqlstr)

		rows, err := db.Query(sqlstr)
		if err != nil {

			golog.Err("newdomain: " + err.Error())
		}
		defer rows.Close()

		
		for rows.Next() {

			rows.Scan(&domaintopush)

		}

		sqlstr = "Update pushdomains set Push = Push +1 WHERE Domain=?"

		tx, err := db.Begin()
		if err != nil {
			golog.Err("SelectDomain: " + err.Error())
		}
		stmt, err := tx.Prepare(sqlstr)
		if err != nil {

			golog.Err(err.Error())
		}
		defer stmt.Close()

		if _, err = stmt.Exec(domaintopush); err != nil {

			golog.Err("SelectDomain: " + err.Error())

		}
		stmt.Close()
		tx.Commit()

	}

	return domaintopush
}
