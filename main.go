package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fasmide/materialassistant/acs"
	"github.com/fasmide/materialassistant/event"
	"github.com/fasmide/materialassistant/label"
	"github.com/holoplot/go-evdev"
)

func main() {
	d, err := evdev.OpenByName("Sycreader RFID Technology Co., Ltd SYC ID&IC USB Reader")
	if err != nil {
		panic(err)
	}

	cardIds, err := event.NewBuilder(d)
	if err != nil {
		panic(err)
	}

	api := acs.API{Endpoint: os.Getenv("MA_ACS_ENDPOINT"), APIToken: os.Getenv("MA_ACS_TOKEN")}
	maker, err := label.NewMaker(89, 36)
	if err != nil {
		panic(err)
	}

	for {
		card := <-cardIds
		fmt.Printf("Card swept: %s\n", card)

		id, err := api.Lookup(card)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Found member: %+v\n", id)

		_, err = maker.MaterialSVG(id, time.Hour*24*365*7, maker.TagUseAllowed())
		//err = maker.DebugSVG()
		if err != nil {
			panic(err)
		}
	}
}
