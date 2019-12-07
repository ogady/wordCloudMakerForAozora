package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"os"
	"time"

	"github.com/ogady/wordCloudMakerForAozora/aozora"
	"github.com/ogady/wordCloudMakerForAozora/morphoAnalyzer"
	"github.com/ogady/wordCloudMakerForAozora/scraper"
	"github.com/ogady/wordCloudMakerForAozora/wordCloud"
)

func main() {

	var (
		output         = flag.String("o", "output.png", "path to output image")
		titleName      = flag.String("t", "銀河鉄道の夜", "Target TitleNames")
		specifiedColor = flag.String("c", "red", "Specify the color to draw from ’red’, ’blue’, ’green’, and ’vivid’.")
	)

	flag.Parse()

	colorsSetting := []color.RGBA{}

	switch *specifiedColor {
	case "red":
		colorsSetting = wordCloud.RedColors
	case "blue":
		colorsSetting = wordCloud.BlueColors
	case "green":
		colorsSetting = wordCloud.GreenColors
	case "vivid":
		colorsSetting = wordCloud.VividColors
	default:
		fmt.Println("色指定がなかったのでデフォルト配色で描画します。")
		colorsSetting = wordCloud.DefaultColors
	}

	htmlURL, err := aozora.GetBookInfoByTitleName(*titleName)

	if err != nil {
		err = fmt.Errorf("作品情報が取得できませんでした。 作品名：%s\n %w", *titleName, err)
		fmt.Println(err)
		os.Exit(1)
	}

	text := scraper.Scrape(htmlURL)

	persedText := morphoAnalyzer.ParseToNode(string(text))
	maxCount := morphoAnalyzer.GetMaxCount(persedText)
	numOfChar := len([]rune(maxCount))

	fmt.Println(numOfChar)

	img := wordCloud.CreateWordCloud(persedText, numOfChar, colorsSetting)
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
