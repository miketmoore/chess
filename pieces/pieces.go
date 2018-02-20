package pieces

import (
	"image"
	"os"

	"github.com/faiface/pixel"
)

var spriteSheetPath = "assets/standard_chess_pieces_sprite_sheet.png"

// ByName represents a map of piece names to sprites
type ByName map[string]*pixel.Sprite

// ByColor represents a map of color name to ByName
type ByColor map[string]ByName

// New constructs a ByColor (chess piece sprites)
func New() ByColor {
	// Load sprite sheet graphic
	pic, err := loadPicture(spriteSheetPath)
	if err != nil {
		panic(err)
	}
	return makePieces(pic)
}

func makePieces(pic pixel.Picture) ByColor {
	return ByColor{
		"black": ByName{
			"king":   newSprite(pic, 0, 0, 40, 40),
			"queen":  newSprite(pic, 40, 0, 90, 40),
			"bishop": newSprite(pic, 90, 0, 140, 40),
			"knight": newSprite(pic, 130, 0, 180, 40),
			"rook":   newSprite(pic, 185, 0, 220, 40),
			"pawn":   newSprite(pic, 230, 0, 270, 40),
		},
		"white": ByName{
			"king":   newSprite(pic, 0, 40, 40, 85),
			"queen":  newSprite(pic, 40, 40, 90, 85),
			"bishop": newSprite(pic, 90, 40, 140, 85),
			"knight": newSprite(pic, 130, 40, 185, 85),
			"rook":   newSprite(pic, 185, 40, 220, 85),
			"pawn":   newSprite(pic, 230, 40, 270, 85),
		},
	}
}

func newSprite(pic pixel.Picture, xa, ya, xb, yb float64) *pixel.Sprite {
	return pixel.NewSprite(pic, pixel.Rect{Min: pixel.V(xa, ya), Max: pixel.V(xb, yb)})
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
