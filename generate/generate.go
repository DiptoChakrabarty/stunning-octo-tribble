package generate

import (
	"encoding/xml"
	"os"
)

const namespace = "http://www.sitemaps.org/schemas/sitemap/0.9"

type XmlLoc struct {
	Value string `xml:"loc"`
}

type UrlSet struct {
	Urls       []XmlLoc `xml:"url"`
	XNamespace string   `xml:"xmlns,attr"`
}

func ConvertXml(convert UrlSet, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	convert.XNamespace = namespace
	enc := xml.NewEncoder(f)
	enc.Indent("", " ")
	if err := enc.Encode(convert); err != nil {
		panic(err)
	}
}
