package figs

import (
	"codingtrig/trig"
	"math"

	"github.com/bit101/blgo"
)

func Fig_1_3() {
	filename := "images/figure_1-3.png"
	width := 2000.0
	height := 900.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetFontSize(50)

	trig.DrawTriangle(1000, 100, 1500, 800, 600, 700, surface)

	te := surface.TextExtents("A")

	surface.FillText("A", 1000-te.Width/2, 80)
	surface.FillText("B", 1520, 820)
	surface.FillText("C", 580-te.Width, 720)

	surface.FillText("c", 1250, 400)
	surface.FillText("a", 920, 810)
	surface.FillText("b", 720, 400)

	surface.SetLineWidth(4)
	surface.Arc(1000, 100, 80, math.Atan2(700, 500), math.Atan2(600, -400))
	surface.Stroke()

	surface.ArcNegative(1500, 800, 80, math.Atan2(-700, -500), math.Atan2(-100, -900))
	surface.Stroke()

	surface.ArcNegative(600, 700, 80, math.Atan2(100, 900), math.Atan2(-600, 400))
	surface.Stroke()

	surface.WriteToPNG(filename)
}
