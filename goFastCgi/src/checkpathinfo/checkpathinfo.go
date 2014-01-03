package checkpathinfo

import (
	"strings"
)

func Check(pathinfo string) string {

	thispathinfo := pathinfo

	splitPathInfo := strings.Split(pathinfo, "/")
	if len(splitPathInfo) > 0 {

		if strings.Index(splitPathInfo[len(splitPathInfo)-1], ".") == -1 {
			thispathinfo = pathinfo + "/index.html"
		}

	}

	if !strings.HasSuffix(thispathinfo, "/index.html") && !strings.HasSuffix(thispathinfo, ".gz") && strings.HasSuffix(thispathinfo,".html") {

		thispathinfo = thispathinfo + ".gz"

	}

	return thispathinfo
}
