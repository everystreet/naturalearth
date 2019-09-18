package data

type Name string

const (
	Boundaries110Name = "boundaries-110m"
	Boundaries50Name  = "boundaries-50m"
	Boundaries10Name  = "boundaries-10m"
	StateLines50Name  = "states-50m"
	StateLines10Name  = "states-10m"
)

func MaxNameLen() int {
	var n int
	for _, name := range []string{
		Boundaries110Name,
		Boundaries50Name,
		Boundaries10Name,
		StateLines50Name,
		StateLines10Name,
	} {
		if l := len(name); l > n {
			n = l
		}
	}
	return n
}
