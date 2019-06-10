package naturalearth_test

import (
	"testing"

	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"

	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	store := &MockStore{t: t}

	prog, err := naturalearth.Update("testdata/ne_110m_admin_0_countries.zip", "", store,
		naturalearth.RenameProperties(map[string]string{
			"NAME_EN": "name_en",
		}),
		naturalearth.AddProperties(
			geojson.Property{
				Name:  "key",
				Value: "value",
			},
		))

	require.NoError(t, err)
	require.NoError(t, prog.Error)
	require.Equal(t, 177, int(prog.Total))
	for range prog.Progress {
	}
	require.Equal(t, 177, int(store.num))
}

type MockStore struct {
	t   *testing.T
	num uint
}

func (s *MockStore) Insert(feat *geojson.Feature, _ string) (string, error) {
	name, ok := feat.Properties.Get("name_en")
	require.Truef(s.t, ok, "missing 'name_en'")
	s.t.Logf("name_en = '%s'", name)

	require.Contains(s.t, feat.Properties, geojson.Property{Name: "key", Value: "value"})

	s.num++
	return "", nil
}
