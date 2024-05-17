package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fasmide/materialassistant/acs"
	"github.com/fasmide/materialassistant/label"
)

func main() {
	api := acs.API{Endpoint: os.Getenv("MA_ACS_ENDPOINT"), APIToken: os.Getenv("MA_ACS_TOKEN")}
	id, err := api.Lookup(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found: %+v\n", id)

	maker, err := label.NewMaker(89, 36)
	if err != nil {
		panic(err)
	}

	_, err = maker.MaterialSVG(id, time.Hour*24*365*7, maker.TagUseAllowed())
	//err = maker.DebugSVG()
	if err != nil {
		panic(err)
	}
}
