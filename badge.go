package main

import (
	"bytes"
	"fmt"
	"github.com/sstallion/go-hid"
)

type badgeEffect uint8

const (
	Left badgeEffect = iota
	Right
	Up
	Down
	Freeze
	Animation
	Snow
	Volume
	Laser
)

type badgeMessage struct {
	Blink  bool
	Frame  bool
	Speed  uint8
	Effect badgeEffect
	Canvas canvas
}

type Badge struct {
	Brightness uint8
	Messages   [nMaxMsges]badgeMessage
}

func (b *Badge) setMessage(i uint8, message badgeMessage) {
	b.Messages[i] = message
}

func (b *Badge) generateHeader() badgeHeader {
	header := defaultHeader()
	header.Brightness = b.Brightness

	for i, message := range b.Messages {
		if message.Canvas.len() == 0 {
			continue
		}
		header.Flash |= boolToUint8(message.Blink) << uint8(i)
		header.Border |= boolToUint8(message.Frame) << uint8(i)
		header.LineConf[i] = ((message.Speed - 1) << 4) | (uint8(message.Effect) & 0x0F)
		header.MsgLen[i] = message.Canvas.len()
	}

	return header
}

func boolToUint8(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

func (b *Badge) send(device *hid.Device) {
	buf := new(bytes.Buffer)

	header := b.generateHeader()
	err := header.toBytes(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, message := range b.Messages {
		buf.Write(message.Canvas.Data[:])
	}

	padding := buf.Len() % 64
	buf.Grow(padding)

	if buf.Len() > 8192 {
		fmt.Println("too many bytes")
		return
	}

	fmt.Printf("%v\n", buf.Bytes())
	_, err = device.Write(buf.Bytes())
	if err != nil {
		fmt.Printf("a: %s\n", err)
		return
	}
}
