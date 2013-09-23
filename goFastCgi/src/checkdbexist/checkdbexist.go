package checkdbexist

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
)

func Checkdb(db *sql.DB,host string, pathinfo string) (int64,string,string,string){
	var locale string
	var themes string
	var title string
	var rowid int64
	
	sqlstr := "select rowid,Locale,Themes,Title from sites where Site='" + host + "' and Pathinfo='" + pathinfo + "'"
	log.Println(sqlstr)

	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&rowid,&locale, &themes,&title)

	}
	rows.Close()

	log.Println("Checkdb:Start", locale, themes)
	
	return rowid,locale,themes,title

}
