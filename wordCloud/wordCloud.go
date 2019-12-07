package wordCloud

import (
	"flag"
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
