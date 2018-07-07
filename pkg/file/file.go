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
	if len(fields) != 3 {
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
