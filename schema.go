package naturalearth

import (
	"github.com/mercatormaps/go-geojson"
)

type Schema func(geojson.Feature, *Meta) (string, error)

type Meta struct {
	props []geojson.Property
}

func (m *Meta) AddProperty(name string, value interface{}) {
	m.props = append(m.props, geojson.Property{Name: name, Value: value})
}
