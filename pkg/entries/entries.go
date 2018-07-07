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
