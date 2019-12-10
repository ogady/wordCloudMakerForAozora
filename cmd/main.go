package main

import (
	"flag"
	"log"
	"os"

	"github.com/ogady/wordCloudMakerForAozora/internal"
)

func main() {

	var (
		output         = flag.String("o", "output.png", "path to output image")
		titleName      = flag.String("t", "銀河鉄道の夜", "target TitleName")
		specifiedColor = flag.String("c", "red", "specify the color to draw from ’red’, ’blue’, ’green’, and ’vivid’.")
	)

	flag.Parse()

	repo := internal.NewWordCloudCreater(*output, *titleName, *specifiedColor)

	err := repo.Execute()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
