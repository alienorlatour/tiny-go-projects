package main

import (
	"time"

	"machine"
)

func main() {
	car := newCarLight(machine.D2, machine.D3, machine.D4)
	walk := newWalkLight(machine.D5, machine.D6)

	for {
		walk.Stop()
		car.Go()
		time.Sleep(time.Second * 5)

		car.Stop()
		walk.Go()
		time.Sleep(time.Second * 5)
	}
}

type carLight struct {
	red, yellow, green machine.Pin
}

func newCarLight(redPin, yellowPin, greenPin machine.Pin) *carLight {
	c := &carLight{
		red:    redPin,
		yellow: yellowPin,
		green:  greenPin,
	}
	c.red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	c.yellow.Configure(machine.PinConfig{Mode: machine.PinOutput})
	c.green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	c.red.High()
	c.yellow.Low()
	c.green.Low()

	return c
}

func (c *carLight) Stop() {
	c.green.Low()
	c.yellow.High()
	time.Sleep(time.Second)
	c.yellow.Low()
	c.red.High()
}

func (c *carLight) Go() {
	c.red.High()
	c.yellow.High()
	time.Sleep(time.Second)
	c.red.Low()
	c.yellow.Low()
	c.green.High()
}

type walkLight struct {
	red, green machine.Pin
}

func newWalkLight(redPin, greenPin machine.Pin) *walkLight {
	w := &walkLight{
		red:   redPin,
		green: greenPin,
	}

	w.red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	w.green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	w.red.High()
	w.green.Low()

	return w
}

func (w *walkLight) Stop() {
	for i := 0; i < 5; i++ {
		w.green.Low()
		time.Sleep(time.Millisecond * 300)
		w.green.High()
		time.Sleep(time.Millisecond * 300)
	}
	w.green.Low()
	w.red.High()
}

func (w *walkLight) Go() {
	w.red.Low()
	w.green.High()
}
