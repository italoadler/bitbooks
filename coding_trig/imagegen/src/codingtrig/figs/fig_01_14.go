package figs

import (
	"math"

	"github.com/bit101/blgo"
)

func Fig_1_14() {
	filename := "images/figure_1-14.png"
	width := 2000.0
	height := 740.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)
	surface.SetLineWidth(4)

	surface.Save()
	surface.Translate(350, 100)
	surface.Rotate(math.Pi / 2)
	surface.MoveTo(0, 0)
	angle := 0.5
	surface.LineTo(math.Cos(angle)*600, math.Sin(angle)*600)
	surface.LineTo(math.Cos(-angle)*600, math.Sin(-angle)*600)
	surface.ClosePath()
	surface.Stroke()
	surface.ArcNegative(0, 0, 100, angle, -angle)
	surface.Stroke()
	surface.Restore()

	surface.Save()
	surface.Translate(1350, 100)
	surface.Rotate(math.Pi / 2)
	surface.MoveTo(0, 0)
	angle = 1.0
	surface.LineTo(math.Cos(1)*600, math.Sin(angle)*600)
	surface.LineTo(math.Cos(-angle)*600, math.Sin(-angle)*600)
	surface.ClosePath()
	surface.Stroke()
	surface.ArcNegative(0, 0, 100, angle, -angle)
	surface.Stroke()
	surface.Restore()

	surface.FillText("A", 333, 80)
	surface.FillText("a", 333, 680)
	surface.FillText("b", 120, 380)
	surface.FillText("c", 540, 380)

	surface.FillText("A", 1333, 80)
	surface.FillText("a", 1333, 480)
	surface.FillText("b", 1050, 240)
	surface.FillText("c", 1600, 240)

	surface.WriteToPNG(filename)
}
