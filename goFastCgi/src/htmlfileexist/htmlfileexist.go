package htmlfileexist

import (
	"addfreeparagraph"
	"checkdbexist"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"createpage"
	"database/sql"
	"domains"
	"getalldbparagraphs"
	"log"
	"sentencesforpr"
)

func StartCheck(htmlfile string, host string, pathinfo string) {

	log.Println("StartCheck", pathinfo)

	var webcontents = domains.WebContents{}
	var rowid int64
	var locale string
	var themes string
	var title string
	var deltamin int
	var paragraphsarr []domains.Paragraph

	db, err := sql.Open("sqlite3", "gofast.db")
	if err != nil {
		log.Fatal(err)
	}

	rowid, locale, themes, title, deltamin = checkdbexist.Checkdb(db, host, pathinfo)

	webcontents.Rowid = rowid
	webcontents.Locale = locale
	webcontents.Themes = themes
	webcontents.Title = title
	webcontents.Site = host
	webcontents.PathInfo = pathinfo

	if deltamin > 5 {

		paragraphsarr = getalldbparagraphs.GetAllPr(db, rowid, host)

		for i, paragraph := range paragraphsarr {

			paragraphsarr[i].Sentences = sentencesforpr.GetSents(db, paragraph.Rowid)
		}

		addfreeparagraph.AddPr(db, rowid, locale, themes)

		if db.Close(); err != nil {

			log.Fatal(err)
		} else {

			webcontents.Paragraphs = paragraphsarr
			createpage.CreatePg(htmlfile, webcontents)

		}
	} else {
		
		log.Println("Dont Create new PAGE!!")
		
	}

}