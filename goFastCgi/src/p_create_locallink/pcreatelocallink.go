package p_create_locallink

import (
	"comutils"
	"math/rand"


)

func CreateLink(keywordsarr []string) string {

	rint :=rand.Intn(2)
		
	var plocallink string
	
	if rint == 0 {
	
		plocallink = "/"+keywordsarr[comutils.Random(0, len(keywordsarr))]+".html"
		
	} else {
		
		plocallink = "/"+keywordsarr[comutils.Random(0, len(keywordsarr))]+"/" +keywordsarr[comutils.Random(0, len(keywordsarr))]+".html"
	}
	
	return plocallink

}