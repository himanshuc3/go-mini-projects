package solver

import "image"

func neighbours(p image.Point) [4]image.Point {
	// NOTE:
	// 1. Defining an array, not a slice - allocated at the compilation step
	return [...]image.Point{
		{p.X, p.Y + 1},
		{p.X, p.Y - 1},
		{p.X + 1, p.Y},
		{p.X - 1, p.Y},
	}
}
