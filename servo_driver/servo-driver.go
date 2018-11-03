package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
)

var (
	// Use mcu pin 10, corresponds to physical pin 19 on the pi
	sdaPin = rpio.Pin(2)
	sclPin = rpio.Pin(3)
)

func i2cStart() {
	sclPin.Low()
	sdaPin.High()
	sclPin.High()
	sdaPin.Low()
}

func i2cStop() {
	sclPin.Low()
	sdaPin.Low()
	sclPin.High()
	sdaPin.High()
}

func i2cDelay() {
	for i := 10000; i > 0; --i {

	}
}

func i2cAddr(addr uint8) {
	addr &= 0x7f

	for i := 7; i >= 0; --i {
		sclPin.Low()
		var bit := (addr >> i) & 0x01

		if(bit == 0) {
			sdaPin.Low()
		} else {
			sdaPin.High()
		}

		sclPin.High()
		i2cDelay()
	}

	sclPin.Low()
	sdaPin.Low()
	sclPin.High()
	i2cDelay()
	i2cAck()
}

func i2cAck() bool {
	sclPin.Low()
	sdaPin.Input()
	i2cDelay()
	sclPin.High()
	
	var success := true	

	if rpio.ReadPin(sdaPin) == High {
		success = false
	}
	
	sclPin.Low()
	sdaPin.Output()
	return success
}

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	sdaPin.Output()
	sclPin.Output()

	// Toggle pin 20 times
	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}

