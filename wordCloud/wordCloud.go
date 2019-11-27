package wordCloud

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"

	"github.com/psykhi/wordclouds"
)

func CreateWordCloud(wordList map[string]int, config *string) image.Image {

	var DefaultColors = []color.RGBA{
		{0x1b, 0x1b, 0x1b, 0xff},
		{0x48, 0x48, 0x4B, 0xff},
		{0x59, 0x3a, 0xee, 0xff},
		{0x65, 0xCD, 0xFA, 0xff},
		{0x70, 0xD6, 0xBF, 0xff},
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
		FontMaxSize:     700,
		FontMinSize:     20,
		RandomPlacement: false,
		FontFile:        "./rounded-l-mplus-2c-medium.ttf",
		Colors:          DefaultColors,
		Width:           3072,
		Height:          3072,
		Mask: MaskConf{"./mask.png", color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 0,
		}},
	}

	flag.Parse()

	// Load config
	conf := DefaultConf
	f, err := os.Open(*config)
	if err == nil {
		defer f.Close()
		reader := bufio.NewReader(f)
		dec := json.NewDecoder(reader)
		err = dec.Decode(&conf)
		if err != nil {
			fmt.Printf("Failed to decode config, using defaults instead: %s\n", err)
		}
	} else {
		fmt.Println("No config file. Using defaults")
	}

	os.Chdir(filepath.Dir(*config))

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
		wordclouds.FontMaxSize(conf.FontMaxSize),
		wordclouds.FontMinSize(conf.FontMinSize),
		wordclouds.Colors(colors),
		wordclouds.MaskBoxes(boxes),
		wordclouds.Height(conf.Height),
		wordclouds.Width(conf.Width),
		wordclouds.RandomPlacement(conf.RandomPlacement))

	img := w.Draw()
	return img
}
