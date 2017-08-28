package ScrapeLinks

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

type Ref struct {
	Title string
	Href string
}

type RefsList struct {
	Url string
	Links []Ref
}

func Run(url string, c chan RefsList) {
	var links []Ref

 	doc, err := goquery.NewDocument(url)

 	if err != nil {
 		panic(err)
 	}

 	doc.Find("a").Each(func(i int, s *goquery.Selection) {
 		Title := strings.TrimSpace(s.Text())
		Href, _ := s.Attr("href")
		ref := Ref{
			Title: Title,
			Href: Href,
	    }
		links = append(links, ref)
 	})
	res := RefsList{
		Url: url,
		Links: links,
	}

	c <- res
 }
