package internal

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
	"time"

	"github.com/ogady/wordCloudMakerForAozora/pkg/aozora"
	"github.com/ogady/wordCloudMakerForAozora/pkg/morphoAnalyzer"
	"github.com/ogady/wordCloudMakerForAozora/pkg/scraper"
	"github.com/ogady/wordCloudMakerForAozora/pkg/wordCloud"
)

type WordCloudCreater struct {
	output         string
	titleName      string
	specifiedColor string
}

func NewWordCloudCreater(output, titleName, specifiedColor string) WordCloudCreater {

	var wordCloudCreater WordCloudCreater

	wordCloudCreater = WordCloudCreater{
		output:         output,
		titleName:      titleName,
		specifiedColor: specifiedColor,
	}
	return wordCloudCreater

}

func (w *WordCloudCreater) Execute() error {
	colorsSetting := []color.RGBA{}

	switch w.specifiedColor {
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

	htmlURL, err := aozora.GetBookInfoByTitleName(w.titleName)

	if err != nil {
		err = fmt.Errorf("作品情報が取得できませんでした。 作品名：%s\n %w", w.titleName, err)
		fmt.Println(err)
		os.Exit(1)
	}

	text := scraper.Scrape(htmlURL)

	persedText := morphoAnalyzer.ParseToNode(string(text))
	maxCount := morphoAnalyzer.GetMaxCount(persedText)
	numOfChar := len([]rune(maxCount))

	fmt.Println(numOfChar)

	img := wordCloud.CreateWordCloud(persedText, numOfChar, colorsSetting)
	outputFile, err := os.Create(w.output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	start := time.Now()

	png.Encode(outputFile, img)

	outputFile.Close()
	fmt.Printf("Done in %v\n", time.Since(start))

	return nil
}
