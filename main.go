package main

import "github.com/go-p5/p5"

// func main() {
// 	webcam, _ := gocv.VideoCaptureDevice(0)
// 	window := gocv.NewWindow("Hello")
// 	img := gocv.NewMat()

// 	for {
// 		webcam.Read(&img)
// 		window.IMShow(img)
// 		window.WaitKey(1)
// 	}
// }

// type Star struct {
// 	x float64
// 	y float64
// 	z float64
// }

// var stars = make([]Star, 100)

// func Stars() Star {
// 	return Star{
// 		x: p5.Random(0, 600),
// 		y: p5.Random(0, 600),
// 		z: 600,
// 	}
// }

// func show() Star {
// 	// p5.Fill(color.White)
// 	p5.Ellipse(Stars().x, Stars().y, 10, 10)
// 	return Star{}
// }

func main() {
	p5.Run(setup, draw)
}

// func setup() {
// 	p5.Canvas(600, 600)
// 	// for i := 0; i < len(stars); i++ {
// 	// 	stars[i] = Stars()
// 	// }
// }

// func draw() {
// 	p5.Background(color.Black)
// 	// for i := range stars {
// 	// 	stars[i] = Stars()
// 	// 	stars[i] = show()
// 	// }
// 	for i := 0; i < len(stars); i++ {
// 		stars[i] = Stars()
// 		stars[i] = show()
// 	}
// }

/////////////////////////////////////////////////////////////////
// func main() {
// 	p5.Run(setup, draw)
// }

// var rez = 20
// var cols, rows int = 1 + width/rez, 1 + height/rez
// var width, height int = 600, 400
// var field [][]float64 = make([][]float64, cols)

// func setup() {
// 	p5.Canvas(400, 400)
// 	for i := 0; i < cols; i++ {
// 		field[i] = make([]float64, rows)
// 		for j := 0; j < rows; j++ {
// 			field[i][j] = p5.Random(0, 1)
// 		}
// 	}
// }

// func draw() {
// 	p5.Background(color.Black)
// 	for i := 0; i < cols; i++ {
// 		for j := 0; j < rows; j++ {
// 			p5.Stroke(color.Alpha16{
// 				A: uint16(field[i][j] * 255.0),
// 			})
// 			p5.Circle(float64(i*rez), float64(j*rez), 1)
// 		}
// 	}

// 	for i := 0; i < cols; i++ {
// 		for j := 0; j < rows; j++ {

// 		}
// 	}
// }
