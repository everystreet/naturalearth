package naturalearth

import (
	"fmt"
	"sync"

	"github.com/mercatormaps/go-shapefile"
)

type Scanner struct {
	shp Shapefile

	scanOnce  sync.Once
	recordsCh chan *shapefile.Record

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

const NumberPropertyName = "number"

func NewScanner(shp Shapefile) *Scanner {
	return &Scanner{
		shp:       shp,
		recordsCh: make(chan *shapefile.Record),
	}
}

func (s *Scanner) Scan() error {
	var err error
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
				close(s.recordsCh)
			}()

			for i := uint(1); ; i++ {
				rec := s.shp.Record()
				if rec == nil {
					break
				}
				s.recordsCh <- rec
			}
		}()
	})

	return err
}

func (s *Scanner) Record() *shapefile.Record {
	rec, ok := <-s.recordsCh
	if !ok {
		return nil
	}
	return rec
}

func (s *Scanner) Err() error {
	return s.err
}

func (s *Scanner) setErr(err error) {
	s.errOnce.Do(func() {
		s.err = err
	})
}
