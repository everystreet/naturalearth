package data

type Name string

const (
	Boundaries110Name = "boundaries-110"
	Boundaries50Name  = "boundaries-50"
	StateLines50Name  = "states-50"
)

func MaxNameLen() int {
	var n int
	for _, name := range []string{
		Boundaries110Name,
		Boundaries50Name,
		StateLines50Name,
	} {
		if l := len(name); l > n {
			n = l
		}
	}
	return n
}
