package scraper

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/ogady/wordCloudMaker/decoder"
)

func Scrape(url string) string {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ドキュメントから
	selection := doc.Find("body > div.main_text")
	text := selection.Text()

	encodedText, err := decoder.Decode("ShiftJIS", []byte(text))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(encodedText)
}
