package figs

import (
	"github.com/bit101/blgo"
)

func Fig_1_13() {
	filename := "images/figure_1-13.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)
	surface.SetLineWidth(4)

	surface.Translate(width/2, height/2)
	surface.StrokeCircle(0, 0, 400)

	surface.MoveTo(0, 0)
	surface.LineTo(410, 0)
	surface.Stroke()

	surface.Save()
	for i := 1; i <= 6; i++ {
		surface.Rotate(-1)
		surface.MoveTo(390, 0)
		surface.LineTo(410, 0)
		surface.Stroke()
	}
	surface.Restore()

	surface.FillText("1 radian", 250, -340)
	surface.FillText("2 radians", -440, -370)
	surface.FillText("3 radians", -660, -70)
	surface.FillText("4 radians", -530, 340)
	surface.FillText("5 radians", 170, 420)
	surface.FillText("6 radians", 420, 160)

	surface.WriteToPNG(filename)
}
