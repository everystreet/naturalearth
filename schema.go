package naturalearth

import (
	"github.com/everystreet/go-geojson/v2"
)

type Schema func(geojson.Feature, *Meta) (string, error)

type Meta struct {
	props []geojson.Property
}

func (m *Meta) AddProperty(name string, value interface{}) {
	m.props = append(m.props, geojson.Property{Name: name, Value: value})
}
