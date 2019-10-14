package figs

import (
	"codingtrig/trig"
	"math"

	"github.com/bit101/blgo"
)

// Fig0302 blah
func Fig0302() {
	filename := "images/figure_3-2.png"
	width := 2000.0
	height := 1300.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(4)
	surface.SetFontSize(50)
	surface.Translate(300, 50)

	trig.DrawRightTriangle(500, 800, 400, 400, surface)
	surface.StrokeRectangle(500, 800, 400, 400)
	surface.StrokeRectangle(100, 400, 400, 400)
	surface.Save()
	size := math.Hypot(400, 400)
	surface.Translate(500, 400)
	surface.Rotate(-math.Pi / 4)
	surface.StrokeRectangle(0, 0, size, size)
	surface.Restore()

	surface.FillText("A", 480, 380)
	surface.FillText("B", 920, 820)
	surface.FillText("C", 430, 880)

	surface.FillText("a", 680, 850)
	surface.FillText("b", 440, 620)
	surface.FillText("c", 740, 600)
	surface.WriteToPNG(filename)
}
