package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/twpayne/go-geom"
)

// use this for auto-updating
// https://drive.google.com/uc?id=138PgJgUCv1Y10LRDsj3eUnOZh3cf_1v4

type StringFloat float64

func (ce CodexEntry) Coords() []float64 {
	fs := make([]float64, 3)
	x, _ := strconv.ParseFloat(ce.X, 64)
	fs[0] = x

	y, _ := strconv.ParseFloat(ce.Y, 64)
	fs[1] = y

	z, _ := strconv.ParseFloat(ce.Z, 64)
	fs[2] = z

	return fs
}

type CodexEntry struct {
	IndexID              interface{} `json:"index_id"`
	HudCategory          string      `json:"hud_category"`
	EnglishName          string      `json:"english_name"`
	CreatedAt            string      `json:"created_at"`
	ReportedAt           string      `json:"reported_at"`
	Cmdrname             string      `json:"cmdrName"`
	System               string      `json:"system"`
	X                    string      `json:"x"`
	Y                    string      `json:"y"`
	Z                    string      `json:"z"`
	Body                 string      `json:"body"`
	Latitude             string      `json:"latitude"`
	Longitude            string      `json:"longitude"`
	Entryid              int         `json:"entryid"`
	Name                 string      `json:"name"`
	NameLocalised        string      `json:"name_localised"`
	Category             string      `json:"category"`
	CategoryLocalised    string      `json:"category_localised"`
	SubCategory          string      `json:"sub_category"`
	SubCategoryLocalised string      `json:"sub_category_localised"`
	RegionName           string      `json:"region_name"`
	RegionNameLocalised  string      `json:"region_name_localised"`
	Distance             float64
}

func LoadEntries(fn string) ([]CodexEntry, error) {
	var ces []CodexEntry

	f, err := os.Open(fn)
	if err != nil {
		return ces, fmt.Errorf("couldn't open codex file: %w", err)
	}

	dec := json.NewDecoder(f)
	err = dec.Decode(&ces)
	if err != nil {
		return ces, fmt.Errorf("couldn't parse codex file: %w", err)
	}

	return ces, nil
}

func NewKeyedCodex(ces []CodexEntry) KeyedCodex {
	m := make(map[string][]CodexEntry)

	for _, e := range ces {
		m[e.Name] = append(m[e.Name], e)
	}

	return m
}

type KeyedCodex map[string][]CodexEntry

func (kc *KeyedCodex) Sort(system string) {
	s, err := GetSystem(system)
	if err != nil {
		log.Fatalf("error fetching reference system: %s", err)
	}
	log.Println("ref system:", s)

	refCoords := geom.Coord{s.Coords.X, s.Coords.Y, s.Coords.Z}
	for _, entries := range *kc {
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
	}

}
