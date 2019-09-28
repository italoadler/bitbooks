package figs

import (
	"math"

	"github.com/bit101/blgo"
)

func Fig_0_1() {
	filename := "images/figure_0-1.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(4)

	for x := 0.0; x < width; x++ {
		y := height/2 - math.Sin(x/width*math.Pi*6)*height*0.45
		surface.LineTo(x, y)
	}
	surface.Stroke()

	surface.WriteToPNG(filename)
}
