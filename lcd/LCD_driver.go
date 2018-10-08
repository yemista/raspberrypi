/*
# Copyright (c) 2014 Adafruit Industries
# Author: Tony DiCola
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.
*/

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"strings"
)

// Commands
const LCD_CLEARDISPLAY        = 0x01
const LCD_RETURNHOME          = 0x02
const LCD_ENTRYMODESET        = 0x04
const LCD_DISPLAYCONTROL      = 0x08
const LCD_CURSORSHIFT         = 0x10
const LCD_DISPLAYCONTROL      = 0x08
const LCD_SETCGRAMADDR        = 0x40
const LCD_SETDDRAMADDR        = 0x80

// Entry flags
const LCD_ENTRYRIGHT          = 0x00
const LCD_ENTRYLEFT           = 0x02
const LCD_ENTRYSHIFTINCREMENT = 0x01
const LCD_ENTRYSHIFTDECREMENT = 0x00

// Control flags
const LCD_DISPLAYON           = 0x04
const LCD_DISPLAYOFF          = 0x00
const LCD_CURSORON            = 0x02
const LCD_CURSOROFF           = 0x00
const LCD_BLINKON             = 0x01
const LCD_BLINKOFF            = 0x00

// Move flags
const LCD_DISPLAYMOVE         = 0x08
const LCD_CURSORMOVE          = 0x00
const LCD_MOVERIGHT           = 0x04
const LCD_MOVELEFT            = 0x00

// Function set flags
const LCD_8BITMODE            = 0x10
const LCD_4BITMODE            = 0x00
const LCD_2LINE               = 0x08
const LCD_1LINE               = 0x00
const LCD_5x10DOTS            = 0x04
const LCD_5x8DOTS             = 0x00

// Offset for up to 4 rows.
const LCD_ROW_OFFSETS         = (0x00, 0x40, 0x14, 0x54)

// Char LCD plate button names.
const SELECT                  = 0
const RIGHT                   = 1
const DOWN                    = 2
const UP                      = 3
const LEFT                    = 4

// Char LCD backpack GPIO numbers.
const LCD_RS         = 1
const LCD_EN         = 2
const LCD_D4         = 3
const LCD_D5         = 4
const LCD_D6         = 5
const LCD_D7         = 6

type char uint8

var (
	RS_PIN = rpio.Pin(LCD_RS)
	CLOCK_PIN = rpio.Pin(LCD_EN)
	D0_PIN = rpio.Pin(LCD_D4)
	D1_PIN = rpio.Pin(LCD_D5)
	D2_PIN = rpio.Pin(LCD_D6)
	D3_PIN = rpio.Pin(LCD_D7)

	data_pins := [4]Pin{
		D0_PIN,
		D1_PIN,
		D2_PIN,
		D3_PIN
	}
)

func init() {
	RS_PIN.Output()
	CLOCK_PIN.Output()

	for i, v := range data_pins {
		v.Output()
	}
}

func clockPulse() {
	CLOCK_PIN.Low()
	CLOCK_PIN.High()
	CLOCK_PIN.Low()
}

func printC(in char) {

	for i, v := range data_pins {
		val := ((in >> i + 4) & 1) > 0

		if val {
			db_pins[i].High()
		} else {
			db_pins[i].Low()
		}
	}

	clockPulse()

	for i, v := range data_pins {
		val := ((in >> i) & 1) > 0

		if val {
			db_pins[i].High()
		} else {
			db_pins[i].Low()
		}
	}

	clockPulse()

}

func setCharMode() {
	RS_PIN.High()
}

func setCmdMode() {
	RS_PIN.Low()
}

func SetCursorBlink() {
	setCmdMode()
	val := LCD_DISPLAYCONTROL | LCD_DISPLAYON | LCD_CURSORON
	printC(val)
}

func Write(in string)  {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()
}
