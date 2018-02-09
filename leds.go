package main

import (
	"fmt"
)

func leds(ch <-chan ControlData) {
	for {
		fmt.Println(<-ch)
	}
}