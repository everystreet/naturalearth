package data

import (
	"fmt"
	"strings"

	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"
)

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
		GetKey: func(feat *geojson.Feature) (string, error) {
			var num uint
			if err := feat.Properties.GetType(naturalearth.NumberPropertyName, &num); err != nil {
				return "", err
			}
			return fmt.Sprintf("boundary_10m-%d", num), nil
		},
	}
}
