package main

import (
	"image/color"

	"github.com/go-p5/p5"
)

var a float64

func main() {
	var x int

	p5.Run(setup, draw)

	_ = x
}

func setup() {
	p5.Canvas(400, 200)
	p5.Background(color.Gray{Y: 220})
}

func draw() {

	p5.Fill(color.RGBA{R: 255, G: 108, B: 105, A: 255})
	p5.Rect(0, 0, 400, 200)

	p5.Translate(200, 100)
	a = a + 0.01
	p5.Rotate(a)
	p5.Rect(30, 10, 38, 58)
	p5.Triangle(-30, -20, -20, 30, 40, -10)
	p5.Circle(0, 0, 10)
}
