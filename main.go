package main

import (
	"flag"
	"fmt"
	"github.com/DiptoChakrabarty/stunning-octo-tribble/site-map"
)

func main() {
	urlFlag := flag.String("url", "https://github.com", "provide the url you wish to build sitemap for")
	maxdepth := flag.Int("depth", 2, "determines the depth of URLS you want to hop")
	flag.Parse()
	fmt.Println(*urlFlag)

	base, links := sitemap.GenerateBaseUrl(urlFlag) // the original base URL from the url
	fmt.Println(base)

	pages := sitemap.GenerateSiteLinks(links, base)
	for _, href := range pages {
		fmt.Println(href)
	}

	fmt.Println("The New Pages")

	linkPages := sitemap.TraverseUrl(*urlFlag, *maxdepth)
	for _, lp := range linkPages {
		fmt.Println(lp)
	}

}
