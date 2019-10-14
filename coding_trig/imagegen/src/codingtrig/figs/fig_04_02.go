package figs

import (
	"codingtrig/trig"
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/util"
)

// Fig0402 blah
func Fig0402() {
	filename := "images/figure_4-2.png"
	width := 1850.0
	height := 700.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(4)
	surface.SetFontSize(50)

	hyp := 500.0
	angle := math.Pi / 6
	w := math.Cos(angle) * hyp
	h := math.Sin(angle) * hyp
	trig.DrawRightTriangle(100, 550, w, h, surface)
	surface.FillText("30°", 550, 550)
	surface.FillText("3", 50, 400)
	surface.FillText("6", 300, 380)
	surface.FillText("sin 30° = 3 / 6 = 0.5", 150, 650)

	trig.DrawRightTriangle(800, 550, w*2, h*2, surface)
	surface.FillText("30°", 1700, 550)
	surface.FillText("6", 750, 350)
	surface.FillText("12", 1200, 250)
	surface.FillText("sin 30° = 6 / 12 = 0.5", 850, 650)

	surface.SetLineWidth(4)
	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}
