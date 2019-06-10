package naturalearth

import (
	"fmt"
	"sync"

	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/go-shapefile"
)

type Scanner struct {
	shp Shapefile

	scanOnce   sync.Once
	featuresCh chan *geojson.Feature

	errOnce sync.Once
	err     error
}

type Shapefile interface {
	AddOptions(...shapefile.Option)
	Info() (*shapefile.Info, error)
	Scan() error
	Record() *shapefile.Record
	Err() error
}

func NewScanner(shp Shapefile) *Scanner {
	return &Scanner{
		shp:        shp,
		featuresCh: make(chan *geojson.Feature),
	}
}

func (s *Scanner) Scan(fieldProps map[string]string) error {
	info, err := s.shp.Info()
	if err != nil {
		return err
	}

	fields := make([]string, len(fieldProps))
	i := 0
	for field := range fieldProps {
		if !info.Fields.Exists(field) {
			return fmt.Errorf("field '%s' does not exist in shapefile", field)
		}
		fields[i] = field
	}
	s.shp.AddOptions(shapefile.FilterFields(fields...))

	s.scanOnce.Do(func() {
		if err = s.shp.Scan(); err != nil {
			return
		}

		go func() {
			defer func() {
				if err := s.shp.Err(); err != nil {
					fmt.Println(err)
					s.setErr(err)
				}
				close(s.featuresCh)
			}()

			for {
				rec := s.shp.Record()
				if rec == nil {
					break
				}

				feat := rec.GeoJSONFeature(shapefile.RenameProperties(fieldProps))
				s.featuresCh <- feat
			}
		}()
	})

	return err
}

func (s *Scanner) Feature() *geojson.Feature {
	feat, ok := <-s.featuresCh
	if !ok {
		return nil
	}
	return feat
}

func (s *Scanner) Err() error {
	return s.err
}

func (s *Scanner) setErr(err error) {
	s.errOnce.Do(func() {
		s.err = err
	})
}
