package wordCloud

import (
	"flag"
	"image"
	"image/color"

	"github.com/psykhi/wordclouds"
)

func CreateWordCloud(wordList map[string]int, numOfChar int) image.Image {

	var DefaultColors = []color.RGBA{
		{0x00, 0x76, 0x2d, 0xff},
		{0x43, 0x76, 0x2d, 0xff},
		{0x73, 0x76, 0x2d, 0xff},
		{0x90, 0x76, 0x2d, 0xff},
		{0xb7, 0x7c, 0x2d, 0xff},
	}

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

	var DefaultConf = Conf{
		RandomPlacement: false,
		FontFile:        "./rounded-l-mplus-2c-medium.ttf",
		Colors:          DefaultColors,
		Width:           2048,
		Height:          2048,
		Mask: MaskConf{"", color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 0,
		}},
	}

	flag.Parse()

	// Load config
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
		wordclouds.FontMaxSize(conf.Width/(numOfChar+5)),
		wordclouds.FontMinSize(conf.Width/(numOfChar+5)/10),
		wordclouds.Colors(colors),
		wordclouds.MaskBoxes(boxes),
		wordclouds.Height(conf.Height),
		wordclouds.Width(conf.Width),
		wordclouds.RandomPlacement(conf.RandomPlacement),
	)

	img := w.Draw()
	return img
}
