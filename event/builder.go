package event

import (
	"fmt"

	"github.com/holoplot/go-evdev"
)

func NewBuilder(ev *evdev.InputDevice) (chan string, error) {
	err := ev.Grab()
	if err != nil {
		return nil, fmt.Errorf("unable to grab input device: %w", err)
	}
	ch := make(chan string)
	go func() {
		defer ev.Ungrab()

		build := ""
		for {
			e, err := ev.ReadOne()
			if err != nil {
				panic(err)
			}
			if e.Type != evdev.EV_KEY {
				continue
			}
			if e.Value != 1 {
				continue
			}

			if e.Code == evdev.KEY_0 {
				build += "0"
			}
			if e.Code == evdev.KEY_1 {
				build += "1"
			}
			if e.Code == evdev.KEY_2 {
				build += "2"
			}
			if e.Code == evdev.KEY_3 {
				build += "3"
			}
			if e.Code == evdev.KEY_4 {
				build += "4"
			}
			if e.Code == evdev.KEY_5 {
				build += "5"
			}
			if e.Code == evdev.KEY_6 {
				build += "6"
			}
			if e.Code == evdev.KEY_7 {
				build += "7"
			}
			if e.Code == evdev.KEY_8 {
				build += "8"
			}
			if e.Code == evdev.KEY_9 {
				build += "9"
			}
			if e.Code == evdev.KEY_A {
				build += "A"
			}
			if e.Code == evdev.KEY_B {
				build += "B"
			}
			if e.Code == evdev.KEY_C {
				build += "C"
			}
			if e.Code == evdev.KEY_D {
				build += "D"
			}
			if e.Code == evdev.KEY_E {
				build += "E"
			}
			if e.Code == evdev.KEY_F {
				build += "F"
			}

			if e.Code == evdev.KEY_ENTER {
				ch <- build
				build = ""
			}
		}
	}()
	return ch, nil
}
