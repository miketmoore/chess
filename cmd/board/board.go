package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	fmt.Println("Draw a chess board")

	// Setup GUI window
	cfg := pixelgl.WindowConfig{
		Title:  "Chess",
		Bounds: pixel.R(0, 0, 400, 400),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Darkgray)

	var squareW float64 = 50
	var squareH float64 = 50
	var r, c float64
	var squares [64]*imdraw.IMDraw
	var colorFlag = true
	var xInc, yInc float64
	i := 0
	for r = 0; r < 8; r++ {
		for c = 0; c < 8; c++ {
			fmt.Printf("i:%d: r:%d x c:%d - xInc:%d x yInc:%d\n", i, int(r), int(c), int(xInc), int(yInc))
			imd := imdraw.New(nil)
			if colorFlag {
				imd.Color = colornames.Black
				colorFlag = false
			} else {
				imd.Color = colornames.White
				colorFlag = true
			}
			imd.Push(pixel.V(xInc, yInc))
			imd.Push(pixel.V(squareW+xInc, squareH+yInc))
			imd.Rectangle(0)
			squares[i] = imd
			xInc += 50
			i++
		}
		colorFlag = !colorFlag
		xInc = 0
		yInc += 50
	}

	fmt.Printf("%T\n", squares[0])
	fmt.Printf("%T\n", squares[1])

	var f = true

	for !win.Closed() {
		win.Update()

		if f {
			fmt.Printf("Drawing...\n")
			for i := 0; i < len(squares); i++ {
				squares[i].Draw(win)
			}
			f = false
		}
	}
}

func main() {
	pixelgl.Run(run)
}
