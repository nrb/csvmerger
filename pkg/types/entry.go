package types

import (
	"fmt"

	"github.com/pkg/errors"
)

type Entry struct {
	Japanese string
	English  string
	Tags     *TagSet
}

func NewEntry(jpText, engText, tags string) *Entry {
	tagSet, _ := NewTagSet(tags)
	return &Entry{
		Japanese: jpText,
		English:  engText,
		Tags:     tagSet,
	}
}

func (e *Entry) ToString() string {
	return fmt.Sprintf("%s,%s,%s", e.Japanese, e.English, e.Tags.ToString())
}

// EntriesAreEqual compares the Japanese and English fields of an Entry for equality.
func EntriesAreEqual(e1, e2 *Entry) bool {
	return e1.Japanese == e2.Japanese && e1.English == e2.English

}

func (e *Entry) MergeTags(source *Entry) error {
	if !EntriesAreEqual(e, source) {
		return errors.New("Cannot merge unequal entries")
	}
	e.Tags.Insert(source.Tags.ToString())
	return nil
}
