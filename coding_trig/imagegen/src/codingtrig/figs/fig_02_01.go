package figs

import (
	"codingtrig/trig"

	"github.com/bit101/blgo"
)

// Fig0201 blah
func Fig0201() {
	filename := "images/figure_2-1.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)

	trig.DrawTriangle(1000, 100, 1500, 800, 600, 700, surface)

	te := surface.TextExtents("angle")

	surface.FillText("angle", 1000-te.Width/2, 80)
	surface.FillText("angle", 1520, 820)
	surface.FillText("angle", 580-te.Width, 720)

	te = surface.TextExtents("side")
	surface.FillText("side", 1250, 400)
	surface.FillText("side", 920, 810)
	surface.FillText("side", 670, 400)

	surface.WriteToPNG(filename)
}
