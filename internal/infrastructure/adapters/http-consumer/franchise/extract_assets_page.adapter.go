package adapters

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const (
	prefixWebScrapper string = "http://"
)

func (f *FranchiseAdapterHTTP) ExtractAssetsPageAdapterHTTP(url string) (*string, error) {
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
				if strings.HasPrefix(href, "http") {
					logoURL = href
				} else {
					logoURL = url + href
				}
			}
		})
	}

	err := c.Visit(fmt.Sprint(prefixWebScrapper, url))
	if err != nil {
		return nil, err
	}
	return &logoURL, nil
}

func (f *FranchiseAdapterHTTP) ExtractFooterPageAdapterHTTP(url string) (*string, error) {
	c := colly.NewCollector()

	var logoURL string

	c.OnHTML("footer", func(e *colly.HTMLElement) {
		// // Extraer información de correo
		// email := e.ChildText(".email")

		// // Extraer información de dirección
		// address := e.ChildText(".address")

		// // Extraer información de ciudad
		// city := e.ChildText(".city")

		// // Puedes hacer algo con la información extraída, como imprimir o almacenar en una estructura de datos
		// fmt.Printf("Email: %s, Address: %s, City: %s\n", email, address, city)
	})

	err := c.Visit(fmt.Sprint(prefixWebScrapper, url))
	if err != nil {
		return nil, err
	}

	return &logoURL, nil
}
