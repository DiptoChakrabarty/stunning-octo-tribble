package sitemap

import (
	"strings"
)

func FilterLinks(base string, links []string) []string {
	var result []string
	for _, l := range links {
		if strings.HasPrefix(l, base) {
			result = append(result, l)
		}
	}
	return result
}
