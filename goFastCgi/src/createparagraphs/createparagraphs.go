package createparagraphs

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"comutils"
//	"database/sql"
	"insertsentences"
	"io/ioutil"
	"log"
	"log/syslog"
	ml "marklib"
	"math/rand"
	"p_create_locallink"
	"prtitlegen"
	"selectmarkfile"
	"time"
	"domains"
	"queue/freeparagraphs"
)

var markfile string
var paragraphs []domains.Paragraph

func CreatePr(golog syslog.Writer, locale string, themes string, keywords []string, phrases []string, hosts []string, quant int) {

//	db, err := sql.Open("sqlite3", "gofast.db")
//	if err != nil {
//		log.Fatal(err)
//	}

	log.Println("start CreatePr")
	prefixLen := 1

//	tx, err := db.Begin()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	stmt, err := tx.Prepare("insert into paragraphs(Created,Locale,Themes,Ptitle,Pphrase,Host,Locallink) values(?,?,?,?,?,?,?)")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer stmt.Close()

	markfile = selectmarkfile.SelectFile(golog, locale, themes)

	log.Println("markfile -> ", markfile)
	golog.Info("markfile -> " + markfile)

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

//	var rs sql.Result
	

	for i := 0; i < quant; i++ {
	
		var clsentensesarr []string
		
		prtitle := prtitlegen.Generate(keywords)
		prphrase := comutils.UpcaseInitial(phrases[rand.Intn(len(phrases))]) + "."
		host := hosts[rand.Intn(len(hosts))]
		locallink := p_create_locallink.CreateLink(keywords)

			clsentensesarr = insertsentences.Insert(c, locale, themes)

//		}
		
		paragraph := domains.Paragraph{
		
			Ptitle: prtitle,
			Pphrase: prphrase,
			Plocallink: locallink,
			Phost: host,
			Sentences: clsentensesarr,
		
		}
		paragraphs = append(paragraphs,paragraph)

	}
//	tx.Commit()
	
	freeparagraphs.CreateParagraphs(golog, locale,themes,paragraphs)
	

}
