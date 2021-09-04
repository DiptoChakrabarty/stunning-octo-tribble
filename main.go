package main

import (
	"flag"
	"fmt"
	"github.com/DiptoChakrabarty/stunning-octo-tribble/link"
	"net/http"
	//"strings"
	//"github.com/DiptoChakrabarty/stunning-octo-tribble/link"
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
}
