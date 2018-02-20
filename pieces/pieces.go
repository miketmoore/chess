package pieces

import (
	"image"
	"os"

	"github.com/faiface/pixel"
)

var spriteSheetPath = "assets/standard_chess_pieces_sprite_sheet.png"

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
	return PiecesMap{
		"black": PieceMap{
			"king":   newSprite(pic, 0, 0, 40, 40),
			"queen":  newSprite(pic, 40, 0, 90, 40),
			"bishop": newSprite(pic, 90, 0, 140, 40),
			"knight": newSprite(pic, 130, 0, 180, 40),
			"rook":   newSprite(pic, 185, 0, 220, 40),
			"pawn":   newSprite(pic, 230, 0, 270, 40),
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
