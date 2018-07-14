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

	if len(os.Args) < 3 {
		log.Fatalf("Need at least 2 files to merge")
		return
	}

	files := os.Args[2:]
	if len(files) < 2 {
		log.Fatalf("Need at least 2 files to merge")
	}

	var merged []*types.Entry
	redefs := make(map[string][]*types.Entry)

	for i := 0; i < len(files); i++ {
		// Load the entries
		es, err := file.CSVToEntries(files[i])
		if err != nil {
			log.Fatalf("Error with file %s: %s", files[i], err)
		}

		// Look for redefinitions
		for _, e := range es {
			rds, ok := entries.FindRedefinition(e, merged)
			if ok {
				key := fmt.Sprintf("%s:%s", files[i], e.ToString())
				redefs[key] = rds
			}
		}
		merged = entries.Merge(merged, es)
	}

	if len(redefs) > 0 {
		fmt.Println("Redefintions were found, can't merge")
		for key, vals := range redefs {
			fmt.Println(key)
			for _, v := range vals {
				fmt.Printf("\t%s\n", v.ToString())
			}
		}
		fmt.Println("Redefintions were found, can't merge")
		return
	}

	for _, m := range merged {
		fmt.Println(m.ToString())
	}

}
