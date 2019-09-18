package naturalearth

import (
	"github.com/mercatormaps/go-geojson"
)

type Option func(*config)

func AddProperty(name string, value interface{}) Option {
	return func(c *config) {
		c.newProps = append(c.newProps, geojson.Property{Name: name, Value: value})
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
				Name:  "source_uri",
				Value: uri,
			},
		},
		oldNewProps: make(map[string]string),
	}
}
