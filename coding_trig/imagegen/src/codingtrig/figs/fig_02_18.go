package figs

import (
	"codingtrig/trig"
	"math"

	"github.com/bit101/blgo"
)

func Fig_2_18() {
	filename := "images/figure_2-18.png"
	width := 2000.0
	height := 800.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)
	surface.SetLineWidth(4)

	trig.DrawTriangle(1000, 100, 1300, 700, 700, 700, surface)

	surface.Arc(1000, 100, 100, math.Atan2(600, 300), math.Atan2(600, -300))
	surface.Stroke()

	surface.Arc(1300, 700, 100, math.Pi, math.Atan2(-600, -300))
	surface.Stroke()

	surface.ArcNegative(700, 700, 100, 0, math.Atan2(-600, 300))
	surface.Stroke()

	surface.Save()
	surface.Translate(1300, 700)
	surface.Rotate(-2.5)
	surface.MoveTo(90, 0)
	surface.LineTo(110, 0)
	surface.Stroke()

	surface.Rotate(-0.25)
	surface.MoveTo(90, 0)
	surface.LineTo(110, 0)
	surface.Stroke()
	surface.Restore()

	surface.Save()
	surface.Translate(700, 700)
	surface.Rotate(-0.4)
	surface.MoveTo(90, 0)
	surface.LineTo(110, 0)
	surface.Stroke()

	surface.Rotate(-0.25)
	surface.MoveTo(90, 0)
	surface.LineTo(110, 0)
	surface.Stroke()
	surface.Restore()

	surface.MoveTo(825, 400)
	surface.LineTo(865, 420)
	surface.Stroke()

	surface.MoveTo(1135, 420)
	surface.LineTo(1175, 400)
	surface.Stroke()

	surface.WriteToPNG(filename)
}
