package main

import (
	"fmt"
	"os"
	"os/exec"
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

		imgReader, err := maker.MaterialSVG(id, time.Hour*24*365*7, maker.TagUseAllowed())
		if err != nil {
			panic(err)
		}

		c := exec.Command("lprint")
		c.Stdin = imgReader

		out, err := c.Output()
		if err != nil {
			fmt.Printf("lprint said:\n%s", string(out))
			panic(err)
		}
	}
}
