package helper

import (
	"encoding/json"
	"fmt"
	"io"
	// "log"
	"os"
	"sort"

	// "github.com/davecgh/go-spew/spew"
	"github.com/olekukonko/tablewriter"
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

func (kc KeyedCodex) Render(w io.Writer, count int) error {
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
	sort.Slice(closest, sortfunc)

	tw := tablewriter.NewWriter(w)
	tw.SetHeader([]string{"Target", "System", "Body", "Distance"})
	tdata := make([][]string, len(kc))
	for _, entry := range closest[:count-1] {
		humanName := codexToHuman[entry.Name]
		dist := fmt.Sprintf("%.2fly", entry.Distance)
		tdata = append(tdata, []string{humanName, entry.System, entry.Body, dist})
	}
	tw.AppendBulk(tdata)

	tw.Render()
	return nil
}

func (kc KeyedCodex) CalculateDistances(reference string, cache map[string]EDSMSystem) error {
	for cat, entries := range kc {
		for idx, entry := range entries {
			d, err := CodexDistance(reference, entry, cache)
			if err != nil {
				return fmt.Errorf("failed to calculate distance from %s to %s: %w", reference, entry.Name, err)
			}

			if d == 0.0 && entry.System != reference {
				return fmt.Errorf("zero distance calculated for non-self system: %s -> %s\n", entry.System, reference)
			}

			entry.Distance = d
			kc[cat][idx] = entry

		}

	}
	return nil
}

func (kc *KeyedCodex) Sort() {
	for _, entries := range *kc {
		sortfunc := func(i, j int) bool {
			return entries[i].Distance < entries[j].Distance
		}
		sort.Slice(entries, sortfunc)
	}

}
