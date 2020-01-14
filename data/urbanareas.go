package data

import (
	"github.com/everystreet/go-geojson/v2"
	"github.com/everystreet/naturalearth"
)

var UrbanAreas50 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: UrbanAreas50Name,
		Schemas: []naturalearth.Schema{
			func(feat geojson.Feature, meta *naturalearth.Meta) (string, error) {
				meta.AddProperty(PropType, TypePropLanduse)
				meta.AddProperty(PropLanduseClass, LanduserClassPropResidential)

				var scaleRank float64
				if err := feat.Properties.GetValue(PropScaleRank, &scaleRank); err != nil {
					return "", err
				}

				switch {
				case scaleRank < 3:
					meta.AddProperty(PropMinZoom, 4)
				default:
					meta.AddProperty(PropMinZoom, 5)
				}

				meta.AddProperty(PropMaxZoom, 5)
				return basicKey("urban_area_50m", feat)
			},
		},
	}
}

var UrbanAreas10 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: UrbanAreas10Name,
		Schemas: []naturalearth.Schema{
			func(feat geojson.Feature, meta *naturalearth.Meta) (string, error) {
				meta.AddProperty(PropType, TypePropLanduse)
				meta.AddProperty(PropLanduseClass, LanduserClassPropResidential)

				var scaleRank float64
				if err := feat.Properties.GetValue(PropScaleRank, &scaleRank); err != nil {
					return "", err
				}

				switch {
				case scaleRank < 6:
					meta.AddProperty(PropMinZoom, 6)
				case scaleRank < 7:
					meta.AddProperty(PropMinZoom, 7)
				case scaleRank < 8:
					meta.AddProperty(PropMinZoom, 8)
				default:
					return "", nil
				}

				meta.AddProperty(PropMaxZoom, 8)
				return basicKey("urban_area_10m", feat)
			},
		},
	}
}
