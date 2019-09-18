package naturalearth

import (
	"net/http"

	"github.com/gosuri/uiprogress"
	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/go-shapefile"
	"github.com/pkg/errors"
)

type Source struct {
	Name           string
	Label          string
	ShouldStore    Checker
	GetKey         KeyGetter
	Opts           []Option
	numFeatures    uint32
	storedFeatures uint32
}

type Checker func(*geojson.Feature) (bool, error)

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

func (s *Source) StoredFeatures() uint32 {
	return s.storedFeatures
}

func (s *Source) open(uri string) (*shapefile.ZipScanner, error) {
	source := NewFile(uri, s.Label+": Downloading", http.DefaultClient)
	scanner, err := source.Open()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open source '%s'", uri)
	}
	return scanner, nil
}

func (s *Source) process(feat *geojson.Feature, store Store, conf config) error {
	if s.ShouldStore != nil {
		if ok, err := s.ShouldStore(feat); err != nil {
			return errors.Wrap(err, "failed to check feature")
		} else if !ok {
			return nil
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

	key, err := s.GetKey(feat)
	if err != nil {
		return errors.Wrap(err, "failed to get key")
	}

	if _, err := store.Insert(feat, key); err != nil {
		return err
	}
	s.storedFeatures++
	return nil
}

func (s *Source) load(zip *shapefile.ZipScanner, uri string, store Store) error {
	conf := defaultConfig(uri)
	for _, opt := range s.Opts {
		opt(&conf)
	}

	scanner := NewScanner(zip)
	if err := scanner.Scan(conf.oldNewProps); err != nil {
		return err
	}

	info, err := zip.Info()
	if err != nil {
		return err
	}
	s.numFeatures = info.NumRecords

	bar := uiprogress.AddBar(int(info.NumRecords))
	bar.AppendCompleted()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return s.Label + ": Applying   "
	})

	for i := 1; ; i++ {
		feat := scanner.Feature()
		if feat == nil {
			break
		}

		if err := s.process(feat, store, conf); err != nil {
			return err
		}
		bar.Set(i)
	}
	return nil
}
