package data

import (
	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"
)

var StateLines50 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: StateLines50Name,
		Schemas: []naturalearth.Schema{
			func(feat geojson.Feature, meta *naturalearth.Meta) (string, error) {
				meta.AddProperty(PropType, TypePropBoundary)
				meta.AddProperty(PropMinZoom, 1)
				meta.AddProperty(PropMaxZoom, 3)
				return basicKey("state_50m", feat)
			},
		},
	}
}

var StateLines10 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: StateLines10Name,
		Schemas: []naturalearth.Schema{
			func(feat geojson.Feature, meta *naturalearth.Meta) (string, error) {
				var minZoom float64
				if err := feat.Properties.GetType(PropMinZoom, &minZoom); err != nil {
					return "", err
				} else if minZoom > 5 {
					return "", nil
				}

				meta.AddProperty(PropType, TypePropBoundary)
				meta.AddProperty(PropMinZoom, 4)
				meta.AddProperty(PropMaxZoom, 4)
				return basicKey("state_10m", feat)
			},
		},
	}
}
