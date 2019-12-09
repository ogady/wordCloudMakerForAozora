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
		return err
	}

	text, err := scraper.Scrape(htmlURL)
	if err != nil {
		err = fmt.Errorf("本情報のスクレイピングに失敗しました。\n%w", err)
		return err
	}

	persedText, err := morphoAnalyzer.ParseToNode(string(text))
	if err != nil {
		err = fmt.Errorf("形態素解析に失敗しました。。\n%w", err)
		return err
	}

	maxCount := morphoAnalyzer.GetMaxCount(persedText)
	numOfChar := len([]rune(maxCount))

	img := wordCloud.CreateWordCloud(persedText, numOfChar, colorsSetting)

	outputFile, err := os.Create(w.output)
	if err != nil {
		err = fmt.Errorf("WordCloud作成に失敗しました。\n%w", err)
		return err
	}

	start := time.Now()

	err = png.Encode(outputFile, img)
	if err != nil {
		err = fmt.Errorf("pngのエンコードに失敗しました。%w", err)
		return err
	}

	outputFile.Close()
	fmt.Printf("Done in %v\n", time.Since(start))

	return nil
}
