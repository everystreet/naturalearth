package data

// PropType property is present on all features.
const PropType = "type"

// TypePropBoundary is the type property value for boundaries.
const TypePropBoundary = "boundary"

// Landcover type property and associated properties.
const (
	TypePropLandcover = "landcover"

	PropLandcoverClass    = "class"
	LandcoverClassPropIce = "ice"

	PropLandcoverSubclass        = "subclass"
	LandcoverSubclassPropGlacier = "glacier"
)

// Landuse type property and associated properties.
const (
	TypePropLanduse = "landuse"

	PropLanduseClass             = "class"
	LanduserClassPropResidential = "residential"
)

// Zoom level properties.
const (
	PropMinZoom = "min_zoom"
	PropMaxZoom = "max_zoom"
)

// featurecla property is found in some shapefiles.
const (
	PropFeatureClass           = "featurecla"
	FeatureClassPropLeaseLimit = "Lease limit"

	PropScaleRank = "scalerank"
)
