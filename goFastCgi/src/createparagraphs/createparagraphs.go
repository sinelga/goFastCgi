package createparagraphs

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"comutils"
	//	"database/sql"
	"domains"
	"insertsentences"
	"io/ioutil"
	"log"
	"log/syslog"
	ml "marklib"
	"math/rand"
	"p_create_locallink"
	"prtitlegen"
	"queue/freeparagraphs"
	"selectmarkfile"
	"time"
)

var markfile string
var paragraphs []domains.Paragraph

func CreatePr(golog syslog.Writer, locale string, themes string, keywords []string, phrases []string, hosts []string, quant int) {

	log.Println("start CreatePr")
	prefixLen := 1

	markfile = selectmarkfile.SelectFile(golog, locale, themes)

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

	for i := 0; i < quant; i++ {

		var clsentensesarr []string

		prtitle := prtitlegen.Generate(keywords)
		prphrase := comutils.UpcaseInitial(phrases[rand.Intn(len(phrases))]) + "."
		host := hosts[rand.Intn(len(hosts))]
		locallink := p_create_locallink.CreateLink(keywords)

		clsentensesarr = insertsentences.Insert(c, locale, themes)

		paragraph := domains.Paragraph{

			Ptitle:     prtitle,
			Pphrase:    prphrase,
			Plocallink: locallink,
			Phost:      host,
			Sentences:  clsentensesarr,
		}
		paragraphs = append(paragraphs, paragraph)

	}

	freeparagraphs.CreateParagraphs(golog, locale, themes, paragraphs)

}
