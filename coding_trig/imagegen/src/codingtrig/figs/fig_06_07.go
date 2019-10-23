package figs

import (
	"codingtrig/trig"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/util"
)

// Fig0607 blah
func Fig0607() {
	filename := "images/figure_6-7.png"
	width := 800.0
	height := 600.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)

	trig.DrawRightTriangle(550, 400, -350, 250, surface)

	surface.SetFontSize(30)
	surface.FillText("A", 160, 400)
	surface.FillText("30 degrees", 90, 440)
	surface.FillText("y = ???", 580, 280)
	surface.FillText("x = 10", 330, 440)

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}
