package naturalearth

import (
	"fmt"
	"net/http"

	"github.com/everystreet/go-geojson/v2"
	"github.com/everystreet/go-shapefile"
	"github.com/gosuri/uiprogress"
)

type Source struct {
	Name        string
	Label       string
	Schemas     []Schema
	numFeatures uint32
	numInserts  uint64
}

type Store interface {
	Insert(*geojson.Feature, string) error
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
		return nil, fmt.Errorf("failed to open source '%s': %w", uri, err)
	}
	return scanner, nil
}

func (s *Source) load(zip *shapefile.ZipScanner, uri string, store Store) error {
	scanner := NewScanner(zip)
	if err := scanner.Scan(); err != nil {
		return err
	}

	info, err := zip.Info()
	if err != nil {
		return err
	}
	s.numFeatures = info.NumRecords

	bars := make([]struct {
		bar              *uiprogress.Bar
		expectedInserts  uint32
		completedInserts uint32
	}, len(s.Schemas))

	for i := range s.Schemas {
		bar := uiprogress.AddBar(int(s.numFeatures))
		bar.PrependFunc(func(b *uiprogress.Bar) string {
			return s.Label + ": Applying   "
		})

		bar.AppendCompleted()
		bars[i].expectedInserts = s.numFeatures
		bar.AppendFunc(func(b *uiprogress.Bar) string {
			return fmt.Sprintf("(%d / %d)", bars[i].completedInserts, bars[i].expectedInserts)
		})

		bars[i].bar = bar
	}

	for recNum := uint(1); ; recNum++ {
		rec := scanner.Record()
		if rec == nil {
			break
		}

		for i, schema := range s.Schemas {
			if err := func() error {
				defer bars[i].bar.Set(int(recNum))

				meta := Meta{}
				feat := rec.GeoJSONFeature().AddProperty(NumberPropertyName, recNum)
				key, err := schema(*feat, &meta)
				if err != nil {
					return err
				}

				if key == "" {
					bars[i].expectedInserts--
					return nil
				} else if err := store.Insert(feat.WithProperties(meta.props...), key); err != nil {
					return err
				}

				bars[i].completedInserts++
				s.numInserts++
				return nil
			}(); err != nil {
				return err
			}
		}
	}

	return nil
}
