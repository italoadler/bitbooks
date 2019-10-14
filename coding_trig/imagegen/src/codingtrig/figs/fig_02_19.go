package figs

import (
	"codingtrig/trig"

	"github.com/bit101/blgo"
)

// Fig0219 blah
func Fig0219() {
	filename := "images/figure_2-19.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)
	surface.SetLineWidth(4)

	trig.DrawTriangle(500, 100, 400, 700, 600, 700, surface)
	trig.DrawTriangle(1250, 100, 700, 500, 1900, 500, surface)

	surface.FillText("acute", 500-surface.TextExtents("acute").Width/2, 800)
	surface.FillText("obtuse", 1250-surface.TextExtents("obtuse").Width/2, 600)

	surface.WriteToPNG(filename)
}
