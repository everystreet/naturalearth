package naturalearth

import (
	"github.com/mercatormaps/go-geojson"
)

type Option func(*config)

func AddProperties(props ...geojson.Property) Option {
	return func(c *config) {
		c.newProps = append(c.newProps, props...)
	}
}

func AddProperty(prop geojson.Property) Option {
	return func(c *config) {
		c.newProps = append(c.newProps, prop)
	}
}

func RenameProperty(old, new string) Option {
	return func(c *config) {
		c.oldNewProps[old] = new
	}
}

type config struct {
	newProps    []geojson.Property
	oldNewProps map[string]string
}

func defaultConfig(uri string) config {
	return config{
		newProps: []geojson.Property{
			{
				Name:  "source",
				Value: "naturalearth",
			},
			{
				Name:  "source",
				Value: uri,
			},
		},
		oldNewProps: make(map[string]string),
	}
}
