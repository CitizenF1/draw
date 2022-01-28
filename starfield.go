package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const with, heith = float64(1024), float64(1024)

var speed = float64(200)

type Star struct {
	pixel.Vec
	Z float64
	P float64
	C color.RGBA
}

var stars [1024]*Star

func init() {
	rand.Seed(4)

	for i := 0; i < len(stars); i++ {
		stars[i] = newStar()
	}
}

func newStar() *Star {
	return &Star{
		pixel.V(random(-with, with), random(-heith, heith)),
		random(0, with), 0, Colors[rand.Intn(len(Colors))],
	}
}

func (s *Star) update(d float64) {
	s.P = s.Z
	s.Z -= d * speed

	if s.Z < 0 {
		s.X = random(-with, with)
		s.Y = random(-heith, heith)
		s.Z = with
		s.P = s.Z
	}
}

func (s *Star) draw(imd *imdraw.IMDraw) {
	p := pixel.V(
		scale(s.X/s.Z, 0, 1, 0, with),
		scale(s.Y/s.Z, 0, 1, 0, heith),
	)

	o := pixel.V(
		scale(s.X/s.P, 0, 1, 0, with),
		scale(s.Y/s.P, 0, 1, 0, heith),
	)

	r := scale(s.Z, 0, with, 3, 0)

	imd.Color = s.C

	if p.Sub(o).Len() > 6 {
		imd.Push(p, o)
		imd.Line(r)
	}

	imd.Push(p)
	imd.Circle(r, 0)
}

func run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Bounds:      pixel.R(0, 0, with, heith),
		VSync:       true,
		Undecorated: true,
	})
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	imd.Precision = 7

	imd.SetMatrix(pixel.IM.Moved(win.Bounds().Center()))

	last := time.Now()

	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))

		if win.Pressed(pixelgl.KeyUp) {
			speed += 10
		}

		if win.Pressed(pixelgl.KeyDown) {
			if speed > 10 {
				speed -= 10
			}
		}

		if win.Pressed(pixelgl.KeySpace) {
			speed = 100
		}

		d := time.Since(last).Seconds()

		last = time.Now()

		imd.Clear()

		for _, s := range stars {
			s.update(d)
			s.draw(imd)
		}

		win.Clear(color.Black)
		imd.Draw(win)
		win.Update()
	}
}

//////////MAIN
func main() {
	pixelgl.Run(run)
}

func random(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func scale(unscaledNum, min, max, minAllowed, maxAllowed float64) float64 {
	return (maxAllowed-minAllowed)*(unscaledNum-min)/(max-min) + minAllowed
}

var Colors = []color.RGBA{
	{255, 255, 255, 255},
}

// package main

// import (
// 	"fmt"
// 	"math"
// 	"time"
// )

// const (
// 	delay      = 16 * time.Millisecond // 60 fps hahaha
// 	coreString = ".,-~:;=!*#$@"
// )

// type SliceType interface {
// 	len() int
// }

// func floatMemset(arr []float64, v float64) {
// 	for i := range arr {
// 		arr[i] = v
// 	}
// }

// func byteMemset(arr []string, v string) {
// 	for i := range arr {
// 		arr[i] = v
// 	}
// }

// func main() {
// 	A := float64(0)
// 	B := float64(0)

// 	var i, j float64
// 	var k int

// 	z := make([]float64, 1760)
// 	b := make([]string, 1760)

// 	fmt.Print("\033[H\033[2J") // clear previous stdout

// 	for {
// 		byteMemset(b, " ")
// 		floatMemset(z, 0)

// 		for j = 0; j < 6.28; j += 0.07 {
// 			for i = 0; i < 6.28; i += 0.02 {
// 				c := math.Sin(i)
// 				d := math.Cos(j)
// 				e := math.Sin(A)
// 				f := math.Sin(j)
// 				g := math.Cos(A)
// 				h := d + 2
// 				D := 1 / (c*h*e + f*g + 5)
// 				l := math.Cos(i)
// 				m := math.Cos(B)
// 				n := math.Sin(B)
// 				t := c*h*g - f*e

// 				x := int(40 + 30*D*(l*h*m-t*n))
// 				y := int(12 + 15*D*(l*h*n+t*m))

// 				o := int(x + 80*y)

// 				N := int(8 * ((f*e-c*d*g)*m - c*d*e - f*g - l*d*n))

// 				if y < 22 && y > 0 && x > 0 && x < 80 && D > z[o] {
// 					z[o] = D

// 					// golang doesn't have ternary operator
// 					point := 0
// 					if N > 0 {
// 						point = N
// 					}

// 					b[o] = string(coreString[point])
// 				}

// 			}
// 		}

// 		print("\x1b[H")

// 		for k = 0; k < 1761; k++ {

// 			v := "\n"

// 			if k%80 > 0 {
// 				v = string(b[k])
// 			}

// 			fmt.Print(v)

// 			A += 0.00004
// 			B += 0.00002
// 		}

// 		time.Sleep(delay)

// 	}

// }
