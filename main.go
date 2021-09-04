package main

import (
	"flag"
	"fmt"
	"github.com/DiptoChakrabarty/stunning-octo-tribble/link"
	"github.com/DiptoChakrabarty/stunning-octo-tribble/site-map"
	"net/http"
)

func main() {
	urlFlag := flag.String("url", "https://github.com", "provide the url you wish to build sitemap for")
	flag.Parse()
	fmt.Println(*urlFlag)

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	links, _ := link.Parse(resp.Body)

	for _, l := range links {
		fmt.Println(l)
	}

	base := sitemap.GenerateBaseUrl(resp) // the original base URL from the url
	fmt.Println(base)

	hrefs := sitemap.GenerateLinks(links, base)
	for _, href := range hrefs {
		fmt.Println(href)
	}
}
