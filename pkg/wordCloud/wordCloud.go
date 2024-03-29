package wordCloud

import (
	"image"
	"image/color"

	"github.com/psykhi/wordclouds"
)

type MaskConf struct {
	File  string     `json:"file"`
	Color color.RGBA `json:"color"`
}

type Conf struct {
	FontMaxSize     int          `json:"font_max_size"`
	FontMinSize     int          `json:"font_min_size"`
	RandomPlacement bool         `json:"random_placement"`
	FontFile        string       `json:"font_file"`
	Colors          []color.RGBA `json:"colors"`
	Width           int          `json:"width"`
	Height          int          `json:"height"`
	Mask            MaskConf     `json:"mask"`
}

func (c *Conf) calcFontMaxSize(numOfChar int) int {
	var fontMaxSize int
	fontMaxSize = int(float32(c.Width) * 0.4 / float32(numOfChar))

	return fontMaxSize
}

func (c *Conf) calcFontMinSize(numOfChar int) int {
	var fontMinSize int
	fontMinSize = int(float32(c.Width) * 0.4 / float32(numOfChar) / 10)
	return fontMinSize
}

func CreateWordCloud(wordList map[string]int, numOfChar int, colorsSetting []color.RGBA) image.Image {

	var DefaultConf = Conf{
		RandomPlacement: false,
		FontFile:        "./rounded-l-mplus-2c-medium.ttf",
		Colors:          colorsSetting,
		Width:           2048,
		Height:          2048,
		Mask: MaskConf{"", color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 0,
		}},
	}

	conf := DefaultConf

	var boxes []*wordclouds.Box
	if conf.Mask.File != "" {
		boxes = wordclouds.Mask(
			conf.Mask.File,
			conf.Width,
			conf.Height,
			conf.Mask.Color)
	}

	colors := make([]color.Color, 0)
	for _, c := range conf.Colors {
		colors = append(colors, c)
	}

	w := wordclouds.NewWordcloud(wordList,
		wordclouds.FontFile(conf.FontFile),
		wordclouds.FontMaxSize(conf.calcFontMaxSize(numOfChar)),
		wordclouds.FontMinSize(conf.calcFontMinSize(numOfChar)),
		wordclouds.Colors(colors),
		wordclouds.MaskBoxes(boxes),
		wordclouds.Height(conf.Height),
		wordclouds.Width(conf.Width),
		wordclouds.RandomPlacement(conf.RandomPlacement),
	)

	img := w.Draw()
	return img
}
