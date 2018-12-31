package view

import (
	"image"
	"os"

	"github.com/faiface/pixel"
)

var spriteSheetPath = "assets/standard_chess_pieces_sprite_sheet.png"

// SpriteByName represents a map of piece names to sprites
type SpriteByName map[string]*pixel.Sprite

// PieceSpriteSet contains one sprite per type of piece
type PieceSpriteSet struct {
	King   *pixel.Sprite
	Queen  *pixel.Sprite
	Bishop *pixel.Sprite
	Knight *pixel.Sprite
	Rook   *pixel.Sprite
	Pawn   *pixel.Sprite
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
