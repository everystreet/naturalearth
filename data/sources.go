package data

import "github.com/mercatormaps/naturalearth"

type Name string

const (
	BoundaryLines110Name = "boundary-lines-110m"
	BoundaryLines50Name  = "boundary-lines-50m"
	BoundaryLines10Name  = "boundary-lines-10m"
	StateLines50Name     = "state-lines-50m"
	StateLines10Name     = "state-lines-10m"
	Glaciers110Name      = "glaciers-110m"
	Glaciers50Name       = "glaciers-50m"
	Glaciers10Name       = "glaciers-10m"
)

var sources = map[Name]func() *naturalearth.Source{
	BoundaryLines110Name: BoundaryLines110,
	BoundaryLines50Name:  BoundaryLines50,
	BoundaryLines10Name:  BoundaryLines10,
	StateLines50Name:     StateLines50,
	StateLines10Name:     StateLines10,
	Glaciers110Name:      Glaciers110,
	Glaciers50Name:       Glaciers50,
	Glaciers10Name:       Glaciers10,
}

func Source(name string) (*naturalearth.Source, bool) {
	f, ok := sources[Name(name)]
	if !ok {
		return nil, false
	}
	return f(), true
}

func MaxNameLen() int {
	var n int
	for name := range sources {
		if l := len(name); l > n {
			n = l
		}
	}
	return n
}
