package getalldbparagraphs

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"domains"
	"log"
	"strconv"
	"strings"
)

func GetAllPr(db *sql.DB, siteid int64,site string) []domains.Paragraph {

//	log.Println("GetAllPr:Start")

	var paragraphsarr []domains.Paragraph

	sqlstr := "select rowid,Ptitle,Pphrase,Host,Locallink from paragraphs where Siteid =" + strconv.FormatInt(siteid, 10)
	log.Println(sqlstr)

	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {

		var prid int64
		var ptitle string
		var pphrase string
		var phost string
		var plocallink string
		var locallink string

		rows.Scan(&prid, &ptitle, &pphrase,&phost,&plocallink)

		domainarr := strings.Split(site,".")
		
		domain := domainarr[len(domainarr)-2]+"."+domainarr[len(domainarr)-1]
		
		locallink =phost+"."+domain+ plocallink
		log.Println(locallink) 

		paragraph := domains.Paragraph{

			Rowid:   prid,
			Siteid:  siteid,
			Ptitle:  ptitle,
			Plocallink: locallink,
			Pphrase: pphrase,
		}

		paragraphsarr = append(paragraphsarr, paragraph)

	}
	rows.Close()
//	log.Println("GetAllPr:paragraphsarr len",len(paragraphsarr))
	return paragraphsarr

}
