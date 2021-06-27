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
		name   string
		origin string
	)

	flag.StringVar(&name, "name", "", "Commander name")
	flag.StringVar(&origin, "origin", "Sol", "Reference system")

	flag.Parse()

	if name == "" {
		log.Fatal("Please supply a Commander name.")
	}

	cm, err := helper.FetchChallengeMap(name)
	if err != nil {
		log.Fatalf("error fetching challenge report: %s", err)
	}

	// fmt.Printf("%#v\n", cr["Aleoids"].Remaining())
	// fmt.Printf("%d\n", len(cr["Aleoids"].Remaining()))
	hm, err := helper.LoadHierarchy("json/hierarchy.json")
	if err != nil {
		log.Fatalf("error loading hierarchy data: %s", err)
	}

	f := helper.FlattenHierarchy(hm)
	// log.Printf("%#v\n", f)
	// for k, v := range f {
	// log.Printf("%s: %#v\n", k, v)
	// }
	// log.Printf("%#v\n", hm["Biology"]["Aleoids"])

	var codexTodo []string
	for _, v := range cm {
		for _, challenge := range v.Remaining() {
			codexTodo = append(codexTodo, challenge)
		}
	}

	var codexKeys []string
	for _, c := range codexTodo {
		codexKeys = append(codexKeys, f[c].Name)
	}

	challenges, err := helper.LoadEntries("json/codex.json")
	if err != nil {
		log.Fatalf("error loading codex entries: %s", err)
	}

	challengeMap := helper.NewKeyedCodex(challenges)
	challengeMap.Sort(origin)
	// bact := challengeMap["$Codex_Ent_Bacterial_06_B_Name;"]

}
