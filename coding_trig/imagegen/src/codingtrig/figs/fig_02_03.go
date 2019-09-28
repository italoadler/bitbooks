package figs

import (
	"github.com/bit101/blgo"
)

func Fig_2_3() {
	filename := "images/figure_2-3.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(4)
	surface.SetFontSize(50)
	// trig.DrawRightTriangle(500, 800, 1000, 700, surface)
	surface.MoveTo(500, 100)
	surface.LineTo(1500, 800)
	surface.Stroke()

	surface.FillCircle(500, 100, 20)
	surface.FillCircle(1500, 800, 20)

	surface.FillText("A", 480, 60)
	surface.FillText("B", 1540, 820)

	surface.FillText("distance", 960, 400)
	surface.WriteToPNG(filename)
}
