package main

import (
	"fmt"
	"strings"

	"github.com/DiptoChakrabarty/stunning-octo-tribble/link"
)

var example = `
<html>
<body>
<h1> Hello </h1>
<a href="/other-page">Link to another page</a>
<a href="/chalo-link">Chalo Theek hai Link</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(example)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
