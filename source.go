package naturalearth

import (
	"fmt"
	"net/http"

	"github.com/gosuri/uiprogress"
	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/go-shapefile"
	"github.com/pkg/errors"
)

type Source struct {
	Name        string
	Label       string
	Schemas     []Schema
	numFeatures uint32
	numInserts  uint64
}

type Schema struct {
	ShouldStore      Filter
	GetKey           KeyGetter
	Opts             []Option
	conf             config
	bar              *uiprogress.Bar
	expectedInserts  uint32
	completedInserts uint32
}

type Filter func(*geojson.Feature) (bool, error)

type KeyGetter func(*geojson.Feature) (string, error)

type Store interface {
	Insert(*geojson.Feature, string) (string, error)
}

func (s *Source) Load(uri string, store Store) error {
	if s.Label == "" {
		s.Label = s.Name
	}

	zip, err := s.open(uri)
	if err != nil {
		return err
	}
	return s.load(zip, uri, store)
}

func (s *Source) NumFeatures() uint32 {
	return s.numFeatures
}

func (s *Source) NumInserts() uint64 {
	return s.numInserts
}

func (s *Source) open(uri string) (*shapefile.ZipScanner, error) {
	source := NewFile(uri, s.Label+": Downloading", http.DefaultClient)
	scanner, err := source.Open()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open source '%s'", uri)
	}
	return scanner, nil
}

func (s *Source) load(zip *shapefile.ZipScanner, uri string, store Store) error {
	fields := []string{}
	for i, schema := range s.Schemas {
		s.Schemas[i].conf = defaultConfig(uri)
		conf := &s.Schemas[i].conf

		for _, opt := range schema.Opts {
			opt(conf)
		}

		for field := range conf.oldNewProps {
			fields = append(fields, field)
		}
	}

	scanner := NewScanner(zip)
	if err := scanner.Scan(fields); err != nil {
		return err
	}

	info, err := zip.Info()
	if err != nil {
		return err
	}
	s.numFeatures = info.NumRecords

	for i := range s.Schemas {
		bar := uiprogress.AddBar(int(s.numFeatures))
		bar.PrependFunc(func(b *uiprogress.Bar) string {
			return s.Label + ": Applying   "
		})

		bar.AppendCompleted()

		insert := &s.Schemas[i]
		insert.expectedInserts = s.numFeatures
		bar.AppendFunc(func(b *uiprogress.Bar) string {
			return fmt.Sprintf("(%d / %d)", insert.completedInserts, insert.expectedInserts)
		})

		insert.bar = bar
	}

	for recNum := uint(1); ; recNum++ {
		rec := scanner.Record()
		if rec == nil {
			break
		}

		for i, schema := range s.Schemas {
			feat := rec.GeoJSONFeature(shapefile.RenameProperties(schema.conf.oldNewProps)).
				AddProperty(NumberPropertyName, recNum)

			ok, err := s.process(feat, store, schema.ShouldStore, schema.GetKey, schema.conf)
			if err != nil {
				return err
			}

			if ok {
				s.Schemas[i].completedInserts++
			} else {
				s.Schemas[i].expectedInserts--
			}
			schema.bar.Set(int(recNum))
		}
	}

	return nil
}

func (s *Source) process(feat *geojson.Feature, store Store, shouldStore Filter, getKey KeyGetter, conf config) (bool, error) {
	if shouldStore != nil {
		if ok, err := shouldStore(feat); err != nil {
			return false, errors.Wrap(err, "failed to check feature")
		} else if !ok {
			return false, nil
		}
	}

	props := conf.newProps
	if numProp, ok := feat.Properties.Get(NumberPropertyName); ok {
		props = append(props, *numProp)
	}

	for _, prop := range feat.Properties {
		for _, name := range conf.oldNewProps {
			if name == prop.Name {
				props = append(props, prop)
			}
		}
	}
	feat = feat.WithProperties(props...)

	key, err := getKey(feat)
	if err != nil {
		return false, errors.Wrap(err, "failed to get key")
	}

	if _, err := store.Insert(feat, key); err != nil {
		return false, err
	}
	s.numInserts++
	return true, nil
}
