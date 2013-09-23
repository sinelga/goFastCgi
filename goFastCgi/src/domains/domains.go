package domains

import (
	"time"
//	"domains"
//	"appengine/datastore"
)

type Tsentences struct {
	Sentences []string
}

type MarkovTask struct {
	Sentences []string
}

type Keyword struct {
	Locale  string
	Themes  string
	Created time.Time
	Updated time.Time
	Keyword string
	Rating  int
	Hits int
}

type Phrase struct {
	Locale  string
	Themes  string
	Created time.Time
	Updated time.Time
	Phrase  string
	Grating int
	Urlstr  string
}

type WebInfo struct {
	Site       string
	Ext        string
	Sentences  []string
	Keywords   []string
	UpKeywords []string
	Title      string
	Phrases    []string
}

type Paragraph struct {

	Rowid	int64
	Siteid int64
	Created  time.Time
	Ptitle string
	Pphrase string
	Plocallink string
	Sentences []string	 	

}

type WebContents struct {
	Rowid	int64
	Locale   string
	Themes   string
	Title	string
	Site     string
	PathInfo string
	Created  time.Time
	Updated  time.Time
	Hits     int
	AllHits  int
	Paragraphs []Paragraph
	
//	//	Mtext    []string
//	Mtext   []byte
//	WebPage []byte
//	//	KeyWords []string
////	KeyWords []byte
//	Title    string
////	Phrases  []byte
//	Somephrases []string
//	Somekeywords []string
}
