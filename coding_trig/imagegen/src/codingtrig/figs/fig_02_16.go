package figs

import (
	"math"

	"github.com/bit101/blgo"
)

func Fig_2_16() {
	filename := "images/figure_2-16.png"
	width := 2000.0
	height := 280.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)
	surface.SetLineWidth(4)

	surface.Save()
	surface.Translate(1000, 100)
	surface.Rotate(math.Pi / 2)
	angle := math.Pi / 2
	surface.MoveTo(math.Cos(angle)*600, math.Sin(angle)*600)
	surface.LineTo(math.Cos(-angle)*600, math.Sin(-angle)*600)
	surface.ClosePath()
	surface.Stroke()
	surface.ArcNegative(0, 0, 30, angle, -angle)
	surface.Stroke()
	surface.Restore()

	surface.FillText("A", 983, 80)
	surface.FillText("a", 983, 200)
	surface.FillText("b", 760, 80)
	surface.FillText("c", 1200, 80)

	surface.WriteToPNG(filename)
}
