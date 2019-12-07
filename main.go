package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/ogady/wordCloudMakerForAozora/morphoAnalyzer"
	"github.com/ogady/wordCloudMakerForAozora/scraper"
	"github.com/ogady/wordCloudMakerForAozora/wordCloud"
)

func main() {

	var (
		output = flag.String("output", "output.png", "path to output image")
		url    = flag.String("url", "https://www.aozora.gr.jp/cards/000081/files/43737_19215.html", "Target URL")
	)

	flag.Parse()

	text := scraper.Scrape(*url)

	persedText := morphoAnalyzer.ParseToNode(string(text))
	maxCount := morphoAnalyzer.GetMaxCount(persedText)
	numOfChar := len([]rune(maxCount))

	fmt.Println(numOfChar)

	img := wordCloud.CreateWordCloud(persedText, numOfChar)
	outputFile, err := os.Create(*output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	start := time.Now()

	png.Encode(outputFile, img)

	outputFile.Close()
	fmt.Printf("Done in %v\n", time.Since(start))
}
