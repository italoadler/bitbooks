package figs

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/util"
)

// Fig0508 blah
func Fig0508() {
	filename := "images/figure_5-8.png"
	width := 2000.0
	height := 1000.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)

	rx := 500.0
	ry := 300.0

	surface.Translate(width/2, height/2)

	surface.SetDash([]float64{10, 0}, 1, 0)

	surface.MoveTo(-width/2, 0)
	surface.LineTo(width/2, 0)
	surface.MoveTo(0, -height/2)
	surface.LineTo(0, height/2)
	surface.Stroke()

	surface.SetDash([]float64{1}, 0, 0)
	surface.SetLineWidth(4)
	surface.SetFontSize(30)

	surface.SetLineWidth(2)
	surface.StrokeEllipse(0, 0, rx, ry)

	surface.FillText("major axis", 150, -10)
	surface.FillText("minor axis", -92, -120)

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}
