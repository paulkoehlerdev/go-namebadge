package main

import (
	"bytes"
	"math/rand"
)

var rocket = []byte{
	// ........ ........
	// ........ ........
	// ........ ........
	// ........ ...xx...
	// ..xxxxxx xxx..x..
	// ..x..... ......x.
	// ..x..... .....x..
	// ..xxxxxx xxx.x...
	// ........ ..xx....
	// ........ ........
	// ........ ........
	0x0, 0x0, 0x0, 0x0, 0b00111111, 0b00100000, 0b00111111, 0x0, 0x0, 0x0, 0x0,
	0x0, 0x0, 0x0, 0b00011000, 0b11101100, 0b110, 0b11101100, 0b00011000, 0x0, 0x0, 0x0,
}

var triangle = []byte{
	0xF0, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF0,
	0x00, 0xF0, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF0, 0x00,
	0x00, 0x00, 0xF0, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF0, 0x00, 0x00,
}

func wholeRocket() badgeMessage {
	buf := new(bytes.Buffer)
	buf.Grow(nPixelsHeight * 10)

	behind := make([]byte, nPixelsHeight*5)
	rand.Read(behind)
	buf.Write(behind)

	triangleRand := make([]byte, len(triangle))
	rand.Read(triangleRand)

	for i := range triangle {
		triangleRand[i] &= triangle[i]
	}

	buf.Write(triangleRand)

	buf.Write(rocket)

	return badgeMessage{
		Blink:  false,
		Frame:  false,
		Speed:  5,
		Effect: Right,
		Canvas: canvas{Data: buf.Bytes(), Length: 10},
	}
}

func explosion() badgeMessage {
	buf := new(bytes.Buffer)

	nFrames := 10
	buf.Grow(nPixelsHeight * 6 * nFrames)

	for i := 0; i < nFrames; i++ {
		frame := make([]byte, nPixelsHeight*6)
		rand.Read(frame)
		buf.Write(frame)
	}

	return badgeMessage{
		Blink:  false,
		Frame:  false,
		Speed:  4,
		Effect: Animation,
		Canvas: canvas{Data: buf.Bytes(), Length: uint16(6 * nFrames)},
	}
}
