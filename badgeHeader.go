package main

import (
	"bytes"
	"encoding/binary"
)

type badgeHeader struct {
	Start      [5]uint8
	Brightness uint8
	Flash      uint8
	Border     uint8
	LineConf   [nMaxMsges]uint8
	MsgLen     [nMaxMsges]uint16
}

func (b badgeHeader) toBytes(buf *bytes.Buffer) error {
	// Write Start
	if err := binary.Write(buf, binary.LittleEndian, b.Start); err != nil {
		return err
	}

	// Write Brightness, Flash, and Border
	if err := binary.Write(buf, binary.LittleEndian, b.Brightness); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, b.Flash); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, b.Border); err != nil {
		return err
	}

	// Write LineConf
	if err := binary.Write(buf, binary.LittleEndian, b.LineConf); err != nil {
		return err
	}

	// Write MsgLen
	if err := binary.Write(buf, binary.BigEndian, b.MsgLen); err != nil {
		return err
	}

	pad := buf.Len() % 64
	buf.Write(make([]byte, pad))

	return nil
	// var packet [reportBufSize]byte
	// copy(packet[:], buf.Bytes())
	// return packet[:], nil
}

func defaultHeader() badgeHeader {
	return badgeHeader{
		Start:      [5]uint8{0x77, 0x61, 0x6e, 0x67, 0x00},
		Brightness: 0xFF, // full brightness for now
		Flash:      0,
		Border:     0,
		LineConf:   [nMaxMsges]uint8{},
		MsgLen:     [nMaxMsges]uint16{},
	}
}
