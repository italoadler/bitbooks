package figs

import (
	"codingtrig/trig"

	"github.com/bit101/blgo"
)

func Fig_1_2() {
	filename := "images/figure_1-2.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)

	trig.DrawTriangle(1000, 100, 1500, 800, 600, 700, surface)

	te := surface.TextExtents("A")

	surface.FillText("A", 1000-te.Width/2, 80)
	surface.FillText("B", 1520, 820)
	surface.FillText("C", 580-te.Width, 720)

	surface.FillText("c", 1250, 400)
	surface.FillText("a", 920, 810)
	surface.FillText("b", 720, 400)

	surface.WriteToPNG(filename)
}
