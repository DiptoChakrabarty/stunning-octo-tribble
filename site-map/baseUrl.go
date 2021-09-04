package sitemap

import (
	"github.com/DiptoChakrabarty/stunning-octo-tribble/link"
	"net/http"
	"net/url"
	"strings"
)

func GenerateBaseUrl(UrlName *string) (string, []link.Link) {
	resp, err := http.Get(*UrlName)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	links, _ := link.Parse(resp.Body)

	redirectUrl := resp.Request.URL // in case users give http redirects to https taken
	baseUrl := &url.URL{
		Scheme: redirectUrl.Scheme,
		Host:   redirectUrl.Host,
	}
	base := baseUrl.String() // the original base URL from the url
	return base, links
}

func GenerateSiteLinks(links []link.Link, base string) []string {
	var result []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			result = append(result, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			result = append(result, l.Href)
		}
	}
	return FilterLinks(base, result)
}
