package data

import (
	"fmt"

	"github.com/mercatormaps/go-geojson"
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
		GetKey: func(feat *geojson.Feature) (string, error) {
			var num uint
			if err := feat.Properties.GetType(naturalearth.NumberPropertyName, &num); err != nil {
				return "", err
			}
			return fmt.Sprintf("glacier_110m-%d", num), nil
		},
	}
}
