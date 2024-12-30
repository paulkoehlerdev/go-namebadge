package main

import (
	"github.com/sstallion/go-hid"
)

const (
	vendorIDS1144  = 0x0416
	productIDS1144 = 0x5020
)

const (
	nMaxMsges     = 8
	nPixelsHeight = 11
)

func staticTextMessage(text string) badgeMessage {
	msg := badgeMessage{
		Blink:  false,
		Frame:  false,
		Speed:  0,
		Effect: Freeze,
		Canvas: newCanvas(uint16(len(text))),
	}
	for i := 0; i < len(text); i++ {
		msg.Canvas.setLetter(i, letterFromRune(rune(text[i])))
	}

	return msg
}

func main() {
	dev, err := hid.OpenFirst(vendorIDS1144, productIDS1144)
	if err != nil {
		panic(err)
	}

	b := Badge{
		Brightness: 0,
	}

	b.setMessage(uint8(0), staticTextMessage("38c3"))
	b.setMessage(uint8(1), staticTextMessage("Paul"))
	b.setMessage(uint8(2), staticTextMessage("Marcel"))

	b.send(dev)
}
