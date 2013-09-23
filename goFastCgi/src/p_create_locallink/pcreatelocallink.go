package p_create_locallink

import (
	"comutils"

)

func CreateLink(keywordsarr []string) string {

	plocallink := "/"+keywordsarr[comutils.Random(0, len(keywordsarr))]+".html"
	
	return plocallink

}