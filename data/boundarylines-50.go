package data

import (
	"github.com/mercatormaps/naturalearth"
)

var BoundaryLines50 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: BoundaryLines50Name,
		Opts: []naturalearth.Option{
			naturalearth.AddProperty(PropType, TypePropBoundary),
			naturalearth.AddProperty(PropMinZoom, 1),
			naturalearth.AddProperty(PropMaxZoom, 3),
		},
		GetKey: BasicKey("boundary_50m"),
	}
}
