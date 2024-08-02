package main

import (
	"time"

	"machine"
)

func main() {
	car := newCarLight(machine.D2, machine.D3, machine.D4)
	walk := newWalkLight(machine.D5, machine.D6)
	button := machine.D7
	c := newCrossing(car, walk, button)

	go c.listenButton()

	for {
		c.Switch()
		select {
		case <-c.buttonPressed:
		case <-time.After(time.Second * 5):
		}
	}
}

type crossing struct {
	cars          *carLight
	walks         *walkLight
	button        *machine.Pin
	pedestriansGo bool
	buttonPressed chan struct{}
}

func newCrossing(cars *carLight, walks *walkLight, button machine.Pin) *crossing {
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	return &crossing{
		cars:          cars,
		walks:         walks,
		button:        &button,
		buttonPressed: make(chan struct{}, 1),
	}
}

func (c *crossing) Switch() {
	if c.pedestriansGo {
		c.pedestriansGo = false
		c.walks.Stop()
		c.cars.Go()
	} else {
		c.pedestriansGo = true
		c.cars.Stop()
		c.walks.Go()
	}
}

func (c *crossing) listenButton() {
	for {
		if !c.pedestriansGo && !c.button.Get() {
			c.buttonPressed <- struct{}{}
		}
		time.Sleep(time.Millisecond * 100)
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
