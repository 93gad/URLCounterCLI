package flags

import (
	"flag"
	"os"
	"strings"
)

func ParseFlags() ([]string, string) {
	var urlsFlag string
	var searchStringFlag string

	osUrls := os.Getenv("urls")
	osSearch := os.Getenv("search")

	flag.StringVar(&urlsFlag, "urls", osUrls, "Список URL-адресов, разделенных запятыми")
	flag.StringVar(&searchStringFlag, "search", osSearch, "Строка для поиска")

	flag.Parse()

	strWithoutSpaces := strings.ReplaceAll(urlsFlag, " ", "")
	urls := strings.Split(strWithoutSpaces, ",")

	return urls, searchStringFlag
}
