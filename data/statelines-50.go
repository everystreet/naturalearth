package data

import (
	"fmt"

	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"
	"github.com/pkg/errors"
)

var StateLines50 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: StateLines50Name,
		Opts: []naturalearth.Option{
			naturalearth.AddProperty(geojson.Property{Name: "type", Value: "state"}),
			naturalearth.AddProperty(geojson.Property{Name: "max_zoom", Value: 3}),
			naturalearth.RenameProperty("scalerank", "scale_rank"),
			naturalearth.RenameProperty("min_zoom", "min_zoom"),
			naturalearth.RenameProperty("adm0_a3", "country_a3"),
		},
		GetKey: func(feat *geojson.Feature) (string, error) {
			var num uint
			if err := feat.Properties.GetType(naturalearth.NumberPropertyName, &num); err != nil {
				return "", errors.Wrap(err, "missing or invalid 'num' property")
			}
			return fmt.Sprintf("state_110-%d", num), nil
		},
	}
}
