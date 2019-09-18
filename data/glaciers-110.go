package data

import (
	"github.com/mercatormaps/naturalearth"
)

var Glaciers110 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: Glaciers110Name,
		Opts: []naturalearth.Option{
			naturalearth.AddProperty(PropType, TypePropLandcover),
			naturalearth.AddProperty(PropMinZoom, 0),
			naturalearth.AddProperty(PropMaxZoom, 1),
			naturalearth.AddProperty(PropLandcoverClass, LandcoverClassPropIce),
			naturalearth.AddProperty(PropLandcoverSubclass, LandcoverSubclassPropGlacier),
		},
		GetKey: BasicKey("glacier_110m"),
	}
}
