package data

import (
	"fmt"

	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"
)

var Boundaries50 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: Boundaries50Name,
		Opts: []naturalearth.Option{
			naturalearth.AddProperty(PropType, TypePropBoundary),
			naturalearth.AddProperty(PropMinZoom, 1),
			naturalearth.AddProperty(PropMaxZoom, 3),
		},
		GetKey: func(feat *geojson.Feature) (string, error) {
			var num uint
			if err := feat.Properties.GetType(naturalearth.NumberPropertyName, &num); err != nil {
				return "", err
			}
			return fmt.Sprintf("boundary_50m-%d", num), nil
		},
	}
}
