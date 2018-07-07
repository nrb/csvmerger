package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const machi = "まち"

func TestCreatingEntry(t *testing.T) {
	tests := []struct {
		name          string
		jpText        string
		engText       string
		tags          string
		expectedEntry *Entry
	}{
		{
			name:    "Empty entry",
			jpText:  "",
			engText: "",
			tags:    "",
			expectedEntry: &Entry{Japanese: "",
				English: "",
				Tags:    &TagSet{Tags: map[string]bool{}},
			},
		},
		{
			name:    "No tags",
			jpText:  machi,
			engText: "city / town",
			tags:    "",
			expectedEntry: &Entry{Japanese: machi,
				English: "city / town",
				Tags:    &TagSet{Tags: map[string]bool{}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := NewEntry(test.jpText, test.engText, test.tags)
			assert.Equal(t, test.expectedEntry, e)
		})
	}
}

func TestEntryToString(t *testing.T) {
	tests := []struct {
		name  string
		entry *Entry
	}{
		{
			name:  "No tags",
			entry: NewEntry(machi, "city / town", ""),
		},
		{
			name:  "Single tag",
			entry: NewEntry(machi, "city / town", "1"),
		},
		{
			name:  "Multiple tags",
			entry: NewEntry(machi, "city / town", "1 12 2"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expectedStr := fmt.Sprintf("%s,%s,%s", test.entry.Japanese, test.entry.English, test.entry.Tags.ToString())
			str := test.entry.ToString()
			assert.Equal(t, expectedStr, str)
		})
	}
}

func TestEntriesAreEqual(t *testing.T) {
	tests := []struct {
		name          string
		entry1        *Entry
		entry2        *Entry
		expectedEqual bool
	}{
		{
			name:          "Empty entries are equal",
			entry1:        NewEntry("", "", ""),
			entry2:        NewEntry("", "", ""),
			expectedEqual: true,
		},
		{
			name:          "Differing tags are still equal",
			entry1:        NewEntry(machi, "city / town", "1"),
			entry2:        NewEntry(machi, "city / town", "1 2"),
			expectedEqual: true,
		},
		{
			name:          "Japanese not equal returns false",
			entry1:        NewEntry(machi, "city / town", ""),
			entry2:        NewEntry("bogus", "city / town", ""),
			expectedEqual: false,
		},
		{
			name:          "English not equal returns false",
			entry1:        NewEntry(machi, "city / town", ""),
			entry2:        NewEntry(machi, "bogus", ""),
			expectedEqual: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, EntriesAreEqual(test.entry1, test.entry2), test.expectedEqual)
		})
	}
}

func TestEntryMergeTags(t *testing.T) {
	tests := []struct {
		name          string
		baseEntry     *Entry
		layeredEntry  *Entry
		expectedEntry *Entry
		expectedErr   bool
	}{
		{
			name:          "Merging the same values results in the same entry",
			baseEntry:     NewEntry(machi, "city / town", "1"),
			layeredEntry:  NewEntry(machi, "city / town", "1"),
			expectedEntry: NewEntry(machi, "city / town", "1"),
		},
		{
			name:          "Merging differing tags results in the union of all tags",
			baseEntry:     NewEntry(machi, "city / town", "1 2 3 4"),
			layeredEntry:  NewEntry(machi, "city / town", "20"),
			expectedEntry: NewEntry(machi, "city / town", "1 2 20 3 4"),
		},
		{
			name:          "Overlapping tags aren't duplicated",
			baseEntry:     NewEntry(machi, "city / town", "1 2 3 4"),
			layeredEntry:  NewEntry(machi, "city / town", "1 3"),
			expectedEntry: NewEntry(machi, "city / town", "1 2 3 4"),
		},
		{
			name:         "If entries aren't equal, return an error",
			baseEntry:    NewEntry(machi, "city / town", "1"),
			layeredEntry: NewEntry("bogus", "city / town", "1"),
			expectedErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.baseEntry.MergeTags(test.layeredEntry)
			assert.Equal(t, test.expectedErr, err != nil)
			if test.expectedEntry != nil {
				assert.Equal(t, test.baseEntry, test.expectedEntry)
			}
		})
	}
}
