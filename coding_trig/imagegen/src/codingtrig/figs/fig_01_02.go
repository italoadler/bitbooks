package figs

import (
	"codingtrig/trig"
	"math"

	"github.com/bit101/blgo"
)

func Fig_0_2() {
	filename := "images/figure_0-2.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)

	trig.DrawRightTriangle(500, 800, 1000, 700, surface)
	surface.FillCircle(1500, 800, 20)
	surface.FillCircle(500, 100, 20)

	surface.SetLineWidth(4)
	surface.Arc(1500, 800, 120, math.Pi, math.Pi+math.Atan2(700, 1000))
	surface.Stroke()

	surface.FillText("angle", 1200, 760)
	surface.FillText("distance", 1000, 400)

	surface.WriteToPNG(filename)
}
