package ScrapeLinks

import (
	"fmt"
	"strconv"
)

func Print(result map[string][]Ref) {
	counter := 1;
	for url, links := range(result) {
		fmt.Println(fmt.Sprintf("\n%s) %s:\n\n   links:\n", strconv.Itoa(counter), url))
		for _, link := range(links) {
			fmt.Println(fmt.Sprintf("      Href: %s, Title: %s", link.Href, link.Title))
		}
		counter = counter + 1
	}
 }
