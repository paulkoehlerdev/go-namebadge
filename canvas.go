package main

type canvas struct {
	Data   []uint8
	Length uint16
}

func newCanvas(length uint16) canvas {
	return canvas{
		Data:   make([]uint8, length*nPixelsHeight),
		Length: length,
	}
}

func (c *canvas) len() uint16 {
	return c.Length
}

func (c *canvas) size() (x, y int) {
	return int(c.Length * 8), nPixelsHeight
}

func (c *canvas) setPixel(x, y int) {
	data_x := x / 8
	data_y := y

	data_offset := x % 8
	data_index := data_x*nPixelsHeight + data_y

	c.Data[data_index] |= 0x80 >> data_offset
}

func (c *canvas) setLetter(i int, bitmap []uint8) {
	copy(c.Data[nPixelsHeight*i:nPixelsHeight*i+nPixelsHeight], bitmap)
}
