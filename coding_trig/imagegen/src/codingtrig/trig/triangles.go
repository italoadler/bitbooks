package trig

import "github.com/bit101/blgo"

func DrawTriangle(x0, y0, x1, y1, x2, y2 float64, s *blgo.Surface) {
	s.Save()
	s.SetSourceRGB(0, 0, 0)
	s.SetLineWidth(4)
	s.MoveTo(x0, y0)
	s.LineTo(x1, y1)
	s.LineTo(x2, y2)
	s.ClosePath()
	s.Stroke()
	s.Restore()
}

func DrawRightTriangle(x, y, w, h float64, s *blgo.Surface) {
	DrawTriangle(x, y, x, y-h, x+w, y, s)
	s.Save()
	s.SetSourceRGB(0, 0, 0)
	s.SetLineWidth(4)
	rSize := 70.0
	rw := rSize
	rh := -rSize
	if w < 0 {
		rw = -rSize
	}
	if h < 0 {
		rh = rSize
	}
	s.StrokeRectangle(x, y, rw, rh)
	s.Restore()
}
