package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTagSet(t *testing.T) {
	tests := []struct {
		name           string
		tags           string
		expectedTagSet *TagSet
		expectedErr    bool
	}{
		{
			name:           "Tags are inserted into a tagset",
			tags:           "1",
			expectedTagSet: &TagSet{Tags: map[string]bool{"1": true}},
		},
		{
			name:           "Duplicate tags are not inserted",
			tags:           "1 1",
			expectedTagSet: &TagSet{Tags: map[string]bool{"1": true}},
		},
		{
			name:           "Empty strings result in initialized but empty map",
			tags:           "",
			expectedTagSet: &TagSet{Tags: map[string]bool{}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, err := NewTagSet(test.tags)
			assert.Equal(t, test.expectedErr, err != nil)
			assert.Equal(t, test.expectedTagSet, ts)
		})
	}
}

func TestTagSetInsert(t *testing.T) {
	tests := []struct {
		name           string
		insertList     []string
		expectedTagSet *TagSet
	}{
		{
			name:           "Inserting a single tag",
			insertList:     []string{"1"},
			expectedTagSet: &TagSet{Tags: map[string]bool{"1": true}},
		},
		{
			name:           "Inserting a string with spaces should result in as many tags",
			insertList:     []string{"1 2"},
			expectedTagSet: &TagSet{Tags: map[string]bool{"1": true, "2": true}},
		},
		{
			name:           "Inserting a string twice doesn't duplicate",
			insertList:     []string{"1", "1"},
			expectedTagSet: &TagSet{Tags: map[string]bool{"1": true}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, err := NewTagSet("")
			require.NoError(t, err)
			for _, entry := range test.insertList {
				ts.Insert(entry)
			}
			assert.Equal(t, test.expectedTagSet, ts)
		})
	}
}

func TestTagSetToString(t *testing.T) {
	tests := []struct {
		name        string
		tagSet      *TagSet
		expectedStr string
	}{
		{
			name:        "Single element",
			tagSet:      &TagSet{Tags: map[string]bool{"1": true}},
			expectedStr: "1",
		},
		{
			name:        "Multiple elements",
			tagSet:      &TagSet{Tags: map[string]bool{"1": true, "2": true}},
			expectedStr: "1 2",
		},
		{
			name:        "String output is sorted",
			tagSet:      &TagSet{Tags: map[string]bool{"22": true, "12": true, "1": true}},
			expectedStr: "1 12 22",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualStr := test.tagSet.ToString()
			assert.Equal(t, test.expectedStr, actualStr)
		})
	}
}

func TestTagSetSort(t *testing.T) {
	tests := []struct {
		name          string
		tags          string
		expectedOrder []string
	}{
		{
			name:          "Sorting works when already in order",
			tags:          "1 2 5 20",
			expectedOrder: []string{"1", "2", "20", "5"},
		},
		{
			name:          "Sorting works when not in order",
			tags:          "20 2 1 5",
			expectedOrder: []string{"1", "2", "20", "5"},
		},
		{
			name:          "Sorting works with a non-numeric value",
			tags:          "verb 1 2 10",
			expectedOrder: []string{"1", "10", "2", "verb"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, err := NewTagSet(test.tags)
			require.NoError(t, err)
			actual := ts.Sort()
			assert.Equal(t, test.expectedOrder, actual)
		})
	}
}
