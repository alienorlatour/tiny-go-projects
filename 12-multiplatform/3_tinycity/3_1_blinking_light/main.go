package main

import (
	"time"

	"machine"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		time.Sleep(time.Second / 2)

		led.Low()
		time.Sleep(time.Second / 2)
	}
}
