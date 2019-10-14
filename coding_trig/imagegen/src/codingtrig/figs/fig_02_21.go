package figs

import (
	"codingtrig/trig"

	"github.com/bit101/blgo"
)

// Fig0221 blah
func Fig0221() {
	filename := "images/figure_2-21.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)

	trig.DrawRightTriangle(500, 800, 1000, 700, surface)
	surface.FillText("leg", 880, 850)
	surface.FillText("leg", 400, 450)
	surface.FillText("hypotenuse", 1000, 420)
	surface.FillText("A", 480, 80)
	surface.FillText("B", 1520, 820)
	surface.FillText("C", 430, 820)

	surface.FillText("a", 900, 750)
	surface.FillText("b", 520, 450)
	surface.FillText("c", 900, 480)
	surface.WriteToPNG(filename)
}
