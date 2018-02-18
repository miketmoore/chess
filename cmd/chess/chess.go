package main

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// var spriteSheetPath = "assets/spritesheet.png"
var spriteSheetPath = "assets/chess-pieces.png"

func run() {
	// Load sprite sheet graphic
	pic, err := loadPicture(spriteSheetPath)
	if err != nil {
		panic(err)
	}

	pieces := makePieces(pic)

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
	for !win.Closed() {
		win.Update()
		mat := pixel.IM
		mat = mat.Moved(win.Bounds().Center())
		// pieces["whiteKing"].Draw(win, mat)
		// pieces["whiteQueen"].Draw(win, mat)
		// pieces["whiteRook"].Draw(win, mat)
		// pieces["whiteKnight"].Draw(win, mat)
		// pieces["whiteBishop"].Draw(win, mat)
		// pieces["whitePawn"].Draw(win, mat)
		// pieces["blackKing"].Draw(win, mat)
		// pieces["blackQueen"].Draw(win, mat)
		// pieces["blackRook"].Draw(win, mat)
		// pieces["blackKnight"].Draw(win, mat)
		// pieces["blackBishop"].Draw(win, mat)
		pieces["blackPawn"].Draw(win, mat)
	}
}

func main() {
	pixelgl.Run(run)
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func makePieces(pic pixel.Picture) map[string]*pixel.Sprite {
	var xInc float64 = 62
	var yInc float64 = 60
	return map[string]*pixel.Sprite{
		"whiteKing":   newSprite(pic, 0, 0, xInc, yInc),
		"whiteQueen":  newSprite(pic, xInc, 0, xInc*2, yInc),
		"whiteRook":   newSprite(pic, xInc*2, 0, xInc*3, yInc),
		"whiteKnight": newSprite(pic, xInc*3, 0, xInc*4, yInc),
		"whiteBishop": newSprite(pic, xInc*4, 0, xInc*5+5, yInc),
		"whitePawn":   newSprite(pic, xInc*5+5, 0, xInc*6, yInc),
		"blackKing":   newSprite(pic, 0, yInc, xInc, yInc*2),
		"blackQueen":  newSprite(pic, xInc, yInc, xInc*2, yInc*2+5),
		"blackRook":   newSprite(pic, xInc*2, yInc, xInc*3, yInc*2),
		"blackKnight": newSprite(pic, xInc*3, yInc, xInc*4, yInc*3),
		"blackBishop": newSprite(pic, xInc*4, yInc, xInc*5+5, yInc*4),
		"blackPawn":   newSprite(pic, xInc*5+5, yInc, xInc*6, yInc*5),
	}
}

func newSprite(pic pixel.Picture, xa, ya, xb, yb float64) *pixel.Sprite {
	return pixel.NewSprite(pic, pixel.Rect{pixel.Vec{xa, ya}, pixel.Vec{xb, yb}})
}
