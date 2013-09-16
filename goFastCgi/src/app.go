package main

import (
	"encoding/json"
//	"fmt"
//	"html/template"
//	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"createfirstpage"
)

type FastCGIServer struct{}

func serveSingle(pattern string, filename string) {
    http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filename)
    })
}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	//	resp.Write([]byte("<h1>Hello,</h1>\n<p>Behold my Go web app.</p>"))
	
		themes :=req.Header.Get("X-THEMES")
		locale :=req.Header.Get("X-LOCALE")
		host :=req.Header.Get("X-DOMAIN")
		pathinfo :=req.Header.Get("X-PATHINFO")
//		queue :=req.Header.Get("X-QUEUE")

	//	fields := make(map[string] string)

//	lst := LoadJson()
//	if req.Method == "POST" {
//		entry := make(map[string]string)
//		entry["author"] = req.FormValue("author")
//		entry["content"] = req.FormValue("content")
//		lst = append(lst, entry)
//		SaveJson(lst)
//	}

//	var tmpl = template.Must(template.New("templateout").Parse("template.html"))
//	template := template.   .MustParseFiles("template.html", nil)
//	tmpl := template.Must(template.New("templateout").ParseFiles("template.html"))
//	tmpl.Execute(resp, lst)

	//    templateout,err := template.New("templateout").Parse("template.html"); err != nil {
	//
	//    	log.Println("err.Error")
	////    	fmt.Fprintln(err.Error())
	//    } else {
	//
	//
	//    }

	//    template.Execute(resp, fields)

	checkfirstpage(resp,req ,locale,themes,host,pathinfo)

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

func checkfirstpage(resp http.ResponseWriter, req *http.Request,locale string, themes string, host string, pathinfo string) {

	log.Println(locale)
	log.Println(themes)
	log.Println(host)
	log.Println(pathinfo)

	htmlfile := string("www/" +locale+"/"+themes+"/"+host+pathinfo)
//	log.Println(htmlfile )
	if _, err := os.Stat(htmlfile); err != nil {

		if os.IsNotExist(err) {
			
			log.Println("file does not exist")
			
			go createfirstpage.CreatePage(locale,themes,host,pathinfo)
			
			http.ServeFile(resp,req,"www/firstpage.html")

			

		} else {
			// other error
			log.Println("somthing wrong???")

		}

	} else {
		log.Println("fileexist")
		http.ServeFile(resp,req,htmlfile)

	}

}
