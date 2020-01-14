package data

import (
	"fmt"

	"github.com/everystreet/go-geojson/v2"
	"github.com/everystreet/naturalearth"
)

func basicKey(prefix string, feat geojson.Feature) (string, error) {
	var num uint
	if err := feat.Properties.GetValue(naturalearth.NumberPropertyName, &num); err != nil {
		return "", err
	}
	return fmt.Sprintf(prefix+"-%d", num), nil
}
