package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/miketmoore/go-chess/board"
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

	squares := board.Build(50, colornames.Darkgray, colornames.Antiquewhite)

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
