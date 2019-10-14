package figs

import (
	"codingtrig/trig"
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/util"
)

// Fig0401 blah
func Fig0401() {
	filename := "images/figure_4-1.png"
	width := 2000.0
	height := 650.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(4)
	surface.SetFontSize(50)

	hyp := 500.0
	angle := math.Pi / 6
	w := math.Cos(angle) * hyp
	h := math.Sin(angle) * hyp
	trig.DrawRightTriangle(100, 500, w, h, surface)
	surface.FillText("A", 550, 500)
	surface.FillText("a", 50, 350)
	surface.FillText("b", 250, 550)
	surface.SetLineWidth(3)
	surface.Ray(300, 340, angle+math.Pi/2, 0, 50)

	surface.SetLineWidth(4)
	angle = math.Pi / 5
	w = math.Cos(angle) * hyp
	h = math.Sin(angle) * hyp
	trig.DrawRightTriangle(650, 500, w, h, surface)
	surface.FillText("A", 1070, 500)
	surface.FillText("a", 600, 350)
	surface.FillText("b", 800, 550)
	surface.SetLineWidth(3)
	surface.Ray(860, 330, angle+math.Pi/2, 0, 50)

	surface.SetLineWidth(4)
	angle = math.Pi / 4
	w = math.Cos(angle) * hyp
	h = math.Sin(angle) * hyp
	trig.DrawRightTriangle(1170, 500, w, h, surface)
	surface.FillText("A", 1540, 500)
	surface.FillText("a", 1110, 350)
	surface.FillText("b", 1310, 550)
	surface.SetLineWidth(3)
	surface.Ray(1360, 300, angle+math.Pi/2, 0, 50)

	surface.SetLineWidth(4)
	angle = math.Pi / 3
	w = math.Cos(angle) * hyp
	h = math.Sin(angle) * hyp
	trig.DrawRightTriangle(1640, 500, w, h, surface)
	surface.FillText("A", 1900, 500)
	surface.FillText("a", 1590, 350)
	surface.FillText("b", 1750, 550)
	surface.SetLineWidth(3)
	surface.Ray(1800, 300, angle+math.Pi/2, 0, 50)

	surface.SetLineWidth(4)
	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}
