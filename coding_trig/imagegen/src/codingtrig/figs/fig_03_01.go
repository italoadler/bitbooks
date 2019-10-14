package figs

import (
	"codingtrig/trig"

	"github.com/bit101/blgo"
)

// Fig0301 blah
func Fig0301() {
	filename := "images/figure_3-1.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)

	trig.DrawRightTriangle(500, 800, 1000, 700, surface)

	surface.FillText("A", 480, 80)
	surface.FillText("B", 1520, 820)
	surface.FillText("C", 430, 820)

	surface.FillText("a", 900, 850)
	surface.FillText("b", 440, 450)
	surface.FillText("c", 960, 400)
	surface.WriteToPNG(filename)
}
