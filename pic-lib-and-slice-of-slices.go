package main

import (
	"bufio"
	"image"
	"image/png"
	"io"
	"os"
)

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		result[i] = make([]uint8, dx)
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			result[i][j] = uint8((i + j) / 2)
		}
	}
	return result
}

func main() {
	Show(Pic)
}

func Show(f func(dx, dy int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}

func ShowImage(m image.Image) {
	file, err := os.Create("./image.png")
	checkError(err)
	w := bufio.NewWriter(file)
	defer w.Flush()

	// _, err = w.WriteString("IMAGE:")
	// checkError(err)
	png.Encode(w, m)

	io.WriteString(w, "\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
