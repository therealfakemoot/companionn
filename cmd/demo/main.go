package main

import (
	// "fmt"
	"flag"
	"log"
	"os"

	// "github.com/davecgh/go-spew/spew"
	"github.com/therealfakemoot/challenge-helper"
)

func main() {
	var (
		name    string
		origin  string
		count   int
		nearest bool
	)

	flag.StringVar(&name, "name", "", "Commander name")
	flag.StringVar(&origin, "origin", "Sol", "Reference system")
	flag.BoolVar(&nearest, "nearest", false, "Print only the nearest codex entry")
	flag.IntVar(&count, "count", 10, "Maximum entries to display.")

	flag.Parse()

	if name == "" {
		log.Fatal("Please supply a Commander name.")
	}

	entries, err := helper.LoadEntries("json/codex.json")
	if err != nil {
		log.Fatalf("error loading codex entries: %s", err)
	}

	keyedCodex := helper.NewKeyedCodex(entries)
	cache := make(map[string]helper.EDSMSystem)
	keyedCodex.CalculateDistances(origin, cache)
	keyedCodex.Sort()

	keyedCodex.Render(os.Stdout, count)

}
