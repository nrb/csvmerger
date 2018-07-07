package types

import (
	"fmt"
	"sort"
	"strings"
)

// TagSet is a set of strings representing unique tags. Tags are not
// guarunteed to be sorted unless retrieved via the Sort method.
type TagSet struct {
	Tags map[string]bool
}

// NewTagSet returns a TagSet with the provided tags.
// If the provided string has spaces, tags will be created by splitting on spaces.
// The emptry string will result in an empty TagSet.
func NewTagSet(tags string) (*TagSet, error) {
	actualTags := strings.Split(tags, " ")
	tagMap := make(map[string]bool)
	for _, tag := range actualTags {
		if tag != "" {
			tagMap[tag] = true
		}
	}
	return &TagSet{Tags: tagMap}, nil
}

// Insert inserts one or more tags into the TagSet.
// Strings with spaces will be split on the spaces.
func (ts *TagSet) Insert(tags string) {
	actualTags := strings.Split(tags, " ")
	for _, tag := range actualTags {
		if tag != "" {
			ts.Tags[tag] = true
		}
	}
}

// ToString returns the tags in a TagSet as a single, space-separated string.
func (ts *TagSet) ToString() string {
	var b strings.Builder
	entries := ts.Sort()
	for _, entry := range entries {
		fmt.Fprintf(&b, "%s ", entry)
	}
	return strings.TrimSpace(b.String())
}

// Sort returns the tags in a TagSet as a sorted slice of strings.
// Sorting is in string order, not numeric.
func (ts *TagSet) Sort() []string {
	var sorted []string
	for key, _ := range ts.Tags {
		sorted = append(sorted, key)
	}
	sort.Strings(sorted)
	return sorted
}
