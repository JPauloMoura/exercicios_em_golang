package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

const BODY = `
<body>
	<p>Links:</p>
	<ul>
		<li><a href="http://web1">web1</a></li>
		<li><a href="http://web2">web2</a></li>
	</ul>
</body>
`

type Links struct {
	urls []string
}

func findLinks() {
	page, err := html.Parse(strings.NewReader(BODY))
	if err != nil {
		panic(err)
	}

	var l Links

	find(&l, page)
	fmt.Println(l.urls)
}

// find intera sobre os elementos htmls buscando aqueles que possue o atributo href
// ao encontrar, seta o valor dele em um slice e busca o proximo elemento
func find(links *Links, n *html.Node) {
	if n.Type == html.ElementNode {
		for _, att := range n.Attr {
			if att.Key == "href" {
				links.urls = append(links.urls, att.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		find(links, c)
	}
}
