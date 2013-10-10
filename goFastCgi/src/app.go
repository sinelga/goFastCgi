package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"createfirstpage"
	"database/sql"
	//	"encoding/json"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"sync"
	//	"strings"
	"clean_pathinfo"
	"htmlfileexist"
)

var startOnce sync.Once
var keywordsarr_fi_FI_finance []string
var phrasesarr_fi_FI_finance []string
var keywordsarr_fi_FI_porno []string
var phrasesarr_fi_FI_porno []string
var keywordsarr_it_IT_finance []string
var phrasesarr_it_IT_finance []string
var keywordsarr_fi_FI_fortune []string
var phrasesarr_fi_FI_fortune []string

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	themes := req.Header.Get("X-THEMES")
	locale := req.Header.Get("X-LOCALE")
	host := req.Header.Get("X-DOMAIN")
	pathinfo := req.Header.Get("X-PATHINFO")

	checkfirstpage(resp, req, locale, themes, host, pathinfo)

}

func main() {

	log.Println("Start")
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)
}

func checkfirstpage(resp http.ResponseWriter, req *http.Request, locale string, themes string, host string, pathinfo string) {

	log.Println(locale)
	log.Println(themes)
	log.Println(host)
	log.Println(pathinfo)

	startOnce.Do(func() {
		startones()
	})
	pathinfostr := clean_pathinfo.CleanPath(pathinfo)

	htmlfile := string("www/" + locale + "/" + themes + "/" + host + pathinfostr)
	log.Println(htmlfile)
	if _, err := os.Stat(htmlfile); err != nil {

		if os.IsNotExist(err) {

			log.Println("file does not exist")
			if locale == "fi_FI" && themes == "finance" {
				createfirstpage.CreatePage(locale, themes, host, pathinfostr, keywordsarr_fi_FI_finance, phrasesarr_fi_FI_finance)
			} else if locale == "fi_FI" && themes == "porno" {
				createfirstpage.CreatePage(locale, themes, host, pathinfostr, keywordsarr_fi_FI_porno, phrasesarr_fi_FI_porno)

			} else if locale == "it_IT" && themes == "finance" {
				createfirstpage.CreatePage(locale, themes, host, pathinfostr, keywordsarr_it_IT_finance, phrasesarr_it_IT_finance)
			}  else if locale == "fi_FI" && themes == "fortune" {
				createfirstpage.CreatePage(locale, themes, host, pathinfostr, keywordsarr_fi_FI_fortune, phrasesarr_fi_FI_fortune)
			}
			
			
			http.ServeFile(resp, req, htmlfile)

		} else {
			// other error
			log.Println("something wrong???")

		}

	} else {
		log.Println("fileexist", htmlfile, "host", host, "pathinfostr", pathinfostr)
		http.ServeFile(resp, req, htmlfile)
		go htmlfileexist.StartCheck(htmlfile, host, pathinfostr)

	}
}

func startones() {

	log.Println("startones")
	db, err := sql.Open("sqlite3", "singo.db")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select Locale,Themes,Keyword from keywords")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var locale string
		var themes string
		var keyword string

		rows.Scan(&locale, &themes, &keyword)
		if locale == "fi_FI" && themes == "finance" {
			keywordsarr_fi_FI_finance = append(keywordsarr_fi_FI_finance, keyword)
		} else if locale == "fi_FI" && themes == "porno" {
			keywordsarr_fi_FI_porno = append(keywordsarr_fi_FI_porno, keyword)

		} else if locale == "it_IT" && themes == "finance" {
			keywordsarr_it_IT_finance = append(keywordsarr_it_IT_finance, keyword)

		} else if locale == "fi_FI" && themes == "fortune" {
			keywordsarr_fi_FI_fortune = append(keywordsarr_fi_FI_fortune, keyword)
		} 

	}
	rows.Close()
	log.Println("keywordsarr_fi_FI_finance", len(keywordsarr_fi_FI_finance))
	log.Println("keywordsarr_fi_FI_porno", len(keywordsarr_fi_FI_porno))
	log.Println("keywordsarr_it_IT_finance", len(keywordsarr_it_IT_finance))
	log.Println("keywordsarr_fi_FI_fortune", len(keywordsarr_fi_FI_fortune))

	rows, err = db.Query("select Locale,Themes,Phrase from phrases")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var locale string
		var themes string
		var phrase string
		rows.Scan(&locale, &themes, &phrase)

		if locale == "fi_FI" && themes == "finance" {
			phrasesarr_fi_FI_finance = append(phrasesarr_fi_FI_finance, phrase)
		} else if locale == "fi_FI" && themes == "porno" {
			phrasesarr_fi_FI_porno = append(phrasesarr_fi_FI_porno, phrase)
		} else if locale == "it_IT" && themes == "finance" {
			phrasesarr_it_IT_finance = append(phrasesarr_it_IT_finance, phrase)
		} else if locale == "fi_FI" && themes == "fortune" {
			phrasesarr_fi_FI_fortune = append(phrasesarr_fi_FI_fortune, phrase)
		}

	}
	rows.Close()
	log.Println("phrasesarr_fi_FI_finance", len(phrasesarr_fi_FI_finance))
	log.Println("phrasesarr_fi_FI_porno", len(phrasesarr_fi_FI_porno))
	log.Println("phrasesarr_it_IT_finance", len(phrasesarr_it_IT_finance))
	log.Println("phrasesarr_fi_FI_fortune", len(phrasesarr_fi_FI_fortune))
	db.Close()

}
