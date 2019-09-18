package data

import (
	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"
)

var StateLines10 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: StateLines10Name,
		Opts: []naturalearth.Option{
			naturalearth.AddProperty(PropType, TypePropBoundary),
			naturalearth.AddProperty(PropMinZoom, 4),
			naturalearth.AddProperty(PropMaxZoom, 4),
		},
		ShouldStore: func(feat *geojson.Feature) (bool, error) {
			var minZoom float64
			if err := feat.Properties.GetType(PropMinZoom, &minZoom); err != nil {
				return false, err
			}
			return minZoom <= 5, nil
		},
		GetKey: BasicKey("state_10m"),
	}
}
