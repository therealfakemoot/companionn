package main

import (
	// "fmt"
	"flag"
	"log"

	"github.com/davecgh/go-spew/spew"
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

	cm, err := helper.FetchChallengeMap(name)
	if err != nil {
		log.Fatalf("error fetching challenge report: %s", err)
	}

	hm, err := helper.LoadHierarchy("json/hierarchy.json")
	if err != nil {
		log.Fatalf("error loading hierarchy data: %s", err)
	}

	humanToCodex := helper.FlattenHierarchy(hm)
	codexToHuman := make(map[string]string)
	for k, v := range humanToCodex {
		codexToHuman[v.Name] = k
	}

	var codexTodo []string
	for _, v := range cm {
		for _, challenge := range v.Remaining() {
			codexTodo = append(codexTodo, challenge)
		}
	}

	challenges, err := helper.LoadEntries("json/codex.json")
	if err != nil {
		log.Fatalf("error loading codex entries: %s", err)
	}

	challengeMap := helper.NewKeyedCodex(challenges)
	challengeMap.Sort(origin)
	// bact := challengeMap["$Codex_Ent_Bacterial_06_B_Name;"]

}
