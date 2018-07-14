package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nrb/csvmerger/pkg/entries"
	"github.com/nrb/csvmerger/pkg/file"
	"github.com/nrb/csvmerger/pkg/types"
)

func main() {

	files := os.Args[1:]
	if len(files) < 2 {
		log.Fatalf("Need at least 2 files to merge")
	}

	var merged []*types.Entry
	for i := 0; i < len(files); i++ {
		e, err := file.CSVToEntries(files[i])
		if err != nil {
			log.Fatalf("Error with file %s: %s", files[i], err)
		}
		merged = entries.Merge(merged, e)
	}

	for _, m := range merged {
		fmt.Println(m.ToString())
	}
}
