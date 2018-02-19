package pieces

import (
	"image"
	"os"

	"github.com/faiface/pixel"
)

var spriteSheetPath = "assets/chess-pieces.png"

type PieceMap map[string]*pixel.Sprite
type PiecesMap map[string]PieceMap

func Build() PiecesMap {
	// Load sprite sheet graphic
	pic, err := loadPicture(spriteSheetPath)
	if err != nil {
		panic(err)
	}
	return makePieces(pic)
}
func makePieces(pic pixel.Picture) PiecesMap {
	var xInc float64 = 62
	var yInc float64 = 60
	return PiecesMap{
		"white": PieceMap{
			"king":   newSprite(pic, 0, 0, xInc, yInc),
			"queen":  newSprite(pic, xInc, 0, xInc*2, yInc),
			"rook":   newSprite(pic, xInc*2, 0, xInc*3, yInc),
			"knight": newSprite(pic, xInc*3, 0, xInc*4, yInc),
			"bishop": newSprite(pic, xInc*4, 0, xInc*5+5, yInc),
			"pawn":   newSprite(pic, xInc*5+5, 0, xInc*6, yInc),
		},
		"black": PieceMap{
			"king":   newSprite(pic, 0, yInc, xInc, yInc*2),
			"queen":  newSprite(pic, xInc, yInc, xInc*2, yInc*2+5),
			"rook":   newSprite(pic, xInc*2, yInc, xInc*3, yInc*2),
			"knight": newSprite(pic, xInc*3, yInc, xInc*4, yInc*3),
			"bishop": newSprite(pic, xInc*4, yInc, xInc*5+5, yInc*4),
			"pawn":   newSprite(pic, xInc*5+5, yInc, xInc*6, yInc*5),
		},
	}
}

func newSprite(pic pixel.Picture, xa, ya, xb, yb float64) *pixel.Sprite {
	return pixel.NewSprite(pic, pixel.Rect{pixel.Vec{xa, ya}, pixel.Vec{xb, yb}})
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
