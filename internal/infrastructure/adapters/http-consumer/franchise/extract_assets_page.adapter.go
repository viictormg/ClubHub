package adapters

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const (
	prefixWebScrapper string = "http://"
)

func (f *FranchiseAdapter) ExtractAssetsPageAdapter(url string) (*string, error) {
	c := colly.NewCollector()

	var logoURL string

	selectors := []string{
		"img[alt*='logo']",
		"img[src*='logo']",
		"img[class*='logo']",
		"img[id*='logo']",
		"div[class*='logo'] img",
		"header img",
		"[role*='banner'] img",
		"[aria-label*='logo'] img",
		"link[rel='icon'], link[rel='shortcut icon']",
	}

	for _, selector := range selectors {
		c.OnHTML(selector, func(e *colly.HTMLElement) {
			href := e.Attr("href")
			if href != "" {
				// Comprobar si la URL es relativa o absoluta
				if strings.HasPrefix(href, "http") {
					logoURL = href
				} else {
					// Convertir la URL relativa a absoluta
					logoURL = url + href
				}
			}
		})
	}

	// Visitar la URL y manejar errores
	logoURL = "imagen-test"
	err := c.Visit(fmt.Sprint(prefixWebScrapper, url))
	if err != nil {
		return nil, err
	}

	return &logoURL, nil
}
