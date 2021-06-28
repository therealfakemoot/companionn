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
		nearest bool
	)

	flag.StringVar(&name, "name", "", "Commander name")
	flag.StringVar(&origin, "origin", "Sol", "Reference system")
	flag.BoolVar(&nearest, "nearest", false, "Print only the nearest codex entry")

	flag.Parse()

	if name == "" {
		log.Fatal("Please supply a Commander name.")
	}

	challenges, err := helper.LoadEntries("json/codex.json")
	if err != nil {
		log.Fatalf("error loading codex entries: %s", err)
	}

	challengeMap := helper.NewKeyedCodex(challenges)
	challengeMap.Sort(origin)
	// bact := challengeMap["$Codex_Ent_Bacterial_06_B_Name;"]

	challengeMap.Render(os.Stdout)

}
