package naturalearth

import (
	"net/http"

	"github.com/mercatormaps/go-geojson"

	"github.com/pkg/errors"
)

type Store interface {
	Insert(*geojson.Feature, string) (string, error)
}

type UpdateProgress struct {
	Total    uint32
	Progress chan uint32
	Error    error
}

func Update(uri, idSuffix string, store Store, opts ...Option) (*UpdateProgress, error) {
	src := NewSource(uri, http.DefaultClient)
	shp, err := src.Open()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open source '%s'", uri)
	}
	defer src.Close()

	conf := defaultConfig()
	for _, opt := range opts {
		opt(&conf)
	}

	s := NewScanner(shp)
	if err := s.Scan(conf.oldNewProps); err != nil {
		return nil, err
	}

	info, err := shp.Info()
	if err != nil {
		return nil, err
	}

	prog := UpdateProgress{
		Total:    info.NumRecords,
		Progress: make(chan uint32, info.NumRecords),
	}

	go func() {
		for i := 1; ; i++ {
			feat := s.Feature()
			if feat == nil {
				break
			}
			feat.Properties = append(feat.Properties, conf.newProps...)

			_, err := store.Insert(feat, idSuffix)
			if err != nil {
				prog.Error = err
				return
			}

			prog.Progress <- uint32(i)
		}
		close(prog.Progress)
	}()

	return &prog, s.Err()
}
