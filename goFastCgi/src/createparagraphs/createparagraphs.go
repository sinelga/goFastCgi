package createparagraphs

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"insertsentences"
	"comutils"
	"prtitlegen"
	"p_create_locallink"
	"io/ioutil"
	"log"
	ml "marklib"
	"math/rand"
	"time"
)

//var keywordsarr_fi_FI_finance []string
var markfile string

func CreatePr(db *sql.DB, locale string, themes string, keywords []string, phrases []string, hosts []string,quant int) {

	log.Println("start CreatePr")
	prefixLen := 1

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into paragraphs(Created,Locale,Themes,Ptitle,Pphrase,Host,Locallink) values(?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var rs sql.Result

	if locale == "fi_FI" && themes == "finance" {

		markfile = "markresources/fi_FI_finance.txt"
	}

	//For start Mark
	rand.Seed(time.Now().UnixNano())
	c := ml.NewChain(prefixLen)
	fData, err := ioutil.ReadFile(markfile)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	c.Write(fData)
	// end For start Mark
	
	for i := 0; i < quant; i++ {
	
		prtitle := prtitlegen.Generate(keywords)
		prphrase := comutils.UpcaseInitial(phrases[rand.Intn(len(phrases))])+"."
		host := hosts[rand.Intn(len(hosts))]
		locallink := p_create_locallink.CreateLink(keywords)
	
		now := time.Now().Unix()
		if rs, err = stmt.Exec(now,locale,themes,prtitle,prphrase,host,locallink); err != nil {

			log.Fatal(err)

		} else {

			paragraphid, _ := rs.LastInsertId()
			log.Println(paragraphid)
			insertsentences.Insert(db, tx, c, locale, themes, paragraphid)

		}

	}
	tx.Commit()

}
