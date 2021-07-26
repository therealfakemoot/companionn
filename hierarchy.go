package companionn

import (
	"encoding/json"
	"fmt"
	"os"
)

// envelope type, these keys are POI superclasses" like "Anomaly", "Biology", etc
type HWrapper map[string]HClass

// these keys are classes like "Aleoids", "E-Type Anomaly"
type HClass map[string]HType

// HType is a specific POI class/species.
type HType map[string]HEntry

type HEntry struct {
	Category    string      `json:"category"`
	Entryid     int         `json:"entryid"`
	Name        string      `json:"name"`
	Platform    string      `json:"platform"`
	Reward      interface{} `json:"reward"`
	SubCategory string      `json:"sub_category"`
}

func LoadHierarchy(fn string) (HWrapper, error) {
	var hw HWrapper

	f, err := os.Open(fn)
	if err != nil {
		return hw, fmt.Errorf("couldn't open hierarchy file: %w", err)
	}

	dec := json.NewDecoder(f)
	err = dec.Decode(&hw)
	if err != nil {
		return hw, fmt.Errorf("couldn't parse hierarchy file: %w", err)
	}

	return hw, nil
}

func FlattenHierarchy(hw HWrapper) map[string]HEntry {
	m := make(map[string]HEntry)
	for _, class := range hw {
		for _, htype := range class {
			for k, hentry := range htype {
				m[k] = hentry
			}
		}
	}

	return m
}

func NameMapping() (map[string]HEntry, map[string]string, error) {
	humanToCodex := make(map[string]HEntry)
	codexToHuman := make(map[string]string)

	rawHierarchy, err := LoadHierarchy("json/hierarchy.json")
	if err != nil {
		return humanToCodex, codexToHuman, fmt.Errorf("error loading hierarchy data: %w", err)
	}

	humanToCodex = FlattenHierarchy(rawHierarchy)
	for k, v := range humanToCodex {
		codexToHuman[v.Name] = k
	}

	return humanToCodex, codexToHuman, nil
}
