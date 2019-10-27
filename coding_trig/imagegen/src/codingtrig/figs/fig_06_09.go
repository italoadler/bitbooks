package figs

import (
	"codingtrig/trig"
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/util"
)

// Fig0609 blah
func Fig0609() {
	filename := "images/figure_6-9.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)

	h := 1000 * math.Tan(math.Pi/6)
	trig.DrawRightTriangle(1500, 800, -1000, h, surface)

	surface.FillCircle(1500, 800-h, 20)
	surface.FillCircle(500, 800, 20)

	surface.SetFontSize(50)
	surface.FillText("point A", 260, 820)
	surface.FillText("point B", 1540, 800-h)
	surface.FillText("angle = ?", 670, 760)
	surface.FillText("B.y - A.y", 1540, 800-h/2)
	surface.FillText("B.x - A.x", 900, 860)

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}
