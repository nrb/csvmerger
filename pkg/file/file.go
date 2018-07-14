package file

import (
	"bufio"
	"os"
	"strings"

	"github.com/nrb/csvmerger/pkg/types"
	"github.com/pkg/errors"
)

func LineToEntry(line string) (*types.Entry, error) {
	fields := strings.Split(line, ",")
	switch {
	case len(fields) > 3:
		// We have an extra comma, check for surrounding quotes
		var start, end int
		for i, f := range fields {
			if strings.HasPrefix(f, `"`) {
				start = i
			}
			if strings.HasSuffix(f, `"`) {
				end = i
			}
		}
		// Try to merge the quoted text into one field
		var newField string
		if start < end {
			newField = strings.Join(fields[start:end+1], ",")
		}
		// Insert the actual field into the fields, and drop any extras
		if newField != "" {
			fields[start] = newField
			fields = append(fields[:end], fields[end+1:]...)
		}
		// Our attempt at fixing the problem didn't work, error
		if len(fields) != 3 {
			return nil, errors.Errorf("Expected 3 fields, got %d for %s", len(fields), line)
		}
	case len(fields) < 3:
		return nil, errors.Errorf("Expected 3 fields, got %d for %s", len(fields), line)
	}
	return types.NewEntry(fields[0], fields[1], fields[2]), nil
}

func CSVToEntries(filePath string) ([]*types.Entry, error) {
	var entries []*types.Entry
	f, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't open file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Don't process empty lines
		t := scanner.Text()
		if t == "" {
			continue
		}
		e, err := LineToEntry(t)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		entries = append(entries, e)

	}
	if err = scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "Error reading file")
	}
	return entries, nil
}
