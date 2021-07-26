package main

import (
	"log"

	"github.com/therealfakemoot/companionn"
)

func main() {
	ces, err := companionn.LoadEntries("json/codex.json")
	if err != nil {
		log.Fatalf("error loading entries: %s\n", err)
	}

	m := make(map[string][]companionn.CodexEntry)
	for _, entry := range ces {
		cat, ok := m[entry.Name]
		if !ok {
			m[entry.Name] = make([]companionn.CodexEntry, 0)
		}
		m[entry.Name] = append(cat, entry)
	}
	// log.Printf("codex count: %d\n", len(ces))

	// for cat, entries := range m {
	// log.Printf("%q: %d\n", cat, len(entries))
	// }
	hmap, err := companionn.LoadHierarchy("json/hierarchy.json")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", hmap["Anomaly"]["E-Type Anomaly"])
}
