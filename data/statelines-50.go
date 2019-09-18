package data

import (
	"github.com/mercatormaps/naturalearth"
)

var StateLines50 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: StateLines50Name,
		Opts: []naturalearth.Option{
			naturalearth.AddProperty(PropType, TypePropBoundary),
			naturalearth.AddProperty(PropMinZoom, 1),
			naturalearth.AddProperty(PropMaxZoom, 3),
		},
		GetKey: BasicKey("state_50m"),
	}
}
