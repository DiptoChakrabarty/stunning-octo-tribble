package sitemap

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/DiptoChakrabarty/stunning-octo-tribble/link"
)

func GenerateBaseUrl(resp *http.Response) string {
	redirectUrl := resp.Request.URL // in case users give http redirects to https taken
	baseUrl := &url.URL{
		Scheme: redirectUrl.Scheme,
		Host:   redirectUrl.Host,
	}
	base := baseUrl.String() // the original base URL from the url
	return base
}

func GenerateLinks(links []link.Link, base string) []string {
	var result []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			result = append(result, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			result = append(result, l.Href)
		}
	}
	return result
}
