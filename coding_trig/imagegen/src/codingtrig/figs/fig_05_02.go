package figs

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

// Fig0502 blah
func Fig0502() {
	filename := "images/figure_5-2.png"
	width := 2000.0
	height := 1000.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)

	surface.Translate(width/2, height/2)
	surface.Scale(1, -1)

	surface.SetDash([]float64{10, 0}, 1, 0)

	surface.MoveTo(-width/2, 0)
	surface.LineTo(width/2, 0)
	surface.MoveTo(0, -height/2)
	surface.LineTo(0, height/2)
	surface.Stroke()

	surface.SetDash([]float64{1}, 0, 0)
	surface.SetLineWidth(4)
	surface.SetFontSize(50)

	angle := math.Pi / 6
	radius := 300.0

	drawTri(surface, angle, radius)
	angle += math.Pi/2 + random.FloatRange(-0.03, 0.03)
	drawTri(surface, angle, radius)
	angle += math.Pi/2 + random.FloatRange(-0.03, 0.03)
	drawTri(surface, angle, radius)
	angle += math.Pi/2 + random.FloatRange(-0.03, 0.03)
	drawTri(surface, angle, radius)
	angle += math.Pi/2 + random.FloatRange(-0.03, 0.03)

	surface.SetDash([]float64{10, 0}, 1, 0)
	surface.SetLineWidth(2)
	surface.StrokeCircle(0, 0, radius)

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

func drawTri(surface *blgo.Surface, angle, radius float64) {
	x := math.Cos(angle) * radius
	y := math.Sin(angle) * radius
	surface.MoveTo(0, 0)
	surface.LineTo(x, 0)
	surface.LineTo(x, y)
	surface.LineTo(0, 0)
	surface.MoveTo(x-30*sign(x), 0)
	surface.LineTo(x-30*sign(x), 30*sign(y))
	surface.LineTo(x, 30*sign(y))
	surface.Stroke()
}

func sign(x float64) float64 {
	if x < 0 {
		return -1
	}
	return 1
}
