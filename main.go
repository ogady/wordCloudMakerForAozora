package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/bluele/mecab-golang"
	"github.com/psykhi/wordclouds"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {

	url := "https://www.aozora.gr.jp/cards/000081/files/456_15050.html"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	// ドキュメントから
	selection := doc.Find("body > div.main_text")
	text := selection.Text()

	encoded, err := Decode("ShiftJIS", []byte(text))

	if err != nil {
		fmt.Println(err)
	}

	m, err := mecab.New("-Owakati")
	if err != nil {
		fmt.Println(err)
	}
	defer m.Destroy()
	persedText := parseToNode(string(encoded), m)

	var config = flag.String("config", "config.json", "path to config file")
	var output = flag.String("output", "output2.png", "path to output image")
	var cpuprofile = flag.String("cpuprofile", "profile", "write cpu profile to file")

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
		Width:           4096,
		Height:          4096,
		Mask: MaskConf{"", color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 0,
		}},
	}

	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

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

	confJson, _ := json.Marshal(conf)
	fmt.Printf("Configuration: %s\n", confJson)
	err = json.Unmarshal(confJson, &conf)
	if err != nil {
		fmt.Println(err)
	}

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

	start := time.Now()
	w := wordclouds.NewWordcloud(persedText,
		wordclouds.FontFile(conf.FontFile),
		wordclouds.FontMaxSize(conf.FontMaxSize),
		wordclouds.FontMinSize(conf.FontMinSize),
		wordclouds.Colors(colors),
		wordclouds.MaskBoxes(boxes),
		wordclouds.Height(conf.Height),
		wordclouds.Width(conf.Width),
		wordclouds.RandomPlacement(conf.RandomPlacement))

	img := w.Draw()
	outputFile, err := os.Create(*output)
	if err != nil {
		// Handle error
	}

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, img)

	// Don't forget to close files
	outputFile.Close()
	fmt.Printf("Done in %v\n", time.Since(start))
}

func Decode(encname string, b []byte) ([]byte, error) {
	enc, err := enc(encname)
	if err != nil {
		return nil, err
	}
	r := bytes.NewBuffer(b)
	decoded, err := ioutil.ReadAll(transform.NewReader(r, enc.NewDecoder()))
	return decoded, err
}

func enc(encname string) (enc encoding.Encoding, err error) {
	switch encname {
	case "ShiftJIS":
		enc = japanese.ShiftJIS
	case "EUCJP":
		enc = japanese.EUCJP
	case "ISO2022JP":
		enc = japanese.ISO2022JP
	default:
		err = fmt.Errorf("Unknown encname %s", encname)
	}
	return
}

func parseToNode(text string, m *mecab.MeCab) map[string]int {

	wordMap := make(map[string]int)
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice(text)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] == "名詞" {
			if !contains(ExcludeTitleIDForLineRensai, node.Surface()) {
				// mapのキーに単語を設定して、バリューに単語のカウントを設定し、キーに対してカウントしていく
				wordMap[node.Surface()]++
				// fmt.Println(fmt.Sprintf("%s %s", node.Surface(), node.Feature()))
			}
		}
		if node.Next() != nil {
			break
		}
	}
	fmt.Println(wordMap)
	return wordMap
}

var ExcludeTitleIDForLineRensai = []string{
	"で",
	"の",
	"よう",
	"ん",
	"あれ",
	"これ",
	"それ",
	"どれ",
	"ここ",
	"どこ",
	"そこ",
	"あそこ",
	"あっち",
	"そっち",
	"こっち",
	"どっち",
}

func contains(sl []string, s string) bool {

	for _, v := range sl {
		if s == v {
			return true
		}
	}
	return false
}
