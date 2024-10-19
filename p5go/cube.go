package main

type Point struct {
	x, y, z float64
}

type QuadFace struct {
	nw, ne, se, sw Point
}

var oneFace = QuadFace{
	nw: Point{-1.0, +1.0, +1.0},
	ne: Point{+1.0, +1.0, +1.0},
	se: Point{+1.0, -1.0, +1.0},
	sw: Point{-1.0, -1.0, +1.0},
}
