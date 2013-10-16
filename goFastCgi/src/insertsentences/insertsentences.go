package insertsentences

import (
	"cleansentence"
//	_ "code.google.com/p/go-sqlite/go1/sqlite3"
//	"database/sql"
	"log"
	"markgen"
	ml "marklib"
	"strings"
)

//func Insert(db *sql.DB, tx *sql.Tx, c *ml.Chain, locale string, themes string, paragraphid int64) []string {

	func Insert(c *ml.Chain,locale string, themes string) []string {

	var clsentensesarr []string

	sentences := markgen.Generate(c, locale, themes)

//	stmt, err := tx.Prepare("insert into sentences(Prid,Sentence) values(?,?)")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer stmt.Close()

	for i := 0; i < len(sentences); i++ {

		words := strings.Split(sentences[i], " ")

		if len(words) < 2 {

			log.Println(sentences[i])

		} else {

			if len(sentences[i]) > 20 {
			
				clsentence := cleansentence.Clean(sentences[i])
				
				clsentensesarr = append(clsentensesarr,clsentence) 

//				if _, err = stmt.Exec(paragraphid, clsentence+"."); err != nil {
//
//					log.Fatal(err)
//				}
			}
		}
	}

return clsentensesarr

}
