package figs

import (
	"codingtrig/trig"

	"github.com/bit101/blgo"
)

// Fig02020 blah
func Fig02020() {
	filename := "images/figure_2-20.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)

	trig.DrawRightTriangle(500, 800, 1000, 700, surface)

	surface.WriteToPNG(filename)
}
