package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type Result struct {
	useName string
	title   string
	likes   string
}

func (r Result) Data() string {
	return fmt.Sprintf("Título: %s | Criado: %s | Curtidas: %s", r.title, r.useName, r.likes)
}

func main() {
	urlList := []string{
		"https://itnext.io/golang-and-clean-architecture-19ae9aae5683",
		// "https://medium.freecodecamp.org/how-to-think-like-a-programmer-lessons-in-problem-solving-d1d8bf1de7d2",
		"https://medium.freecodecamp.org/code-comments-the-good-the-bad-and-the-ugly-be9cc65fbf83",
		// "https://uxdesign.cc/learning-to-code-or-sort-of-will-make-you-a-better-product-designer-e76165bdfc2d",
	}

	fmt.Printf("sem o uso de goroutines:\n")
	//cronometro
	init := time.Now()

	for _, url := range urlList {
		r := scrap(url)
		fmt.Println(r)
	}

	fmt.Println("tempo", time.Since(init).Seconds())
}

func scrap(url string) Result {
	var r Result
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("400: %s |\n %s", url, err.Error())
		return r
	}

	defer resp.Body.Close()

	body := resp.Body
	htmlParse, err := html.Parse(body)
	if err != nil {
		fmt.Printf("500: %s ", err.Error())
		return r
	}

	a := getFirtsTextNode(getFirstElementByClass(htmlParse, "span", "author-card-name"))
	if a != nil {
		r.useName = a.Data
	} else {
		fmt.Printf("404 useName")
	}

	h1 := getFirtsTextNode(getFirstElementByClass(htmlParse, "h1", "post-full-title"))

	if h1 != nil {
		r.title = h1.Data
	} else {
		fmt.Printf("404 title")
	}

	footer := getFirstElementByClass(htmlParse, "footer", "u-paddingTop10")
	btnLikes := getFirtsTextNode(getFirstElementByClass(footer, "button", "js-multirecommendCountButton"))

	if btnLikes != nil {
		r.likes = btnLikes.Data
	} else {
		fmt.Printf("404 likes")
	}
	return r
}

func getFirtsTextNode(htmlParse *html.Node) *html.Node {
	if htmlParse == nil {
		return nil
	}

	for m := htmlParse.FirstChild; m != nil; m = m.NextSibling {
		if m.Type == html.TextNode {
			return m
		}

		r := getFirtsTextNode(m)

		if r != nil {
			return r
		}
	}
	return nil
}

func getFirstElementByClass(htmlParsed *html.Node, elm, className string) *html.Node {
	if htmlParsed == nil {
		return nil
	}
	for m := htmlParsed.FirstChild; m != nil; m = m.NextSibling {
		if m.Data == elm && hasClass(m.Attr, className) {
			return m
		}
		r := getFirstElementByClass(m, elm, className)
		if r != nil {
			return r
		}
	}
	return nil
}

func hasClass(attribs []html.Attribute, className string) bool {
	for _, attr := range attribs {
		if attr.Key == "class" && strings.Contains(attr.Val, className) {
			return true
		}
	}
	return false
}
