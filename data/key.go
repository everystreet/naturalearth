package data

import (
	"fmt"

	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"
)

func BasicKey(prefix string) naturalearth.KeyGetter {
	return func(feat *geojson.Feature) (string, error) {
		var num uint
		if err := feat.Properties.GetType(naturalearth.NumberPropertyName, &num); err != nil {
			return "", err
		}
		return fmt.Sprintf(prefix+"-%d", num), nil
	}
}
