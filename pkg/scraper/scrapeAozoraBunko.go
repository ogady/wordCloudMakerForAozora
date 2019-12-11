package scraper

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/ogady/wordCloudMakerForAozora/pkg/decoder"
)

func Scrape(url string) (string, error) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		err = fmt.Errorf("Document Constructorの初期化に失敗しました。\n %w", err)
		return "", err
	}

	selection := doc.Find("body > div.main_text")
	text := selection.Text()

	encodedText, err := decoder.Decode("ShiftJIS", []byte(text))
	if err != nil {
		err = fmt.Errorf("UTF8への変換に失敗しました。 \n %w", err)
		return "", err
	}

	return string(encodedText), nil
}
