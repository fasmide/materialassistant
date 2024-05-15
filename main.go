package main

import (
	"github.com/fasmide/materialassistant/label"
)

func main() {
	maker := label.Maker{
		MmHeight: 413,
		MmWidth:  991,
	}
	_, err := maker.MaterialSVG("midemide")
	if err != nil {
		panic(err)
	}
}
