package figs

import (
	"github.com/bit101/blgo"
)

func Fig_2_12() {
	filename := "images/figure_2-12.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)
	surface.SetLineWidth(4)

	surface.Translate(width/2, height/2)
	surface.SetDash([]float64{20, 10}, 1, 0)
	surface.StrokeCircle(0, 0, 400)

	surface.SetDash([]float64{0}, 0, 0)
	surface.MoveTo(0, 0)
	surface.ArcNegative(0, 0, 400, 0, -1)
	surface.ClosePath()
	surface.ArcNegative(0, 0, 100, 0, -1)
	surface.Stroke()

	angleText := "angle = 1 radian"
	te := surface.TextExtents(angleText)
	surface.FillText(angleText, -te.Width/2, te.Height+10)

	radiusText := "radius = 1"
	te = surface.TextExtents(radiusText)
	surface.FillText(radiusText, -150, -200)

	arcText := "length of"
	te = surface.TextExtents(arcText)
	surface.FillText(arcText, 370, -200)

	arcText2 := "arc = 1"
	te = surface.TextExtents(arcText2)
	surface.FillText(arcText2, 410, -140)

	degreeText := "(57.2958... degrees)"
	surface.SetFontSize(30)
	te = surface.TextExtents(degreeText)
	surface.FillText(degreeText, -te.Width/2, 110)

	surface.WriteToPNG(filename)
}
