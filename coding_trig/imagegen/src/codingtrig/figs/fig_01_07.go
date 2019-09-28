package figs

import (
	"fmt"
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/color"
)

func Fig_1_7() {
	filename := "images/figure_1-7.png"
	width := 2000.0
	height := 500.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(4)
	surface.SetFontSize(50)
	angles := []int{0, -90, -180, -270, -360}

	diam := width / 5
	radius := diam/2 - 20
	for i := 0.0; i < 5; i++ {
		x := i*diam + diam/2
		surface.Save()
		surface.Translate(x, height*0.45)
		surface.StrokeCircle(0, 0, radius)
		surface.MoveTo(0, 0)
		surface.Arc(0, 0, radius, 0, i*math.Pi/2)
		surface.ClosePath()
		surface.SetSourceColor(color.Grey(0.75))
		surface.FillPreserve()
		surface.SetSourceRGB(0, 0, 0)
		surface.Stroke()

		label := fmt.Sprintf("%d°", angles[int(i)])
		surface.FillText(label, -surface.TextExtents(label).Width/2, 250)

		surface.Restore()
	}

	surface.WriteToPNG(filename)
}
