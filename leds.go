package main

import (
	"encoding/binary"
	"time"
	"github.com/jgarff/rpi_ws281x/golang/ws2811"
	"fmt"
)

const (
	LED_PIN = 18
	LED_COUNT = 60
	LED_FREQ = 800000
	LED_DMA = 5
	LED_BRIGHTNESS = 255
	LED_INVERT = false
	LED_CHANNEL = 0
)

func color(red uint8, green uint8, blue uint8) uint32 {
	u := []uint8{red, green, blue, 0}
	return binary.LittleEndian.Uint32(u)
}

func colorWipe(color uint32) error {
	for i := 0; i < LED_COUNT; i++ {
		ws2811.SetLed(i, color)
		err := ws2811.Render()
		if err != nil {
			ws2811.Clear()
			return err
		}

		time.Sleep(50 * time.Millisecond)
	}

	return nil
}

func wheel(pos byte) uint32 {
	if pos < 85 {
		return color(pos * 3, 255 - pos * 3, 0)
	} else if pos < 170 {
		pos -= 85
		return color(255 - pos * 3, 0, pos * 3)
	} else {
		pos -= 170
		return color(0, pos * 3, 255 - pos * 3)
	}
}

func rainbow() {
	for j := 0; j < 256; j++ {
		for i := 0; i < LED_COUNT; i++ {
			ws2811.SetLed(i, wheel(uint8(i + j) & 255))
		}
		ws2811.Render()
		time.Sleep(20 * time.Millisecond)
	}
}

func rainbowCycle() {
	for j := 0; j < 256; j++ {
		for i := 0; i < LED_COUNT; i++ {
			ws2811.SetLed(i, wheel(uint8((i * 256 / LED_COUNT) + j) & 255))
		}
		ws2811.Render()
		time.Sleep(20 * time.Millisecond)
	}
}

func leds(ch <-chan ControlData, hasData <-chan bool) {
	defer ws2811.Fini()
	err := ws2811.Init(LED_PIN, LED_COUNT, LED_BRIGHTNESS)

	if err != nil {
		fmt.Println(err.Error())
	}

	var red uint8 = 255
	var green uint8 = 255
	var blue uint8 = 255

	var anim uint8

	ticker := time.NewTicker(20 * time.Millisecond)

	for {
		select {
		case input := <-ch:
			if input.Key != nil {
				switch input.Key {
				case "red":
					red = input.Value
				case "green":
					green = input.Value
				case "blue":
					blue = input.Value
				case "animation":
					anim = input.Value
				}
			}
		case t := <-ticker.C:
			switch anim {
			case 0:
				//Off
				colorWipe(color(0, 0, 0))
			case 1:
				//Solid Color
				colorWipe(color(red, green, blue))
			case 2:
				//Rainbow
				rainbow()
			case 3:
				//Rainbow cycle
				rainbowCycle()
			}
		}
	}
}