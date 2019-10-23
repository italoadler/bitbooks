package figs

import (
	"codingtrig/trig"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/util"
)

// Fig0608 blah
func Fig0608() {
	filename := "images/figure_6-8.png"
	width := 800.0
	height := 600.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)

	trig.DrawRightTriangle(550, 400, -350, 250, surface)

	surface.SetFontSize(30)
	surface.FillText("A", 160, 410)
	surface.FillText("B.y - A.y", 580, 280)
	surface.FillText("B.x - A.x", 320, 440)

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}
