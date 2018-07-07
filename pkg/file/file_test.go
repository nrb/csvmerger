package file

import (
	"path/filepath"
	"testing"

	"github.com/nrb/csvmerger/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestLineToEntry(t *testing.T) {
	tests := []struct {
		name          string
		line          string
		expectedEntry *types.Entry
		expectedErr   bool
	}{
		{
			name:        "Line doesn't have enough fields",
			line:        "まち,city / town",
			expectedErr: true,
		},
		{
			name:          "Empty tags field works",
			line:          "まち,city / town,",
			expectedEntry: types.NewEntry("まち", "city / town", ""),
		},
		{
			name:          "Full line works",
			line:          "まち,city / town,1 2 3",
			expectedEntry: types.NewEntry("まち", "city / town", "1 2 3"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := LineToEntry(test.line)
			assert.Equal(t, test.expectedErr, err != nil)
			if test.expectedEntry != nil {
				assert.Equal(t, test.expectedEntry, actual)
			}
		})
	}
}

func TestCSVToEntries(t *testing.T) {
	tests := []struct {
		name            string
		fileName        string
		expectedEntries []*types.Entry
		expectedErr     bool
	}{
		{
			name:     "Short valid file",
			fileName: "validfile.csv",
			expectedEntries: []*types.Entry{
				types.NewEntry("まち", "city / town", "1 2 3"),
				types.NewEntry("うち", "house / home", "2 3"),
			},
		},
		{
			name:        "Partially valid file returns error",
			fileName:    "partialvalidfile.csv",
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fullFileName := filepath.Join("testdata", test.fileName)
			entries, err := CSVToEntries(fullFileName)
			t.Log(err)
			assert.Equal(t, test.expectedErr, err != nil)
			if test.expectedEntries != nil {
				assert.Equal(t, test.expectedEntries, entries)
			}
		})
	}
}
