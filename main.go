package main

import (
	_ "embed"
	"github.com/sstallion/go-hid"
)

const (
	vendorIDS1144  = 0x0416
	productIDS1144 = 0x5020
)

const (
	nMaxMsges     = 8
	nPixelsHeight = 11
	nPixelsWidth  = 44
)

func staticTextCanvas(text string) canvas {
	c := newCanvas(uint16(len(text)))
	for i := 0; i < len(text); i++ {
		c.setLetter(i, letterFromRune(rune(text[i])))
	}

	return c
}

func staticTextMessage(text string) badgeMessage {
	return badgeMessage{
		Blink:  false,
		Frame:  false,
		Speed:  0,
		Effect: Freeze,
		Canvas: staticTextCanvas(text),
	}
}

func main() {

	dev, err := hid.OpenFirst(vendorIDS1144, productIDS1144)
	if err != nil {
		panic(err)
	}

	b := Badge{
		Brightness: 0,
	}

	b.setMessage(uint8(0), badgeMessage{
		Blink:  false,
		Frame:  false,
		Speed:  8,
		Effect: Left,
		Canvas: staticTextCanvas("Happy New Year!"),
	})

	b.setMessage(uint8(1), badgeMessage{
		Blink:  true,
		Frame:  false,
		Speed:  3,
		Effect: Freeze,
		Canvas: staticTextCanvas("2025"),
	})

	b.setMessage(uint8(2), wholeRocket())
	b.setMessage(uint8(3), explosion())

	b.setMessage(uint8(4), badgeMessage{
		Blink:  false,
		Frame:  false,
		Speed:  1,
		Effect: Snow,
		Canvas: staticTextCanvas("2025"),
	})

	b.send(dev)
}
