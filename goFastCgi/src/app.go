package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"createfirstpage"
	"database/sql"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"sync"
)

var startOnce sync.Once
var keywordsarr_fi_FI_finance []string
var phrasesarr_fi_FI_finance []string

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	//	resp.Write([]byte("<h1>Hello,</h1>\n<p>Behold my Go web app.</p>"))

	themes := req.Header.Get("X-THEMES")
	locale := req.Header.Get("X-LOCALE")
	host := req.Header.Get("X-DOMAIN")
	pathinfo := req.Header.Get("X-PATHINFO")

	checkfirstpage(resp, req, locale, themes, host, pathinfo)

}

func main() {

	listener, _ := net.Listen("tcp", "127.0.0.1:9000")
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)
}

func SaveJson(j []map[string]string) {
	file, err := os.Create("data.json")
	if err == nil {
		json := json.NewEncoder(file)
		json.Encode(j)
		file.Close()
	}
}
func LoadJson() []map[string]string {
	lst := make([]map[string]string, 0)
	file, err := os.Open("data.json")
	if err != nil {
		return lst
	}
	json := json.NewDecoder(file)
	json.Decode(&lst)
	file.Close()
	return lst
}

func checkfirstpage(resp http.ResponseWriter, req *http.Request, locale string, themes string, host string, pathinfo string) {

	log.Println(locale)
	log.Println(themes)
	log.Println(host)
	log.Println(pathinfo)

	startOnce.Do(func() {
		startones()
	})
	
	var pathinfostr string
	if pathinfo == "/" {
		
		pathinfostr ="/index.html"
	
	} else {
		pathinfostr = pathinfo
	
	}
			
	htmlfile := string("www/" + locale + "/" + themes + "/" + host + pathinfostr)
//	log.Println(htmlfile )
	if _, err := os.Stat(htmlfile); err != nil {

		if os.IsNotExist(err) {

			log.Println("file does not exist")

			go createfirstpage.CreatePage(locale, themes, host, pathinfostr,keywordsarr_fi_FI_finance,phrasesarr_fi_FI_finance)

			http.ServeFile(resp, req, "www/firstpage.html")

		} else {
			// other error
			log.Println("somthing wrong???")

		}

	} else {
		log.Println("fileexist")
		http.ServeFile(resp, req, htmlfile)
	}
}

func startones() {

	log.Println("startones")
	db, err := sql.Open("sqlite3", "singo.db")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select keyword from keywords where locale='fi_FI' and themes='finance'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var keyword string
		rows.Scan(&keyword)
		keywordsarr_fi_FI_finance = append(keywordsarr_fi_FI_finance, keyword)

	}
	rows.Close()
	log.Println(len(keywordsarr_fi_FI_finance))

	rows, err = db.Query("select phrase from phrases where locale='fi_FI' and themes='finance'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var phrase string
		rows.Scan(&phrase)
		phrasesarr_fi_FI_finance = append(phrasesarr_fi_FI_finance, phrase)

	}
	rows.Close()
	log.Println(len(phrasesarr_fi_FI_finance))

}
