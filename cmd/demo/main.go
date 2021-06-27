package main

import (
	"sort"
	// "fmt"
	"flag"
	"log"

	"github.com/therealfakemoot/challenge-helper"
	"github.com/twpayne/go-geom"
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

	challengeMap := helper.KeyedCodex(challenges)
	bact := challengeMap["$Codex_Ent_Bacterial_06_B_Name;"]
	log.Printf("%#v\n", bact)

	s, err := helper.GetSystem(origin)
	if err != nil {
		log.Fatalf("error fetching reference system: %s", err)
	}
	log.Println("ref system:", s)

	refCoords := geom.Coord{s.Coords.X, s.Coords.Y, s.Coords.Z}
	// nearest := make(map[string]helper.CodexEntry)
	log.Println("refCoords:", refCoords)

	// challengeMap is a map of slices. i should be able to use a clever
	// slice.Sort func here

	for codexKey, entries := range challengeMap {
		sortfunc := func(i, j int) bool {
			iLine := geom.NewLineString(geom.XYZ)
			iLine.SetCoords([]geom.Coord{refCoords, entries[i].Coords()})
			entries[i].Distance = iLine.Length()

			jLine := geom.NewLineString(geom.XYZ)
			jLine.SetCoords([]geom.Coord{refCoords, entries[j].Coords()})
			entries[j].Distance = jLine.Length()

			return iLine.Length() < jLine.Length()
		}
		sort.Slice(entries, sortfunc)
		log.Printf("%#v\n", entries)
		log.Println(codexKey)
		log.Println(len(entries))
		return

		/*
			for system, v := range entries {
				c, err := v.Coords()
				if err != nil {
					log.Printf("error getting system coords for %s: %s", v.System, err)
				}
				log.Printf("%s| %#v\n", v.System, c)

			}
		*/

	}

}
