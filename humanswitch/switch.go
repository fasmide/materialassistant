package humanswitch

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio"
)

type Humanswitch []rpio.Pin

func (h Humanswitch) Read() int {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	for i, v := range h {
		// Pull up and read value
		v.PullUp()
		if v.Read() == 0 {
			return i
		}
	}

	return len(h)
}
