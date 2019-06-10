package naturalearth

import (
	"github.com/mercatormaps/go-geojson"
)

type Option func(*config)

func AddProperties(props ...geojson.Property) Option {
	return func(c *config) {
		c.newProps = props
	}
}

func RenameProperties(oldNew map[string]string) Option {
	return func(c *config) {
		c.oldNewProps = oldNew
	}
}

type config struct {
	newProps    []geojson.Property
	oldNewProps map[string]string
}

func defaultConfig() config {
	return config{}
}
