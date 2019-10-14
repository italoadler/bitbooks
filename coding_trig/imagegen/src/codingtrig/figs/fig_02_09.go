package figs

import (
	"fmt"

	"github.com/bit101/blgo"
)

// Fig0209 blah
func Fig0209() {
	filename := "images/figure_2-9.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(30)
	surface.SetLineWidth(4)

	surface.Translate(width/2-400, height/2-400)
	surface.MoveTo(0, 0)
	surface.LineTo(0, 800)
	surface.MoveTo(0, 0)
	surface.LineTo(800, 0)

	i := 0
	for x := 0.0; x <= 800.0; x += 100 {
		surface.MoveTo(x, -10)
		surface.LineTo(x, 10)
		surface.Stroke()
		if i != 0 {
			c := fmt.Sprintf("%d", i)
			surface.FillText(c, x-surface.TextExtents(c).Width/2, -15)
		}
		i++
	}
	i = 0
	for y := 0.0; y <= 800.0; y += 100 {
		surface.MoveTo(-10, y)
		surface.LineTo(10, y)
		surface.Stroke()
		if i != 0 {
			c := fmt.Sprintf("%d", i)
			surface.FillText(c, 15, y+surface.TextExtents(c).Height/2)
		}
		i++
	}
	surface.FillText("0", 10, -15)

	surface.WriteToPNG(filename)
}
