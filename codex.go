package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/twpayne/go-geom"
)

// use this for auto-updating
// https://drive.google.com/uc?id=138PgJgUCv1Y10LRDsj3eUnOZh3cf_1v4

type CodexEntry struct {
	IndexID              interface{} `json:"index_id"`
	HudCategory          string      `json:"hud_category"`
	EnglishName          string      `json:"english_name"`
	CreatedAt            string      `json:"created_at"`
	ReportedAt           string      `json:"reported_at"`
	Cmdrname             string      `json:"cmdrName"`
	System               string      `json:"system"`
	X                    json.Number `json:"x"`
	Y                    json.Number `json:"y"`
	Z                    json.Number `json:"z"`
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

func (kc KeyedCodex) Render(w io.Writer) error {
	_, codexToHuman, err := NameMapping()
	if err != nil {
		return fmt.Errorf("error loading name-mapping: %w", err)
	}

	closest := make([]CodexEntry, len(kc))
	i := 0
	for _, v := range kc {
		closest[i] = v[0]
		i++
		// closest = append(closest, v[0])
	}
	sortfunc := func(i, j int) bool {
		return closest[i].Distance < closest[j].Distance
	}
	log.Printf("%#v\n", closest)
	sort.Slice(closest, sortfunc)

	tw := tablewriter.NewWriter(w)
	tw.SetHeader([]string{"Target", "System", "Distance"})
	tdata := make([][]string, len(kc))
	for _, entry := range closest {
		humanName := codexToHuman[entry.Name]
		tdata = append(tdata, []string{humanName, entry.System, fmt.Sprintf("%.2fly", entry.Distance)})
	}
	tw.AppendBulk(tdata)

	tw.Render()
	return nil
}

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
