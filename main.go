package main

import (
	"time"

	"github.com/fasmide/materialassistant/label"
)

func main() {
	maker, err := label.NewMaker(89, 36)
	if err != nil {
		panic(err)
	}

	_, err = maker.MaterialSVG("Falke Carlsen", time.Hour*24*365*75-(time.Hour*24*8), maker.TagUseAllowed())
	//err = maker.DebugSVG()
	if err != nil {
		panic(err)
	}
}
