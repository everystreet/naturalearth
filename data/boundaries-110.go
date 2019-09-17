package data

import (
	"fmt"

	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"
	"github.com/pkg/errors"
)

var Boundaries110 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: Boundaries110Name,
		Opts: []naturalearth.Option{
			naturalearth.AddProperty(geojson.Property{Name: "type", Value: "boundary"}),
			naturalearth.AddProperty(geojson.Property{Name: "max_zoom", Value: 0}),
			naturalearth.RenameProperty("scalerank", "scale_rank"),
			naturalearth.RenameProperty("min_zoom", "min_zoom"),
		},
		GetKey: func(feat *geojson.Feature) (string, error) {
			var num uint
			if err := feat.Properties.GetType(naturalearth.NumberPropertyName, &num); err != nil {
				return "", errors.Wrap(err, "missing or invalid 'num' property")
			}
			return fmt.Sprintf("boundary_110-%d", num), nil
		},
	}
}
