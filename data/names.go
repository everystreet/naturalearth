package data

type Name string

const (
	BoundaryLines110Name = "boundary-lines-110m"
	BoundaryLines50Name  = "boundary-lines-50m"
	BoundaryLines10Name  = "boundary-lines-10m"
	StateLines50Name     = "state-lines-50m"
	StateLines10Name     = "state-lines-10m"
	Glaciers110Name      = "glaciers-110m"
)

func MaxNameLen() int {
	var n int
	for _, name := range []string{
		BoundaryLines110Name,
		BoundaryLines50Name,
		BoundaryLines10Name,
		StateLines50Name,
		StateLines10Name,
		Glaciers110Name,
	} {
		if l := len(name); l > n {
			n = l
		}
	}
	return n
}
