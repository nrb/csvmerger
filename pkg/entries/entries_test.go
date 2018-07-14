package entries

import (
	"testing"

	"github.com/nrb/csvmerger/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name           string
		needle         *types.Entry
		haystack       []*types.Entry
		expectedFound  bool
		expectedTarget *types.Entry
	}{
		{
			name:   "Needle is in haystack, with matching tags",
			needle: types.NewEntry("まち", "city / town", "1"),
			haystack: []*types.Entry{
				types.NewEntry("まち", "city / town", "1"),
			},
			expectedFound:  true,
			expectedTarget: types.NewEntry("まち", "city / town", "1"),
		},
		{
			name:   "Needle is in haystack, without matching tags",
			needle: types.NewEntry("まち", "city / town", "1"),
			haystack: []*types.Entry{
				types.NewEntry("まち", "city / town", "1 2 3"),
			},
			expectedFound:  true,
			expectedTarget: types.NewEntry("まち", "city / town", "1 2 3"),
		},
		{
			name:          "Needle is not in empty haystack",
			needle:        types.NewEntry("まち", "city / town", "1"),
			haystack:      []*types.Entry{},
			expectedFound: false,
		},
		{
			name:   "Needle is not in haystack",
			needle: types.NewEntry("まち", "city / town", "1"),
			haystack: []*types.Entry{
				types.NewEntry("うち", "house / home", "1 2 3"),
			},
			expectedFound: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			target, found := Find(test.needle, test.haystack)
			assert.Equal(t, test.expectedFound, found)
			if test.expectedTarget != nil {
				assert.Equal(t, test.expectedTarget, target)
			}
		})
	}
}

func TestMergeLists(t *testing.T) {
	original := []*types.Entry{
		types.NewEntry("うち", "house / home", "1"),
		types.NewEntry("まち", "city / town", "2 3"),
	}
	new := []*types.Entry{
		types.NewEntry("うち", "house / home", "17 2"),
		types.NewEntry("まち", "city / town", "3"),
		types.NewEntry("じんじゃ", "shrine", "4"),
	}

	expected := []*types.Entry{
		types.NewEntry("うち", "house / home", "1 17 2"),
		types.NewEntry("まち", "city / town", "2 3"),
		types.NewEntry("じんじゃ", "shrine", "4"),
	}

	actual := Merge(original, new)
	assert.Equal(t, expected, actual)
}

func TestFindRedefinition(t *testing.T) {
	tests := []struct {
		name          string
		needle        *types.Entry
		haystack      []*types.Entry
		expectedSlice []*types.Entry
		expectedFound bool
	}{
		{
			name:     "Same Japanese but different English is a redefintion",
			needle:   types.NewEntry("まち", "city", "1"),
			haystack: []*types.Entry{types.NewEntry("まち", "town", "1")},
			expectedSlice: []*types.Entry{
				types.NewEntry("まち", "town", "1"),
			},
			expectedFound: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, ok := FindRedefinition(test.needle, test.haystack)
			assert.Equal(t, test.expectedFound, ok)
			assert.Equal(t, test.expectedSlice, actual)
		})
	}
}
