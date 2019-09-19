package data

import (
	"strings"

	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"
)

var BoundaryLines110 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: BoundaryLines110Name,
		Opts: []naturalearth.Option{
			naturalearth.AddProperty(PropType, TypePropBoundary),
			naturalearth.AddProperty(PropMinZoom, 0),
			naturalearth.AddProperty(PropMaxZoom, 0),
		},
		GetKey: BasicKey("boundary_110m"),
	}
}

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

var BoundaryLines10 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: BoundaryLines10Name,
		Opts: []naturalearth.Option{
			naturalearth.AddProperty(PropType, TypePropBoundary),
			naturalearth.AddProperty(PropMinZoom, 4),
			naturalearth.AddProperty(PropMaxZoom, 4),
		},
		ShouldStore: func(feat *geojson.Feature) (bool, error) {
			var class string
			if err := feat.Properties.GetType(PropFeatureClass, &class); err != nil {
				return false, err
			}
			return !strings.EqualFold(class, FeatureClassPropLeaseLimit), nil
		},
		GetKey: BasicKey("boundary_10m"),
	}
}
