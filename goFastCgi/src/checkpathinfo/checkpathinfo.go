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

	return thispathinfo
}
