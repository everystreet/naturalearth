package data

import (
	"github.com/mercatormaps/naturalearth"
)

var BoundaryLines110 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: BoundaryLines110Name,
		Opts: []naturalearth.Option{
			naturalearth.AddProperty(PropType, TypePropBoundary),
			naturalearth.AddProperty(PropMinZoom, 0),
			naturalearth.AddProperty(PropMaxZoom, 0),
		},
		GetKey: BasicKey("boundary_110m"),
	}
}
