package main

import (
	"flag"
	"fmt"
	"github.com/DiptoChakrabarty/stunning-octo-tribble/generate"
	"github.com/DiptoChakrabarty/stunning-octo-tribble/site-map"
)

func main() {
	urlFlag := flag.String("url", "https://github.com", "provide the url you wish to build sitemap for")
	maxDepth := flag.Int("depth", 2, "determines the depth of URLS you want to hop")
	maxLinks := flag.Int("limit", 500, "maximum number of links to checks")
	fName := flag.String("filename", "out.xml", "Write xml generated to this file")
	flag.Parse()
	fmt.Println(*urlFlag)

	var convertXml generate.UrlSet
	linkPages := sitemap.TraverseUrl(*urlFlag, *maxDepth, *maxLinks)
	for _, lp := range linkPages {
		//fmt.Println(lp)
		convertXml.Urls = append(convertXml.Urls, generate.XmlLoc{lp})
	}

	generate.ConvertXml(convertXml, *fName)

}
