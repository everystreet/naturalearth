package data

import (
	"fmt"

	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"
)

var Boundaries110 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: Boundaries110Name,
		Opts: []naturalearth.Option{
			naturalearth.AddProperty(PropType, TypePropBoundary),
			naturalearth.AddProperty(PropMinZoom, 0),
			naturalearth.AddProperty(PropMaxZoom, 0),
		},
		GetKey: func(feat *geojson.Feature) (string, error) {
			var num uint
			if err := feat.Properties.GetType(naturalearth.NumberPropertyName, &num); err != nil {
				return "", err
			}
			return fmt.Sprintf("boundary_110m-%d", num), nil
		},
	}
}
