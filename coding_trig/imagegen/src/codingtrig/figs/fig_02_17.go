package figs

import (
	"math"

	"github.com/bit101/blgo"
)

func Fig_2_17() {
	filename := "images/figure_2-17.png"
	width := 2000.0
	height := 800.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)
	surface.SetLineWidth(4)

	surface.StrokePolygon(width/2, 500, 400, 3, -math.Pi/2)

	surface.FillText("60°", 960, 240)
	surface.FillText("60°", 1200, 670)
	surface.FillText("60°", 720, 670)

	surface.WriteToPNG(filename)
}
