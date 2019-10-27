package figs

import (
	"codingtrig/trig"
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/util"
)

// Fig0607 blah
func Fig0607() {
	filename := "images/figure_6-7.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)

	h := 1000 * math.Tan(math.Pi/6)
	trig.DrawRightTriangle(1500, 800, -1000, h, surface)

	surface.SetFontSize(50)
	surface.FillText("A", 420, 820)
	surface.FillText("30 degrees", 670, 760)
	surface.FillText("y = ???", 1540, 800-h/2)
	surface.FillText("x = 10", 900, 860)

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}
