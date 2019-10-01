package data

import (
	"strings"

	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"
)

var BoundaryLines110 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: BoundaryLines110Name,
		Schemas: []naturalearth.Schema{
			func(feat geojson.Feature, meta *naturalearth.Meta) (string, error) {
				meta.AddProperty(PropType, TypePropBoundary)
				meta.AddProperty(PropMinZoom, 0)
				meta.AddProperty(PropMaxZoom, 0)
				return basicKey("boundary_110m", feat)
			},
		},
	}
}

var BoundaryLines50 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: BoundaryLines50Name,
		Schemas: []naturalearth.Schema{
			func(feat geojson.Feature, meta *naturalearth.Meta) (string, error) {
				meta.AddProperty(PropType, TypePropBoundary)
				meta.AddProperty(PropMinZoom, 1)
				meta.AddProperty(PropMaxZoom, 3)
				return basicKey("boundary_50m", feat)
			},
		},
	}
}

var BoundaryLines10 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: BoundaryLines10Name,
		Schemas: []naturalearth.Schema{
			func(feat geojson.Feature, meta *naturalearth.Meta) (string, error) {
				var class string
				if err := feat.Properties.GetType(PropFeatureClass, &class); err != nil {
					return "", err
				} else if strings.EqualFold(class, FeatureClassPropLeaseLimit) {
					return "", nil
				}

				meta.AddProperty(PropType, TypePropBoundary)
				meta.AddProperty(PropMinZoom, 4)
				meta.AddProperty(PropMaxZoom, 4)
				return basicKey("boundary_10m", feat)
			},
		},
	}
}
