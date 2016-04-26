package stl

type Solid struct {
	Name   string
	Facets []Facet
}

type Facet struct {
	Normal   Vec3
	Vertices [3]Vec3
}

type Vec3 struct {
	X, Y, Z float64
}
