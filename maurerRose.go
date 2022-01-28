// // Maurer Rose

package main

import (
	"image/color"
	"math"

	"github.com/go-p5/p5"
)

var n float64 = 6
var d float64 = 71

var n2 float64 = 6
var d2 float64 = 71

func setup() {
	p5.Canvas(400, 400)
}

func draw() {
	p5.Background(color.Black)

	p5.Stroke(color.RGBA{
		R: 255,
		G: 0,
		B: 255,
		A: 100,
	})
	p5.StrokeWidth(1)
	var px, py float64
	flag := false

	for i := 0; i < 361; i++ {
		var k float64 = float64(i) * d
		var r = 200 * math.Sin(n*k)
		var x = r * math.Cos(float64(k))
		var y = r * math.Sin(float64(k))

		if !flag {
			flag = true
		} else {
			p5.Line(px+200, py+200, x+200, y+200)
		}
		px = x
		py = y
	}

	p5.Stroke(color.White)
	p5.StrokeWidth(0.2)
	for j := 0; j < 361; j++ {
		var k float64 = float64(j) * d2
		var r = 150 * math.Sin(n2*k)
		var x = r * math.Cos(float64(k))
		var y = r * math.Sin(float64(k))

		if !flag {
			flag = true
		} else {
			p5.Line(px+200, py+200, x+200, y+200)
			// p5.Arc(200, 200, 200, 200, px, py)
		}
		px = x
		py = y
	}
	n2 += 0.000001
	d2 += 0.000003

	n += 0.000001
	d += 0.000003
}
