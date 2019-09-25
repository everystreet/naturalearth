package data

import "github.com/mercatormaps/naturalearth"

var Glaciers110 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: Glaciers110Name,
		Schemas: []naturalearth.Schema{
			{
				Opts: []naturalearth.Option{
					naturalearth.AddProperty(PropType, TypePropLandcover),
					naturalearth.AddProperty(PropMinZoom, 0),
					naturalearth.AddProperty(PropMaxZoom, 1),
					naturalearth.AddProperty(PropLandcoverClass, LandcoverClassPropIce),
					naturalearth.AddProperty(PropLandcoverSubclass, LandcoverSubclassPropGlacier),
				},
				GetKey: BasicKey("glacier_110m"),
			},
		},
	}
}

var Glaciers50 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: Glaciers50Name,
		Schemas: []naturalearth.Schema{
			{
				Opts: []naturalearth.Option{
					naturalearth.AddProperty(PropType, TypePropLandcover),
					naturalearth.AddProperty(PropMinZoom, 2),
					naturalearth.AddProperty(PropMaxZoom, 4),
					naturalearth.AddProperty(PropLandcoverClass, LandcoverClassPropIce),
					naturalearth.AddProperty(PropLandcoverSubclass, LandcoverSubclassPropGlacier),
				},
				GetKey: BasicKey("glacier_50m"),
			},
		},
	}
}

var Glaciers10 = func() *naturalearth.Source {
	return &naturalearth.Source{
		Name: Glaciers10Name,
		Schemas: []naturalearth.Schema{
			{
				Opts: []naturalearth.Option{
					naturalearth.AddProperty(PropType, TypePropLandcover),
					naturalearth.AddProperty(PropMinZoom, 5),
					naturalearth.AddProperty(PropMaxZoom, 6),
					naturalearth.AddProperty(PropLandcoverClass, LandcoverClassPropIce),
					naturalearth.AddProperty(PropLandcoverSubclass, LandcoverSubclassPropGlacier),
				},
				GetKey: BasicKey("glacier_10m"),
			},
		},
	}
}
