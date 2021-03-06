package figs

import (
	"codingtrig/trig"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/util"
)

// Fig0407 blah
func Fig0407() {
	filename := "images/figure_4-7.png"
	width := 2000.0
	height := 1000.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)

	surface.Translate(width/2, height/2)

	surface.SetDash([]float64{10, 0}, 1, 0)

	surface.MoveTo(-width/2, 0)
	surface.LineTo(width/2, 0)
	surface.MoveTo(0, -height/2)
	surface.LineTo(0, height/2)
	surface.Stroke()

	surface.SetDash([]float64{1}, 0, 0)
	surface.SetLineWidth(4)
	surface.SetFontSize(50)

	trig.DrawRightTriangle(-300, 0, 300, -200, surface)
	trig.DrawRightTriangle(200, 0, -200, -300, surface)

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}
