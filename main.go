package main

import (
	"flag"
	"fmt"
	"github.com/DiptoChakrabarty/stunning-octo-tribble/site-map"
)

func main() {
	urlFlag := flag.String("url", "https://github.com", "provide the url you wish to build sitemap for")
	flag.Parse()
	fmt.Println(*urlFlag)

	base, links := sitemap.GenerateBaseUrl(urlFlag) // the original base URL from the url
	fmt.Println(base)

	pages := sitemap.GenerateLinks(links, base)
	for _, href := range pages {
		fmt.Println(href)
	}
}
