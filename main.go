package main

import (
	"time"

	"github.com/fasmide/materialassistant/label"
)

func main() {
	maker, err := label.NewMaker(1191, 413)
	if err != nil {
		panic(err)
	}

	_, err = maker.MaterialSVG("KristianKristianKristianKristian", time.Hour*24*(365+(365/2)), maker.TagSkab())
	if err != nil {
		panic(err)
	}
}
