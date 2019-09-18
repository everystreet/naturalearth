package data

import (
	"fmt"

	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"
)

var Boundaries10 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: Boundaries10Name,
		Opts: []naturalearth.Option{
			naturalearth.AddProperty(PropType, TypePropBoundary),
			naturalearth.AddProperty(PropMinZoom, 4),
			naturalearth.AddProperty(PropMaxZoom, 4),
		},
		GetKey: func(feat *geojson.Feature) (string, error) {
			var num uint
			if err := feat.Properties.GetType(naturalearth.NumberPropertyName, &num); err != nil {
				return "", err
			}
			return fmt.Sprintf("boundary_10m-%d", num), nil
		},
	}
}
