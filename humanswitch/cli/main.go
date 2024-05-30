package main

import (
	"fmt"

	"github.com/fasmide/materialassistant/humanswitch"
	"github.com/stianeikeland/go-rpio"
)

func main() {
	h := humanswitch.Humanswitch{rpio.Pin(2),
		rpio.Pin(3),
		rpio.Pin(4),
		rpio.Pin(17),
	}

	i := h.Read()
	fmt.Printf("Read: %d", i)

}
