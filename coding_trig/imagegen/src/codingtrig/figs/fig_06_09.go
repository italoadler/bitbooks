package figs

import (
	"codingtrig/trig"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/util"
)

// Fig0609 blah
func Fig0609() {
	filename := "images/figure_6-9.png"
	width := 800.0
	height := 600.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)

	trig.DrawRightTriangle(550, 400, -350, 250, surface)

	surface.FillCircle(200, 400, 10)
	surface.FillCircle(550, 150, 10)

	surface.SetFontSize(30)
	surface.FillText("point A", 60, 410)
	surface.FillText("point B", 580, 140)
	surface.FillText("B.y - A.y", 580, 280)
	surface.FillText("B.x - A.x", 320, 440)

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}
