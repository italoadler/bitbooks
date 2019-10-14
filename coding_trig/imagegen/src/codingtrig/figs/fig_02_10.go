package figs

import (
	"fmt"

	"github.com/bit101/blgo"
)

// Fig0210 blah
func Fig0210() {
	filename := "images/figure_2-10.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(30)
	surface.SetLineWidth(4)

	surface.Translate(width/2, height/2)
	surface.MoveTo(-400, 0)
	surface.LineTo(400, 0)
	surface.MoveTo(0, -400)
	surface.LineTo(0, 400)

	i := -4
	for x := -400.0; x <= 400.0; x += 100 {
		surface.MoveTo(x, -10)
		surface.LineTo(x, 10)
		surface.Stroke()
		if i != 0 {
			c := fmt.Sprintf("%d", i)
			surface.FillText(c, x-surface.TextExtents(c).Width/2, -15)
		}
		i++
	}
	i = -4
	for y := -400.0; y <= 400.0; y += 100 {
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
