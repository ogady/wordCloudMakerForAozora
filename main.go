package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/ogady/wordCloudMaker/decoder"
	"github.com/ogady/wordCloudMaker/morphoAnalyzer"
	"github.com/ogady/wordCloudMaker/wordCloud"
)

func main() {

	var config = flag.String("config", "config.json", "path to config file")
	var output = flag.String("output", "output.png", "path to output image")

	url := "https://www.aozora.gr.jp/cards/001235/files/49866_41897.html"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	// ドキュメントから
	selection := doc.Find("body > div.main_text")
	text := selection.Text()

	encoded, err := decoder.Decode("ShiftJIS", []byte(text))

	if err != nil {
		fmt.Println(err)
	}

	persedText := morphoAnalyzer.ParseToNode(string(encoded))
	img := wordCloud.CreateWordCloud(persedText, config)
	outputFile, err := os.Create(*output)
	if err != nil {
		// Handle error
	}
	start := time.Now()
	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, img)

	// Don't forget to close files
	outputFile.Close()
	fmt.Printf("Done in %v\n", time.Since(start))
}
