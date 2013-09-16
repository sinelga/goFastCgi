package domains

import (
	"time"
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
//type Keywordkey struct {
//	Key *datastore.Key
//	Keyword string
//
//}
//
//type Phrasekey struct {
//	Key *datastore.Key
//	Phrase  string
//
//}

//type Restorekeys struct {
//	Keysarr []*datastore.Key
//}

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

type WebContents struct {
	Locale   string
	Themes   string
	Site     string
	PathInfo string
	Created  time.Time
	Updated  time.Time
	Hits     int
	AllHits  int
	//	Mtext    []string
	Mtext   []byte
	WebPage []byte
	//	KeyWords []string
//	KeyWords []byte
	Title    string
//	Phrases  []byte
	Somephrases []string
	Somekeywords []string
}
