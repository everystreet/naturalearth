package data

import (
	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"
)

var UrbanAreas50 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: UrbanAreas50Name,
		Schemas: []naturalearth.Schema{
			{
				Opts: []naturalearth.Option{
					naturalearth.AddProperty(PropType, TypePropLanduse),
					naturalearth.AddProperty(PropMinZoom, 4),
					naturalearth.AddProperty(PropMaxZoom, 4),
				},
				ShouldStore: scaleRankIsLessThan(3),
				GetKey:      BasicKey("urban_area_50m_0"),
			},
			{
				Opts: []naturalearth.Option{
					naturalearth.AddProperty(PropType, TypePropLanduse),
					naturalearth.AddProperty(PropMinZoom, 5),
					naturalearth.AddProperty(PropMaxZoom, 5),
				},
				GetKey: BasicKey("urban_area_50m_1"),
			},
		},
	}
}

var UrbanAreas10 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: UrbanAreas10Name,
		Schemas: []naturalearth.Schema{
			{
				Opts: []naturalearth.Option{
					naturalearth.AddProperty(PropType, TypePropLanduse),
					naturalearth.AddProperty(PropMinZoom, 6),
					naturalearth.AddProperty(PropMaxZoom, 6),
				},
				ShouldStore: scaleRankIsLessThan(6),
				GetKey:      BasicKey("urban_area_10m_0"),
			},
			{
				Opts: []naturalearth.Option{
					naturalearth.AddProperty(PropType, TypePropLanduse),
					naturalearth.AddProperty(PropMinZoom, 7),
					naturalearth.AddProperty(PropMaxZoom, 7),
				},
				ShouldStore: scaleRankIsLessThan(7),
				GetKey:      BasicKey("urban_area_10m_1"),
			},
			{
				Opts: []naturalearth.Option{
					naturalearth.AddProperty(PropType, TypePropLanduse),
					naturalearth.AddProperty(PropMinZoom, 8),
					naturalearth.AddProperty(PropMaxZoom, 8),
				},
				ShouldStore: scaleRankIsLessThan(8),
				GetKey:      BasicKey("urban_area_10m_2"),
			},
		},
	}
}

func scaleRankIsLessThan(n float64) naturalearth.Filter {
	return func(feat *geojson.Feature) (bool, error) {
		var scaleRank float64
		if err := feat.Properties.GetType(PropScaleRank, &scaleRank); err != nil {
			return false, err
		}
		return scaleRank < n, nil
	}
}
