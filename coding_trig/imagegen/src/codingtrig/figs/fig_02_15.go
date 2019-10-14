package figs

import (
	"math"

	"github.com/bit101/blgo"
)

// Fig0215 blah
func Fig0215() {
	filename := "images/figure_2-15.png"
	width := 2000.0
	height := 280.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)
	surface.SetLineWidth(4)

	surface.Save()
	surface.Translate(1000, 100)
	surface.Rotate(math.Pi / 2)
	surface.MoveTo(0, 0)
	angle := 1.5
	surface.LineTo(math.Cos(angle)*600, math.Sin(angle)*600)
	surface.LineTo(math.Cos(-angle)*600, math.Sin(-angle)*600)
	surface.ClosePath()
	surface.Stroke()
	surface.ArcNegative(0, 0, 30, angle, -angle)
	surface.Stroke()
	surface.Restore()

	surface.FillText("A", 983, 80)
	surface.FillText("a", 983, 200)
	surface.FillText("b", 760, 100)
	surface.FillText("c", 1200, 100)

	surface.WriteToPNG(filename)
}
