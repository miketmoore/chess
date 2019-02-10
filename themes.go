package chess

import "image/color"

// Themes is a collection of color themes for the board
var Themes = map[string]map[string]color.RGBA{
	"classic": map[string]color.RGBA{
		"black": color.RGBA{0, 0, 0, 255},
		"white": color.RGBA{255, 255, 255, 255},
	},
	"coral": map[string]color.RGBA{
		"black": color.RGBA{112, 162, 163, 255},
		"white": color.RGBA{177, 228, 185, 255},
	},
	"emerald": map[string]color.RGBA{
		"black": color.RGBA{111, 143, 114, 255},
		"white": color.RGBA{173, 189, 143, 255},
	},
	"sandcastle": map[string]color.RGBA{
		"black": color.RGBA{184, 139, 74, 255},
		"white": color.RGBA{227, 193, 111, 255},
	},
}
