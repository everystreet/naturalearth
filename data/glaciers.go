package data

import (
	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"
)

var Glaciers110 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: Glaciers110Name,
		Schemas: []naturalearth.Schema{
			func(feat geojson.Feature, meta *naturalearth.Meta) (string, error) {
				meta.AddProperty(PropType, TypePropLandcover)
				meta.AddProperty(PropMinZoom, 0)
				meta.AddProperty(PropMaxZoom, 1)
				meta.AddProperty(PropLandcoverClass, LandcoverClassPropIce)
				meta.AddProperty(PropLandcoverSubclass, LandcoverSubclassPropGlacier)
				return basicKey("glacier_110m", feat)
			},
		},
	}
}

var Glaciers50 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: Glaciers50Name,
		Schemas: []naturalearth.Schema{
			func(feat geojson.Feature, meta *naturalearth.Meta) (string, error) {
				meta.AddProperty(PropType, TypePropLandcover)
				meta.AddProperty(PropMinZoom, 2)
				meta.AddProperty(PropMaxZoom, 4)
				meta.AddProperty(PropLandcoverClass, LandcoverClassPropIce)
				meta.AddProperty(PropLandcoverSubclass, LandcoverSubclassPropGlacier)
				return basicKey("glacier_50m", feat)
			},
		},
	}
}

var Glaciers10 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: Glaciers10Name,
		Schemas: []naturalearth.Schema{
			func(feat geojson.Feature, meta *naturalearth.Meta) (string, error) {
				meta.AddProperty(PropType, TypePropLandcover)
				meta.AddProperty(PropMinZoom, 5)
				meta.AddProperty(PropMaxZoom, 6)
				meta.AddProperty(PropLandcoverClass, LandcoverClassPropIce)
				meta.AddProperty(PropLandcoverSubclass, LandcoverSubclassPropGlacier)
				return basicKey("glacier_10m", feat)
			},
		},
	}
}
