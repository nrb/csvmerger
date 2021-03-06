package entries

import (
	"github.com/nrb/csvmerger/pkg/types"
)

// Find will find an equivalent entry (the needle) in a slice of entries (the haystack).
// If the entry is found, a pointer to it is returned.
// If the entry is not found, the pointer is nil.
func Find(needle *types.Entry, haystack []*types.Entry) (*types.Entry, bool) {
	var found bool
	var target *types.Entry
	for _, e := range haystack {
		if types.EntriesAreEqual(e, needle) {
			found = true
			target = e
			break
		}
	}
	return target, found
}

// Update merges tags from a source Entry into a target entry.
func Update(target, source *types.Entry) error {
	return target.MergeTags(source)
}

// Merge combines two slices of entries. If an entry from new
// is not in original, it will be added. If an entry is in both,
// the tags from both will be merged.
func Merge(original, new []*types.Entry) []*types.Entry {
	for _, e := range new {
		o, ok := Find(e, original)
		if ok {
			Update(o, e)
		} else {
			original = append(original, e)
		}
	}
	return original
}

// FindRedefinition searches a slice of *types.Entry for entries that define
// different English terms for Japanese terms, and vice versa.
// Any redefinitions found are returned as a slice.
func FindRedefinition(needle *types.Entry, haystack []*types.Entry) ([]*types.Entry, bool) {
	redefs := []*types.Entry{}
	var redefined bool
	for _, e := range haystack {
		if types.EntriesRedefined(needle, e) {
			redefined = true
			redefs = append(redefs, e)
		}
	}
	return redefs, redefined
}
