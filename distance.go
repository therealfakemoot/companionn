package companionn

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/twpayne/go-geom"
)

type Coords struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type EDSMSystem struct {
	Name         string `json:"name"`
	Coords       Coords `json:"coords"`
	Coordslocked bool   `json:"coordsLocked"`
}

func (s EDSMSystem) GetCoords() []float64 {
	return []float64{s.Coords.X, s.Coords.Y, s.Coords.Z}
}

func GetSystem(system string, cache map[string]EDSMSystem) (EDSMSystem, error) {
	if s, ok := cache[system]; ok {
		return s, nil
	}
	var s EDSMSystem
	u, err := url.Parse("https://www.edsm.net/api-v1/system?systemName=sol&showCoordinates=1")
	q := url.Values{}
	q.Set("systemName", system)
	q.Set("showCoordinates", "1")
	u.RawQuery = q.Encode()
	if err != nil {
		return s, fmt.Errorf("error parsing hard coded url, dummy: %w", err)
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return s, fmt.Errorf("error fetching system info for %s: %w", system, err)
	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&s)
	if err != nil {
		return s, fmt.Errorf("error decoding system json for %s: %w", system, err)
	}

	cache[system] = s
	log.Printf("Added %q to cache", s.Name)
	return s, nil
}

func EDSMDistance(s1, s2 string, cache map[string]EDSMSystem) (float64, error) {
	line := geom.NewLineString(geom.XYZ)
	a, err := GetSystem(s1, cache)
	if err != nil {
		return 0.0, err
	}

	b, err := GetSystem(s2, cache)
	if err != nil {
		return 0.0, err
	}

	line.SetCoords([]geom.Coord{a.GetCoords(), b.GetCoords()})

	return line.Length(), nil
}

func dirtyDump(fn string, entries []CodexEntry) {
	f, _ := os.OpenFile("dirtyDump.json", os.O_RDWR|os.O_TRUNC, 0644)
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.Encode(entries)

}

func CodexDistance(s1 string, s2 CodexEntry, cache map[string]EDSMSystem) (float64, error) {
	line := geom.NewLineString(geom.XYZ)
	origin, err := GetSystem(s1, cache)
	if err != nil {
		log.Printf("error fetching reference system: %s\n", err)
		return 0.0, err
	}

	x, err := s2.X.Float64()
	if err != nil {
		return 0.0, fmt.Errorf("error loading X coord for %s: %w", s2.System, err)
	}
	y, err := s2.Y.Float64()
	if err != nil {
		return 0.0, fmt.Errorf("error loading Y coord for %s: %w", s2.System, err)
	}
	z, err := s2.Z.Float64()
	if err != nil {
		return 0.0, fmt.Errorf("error loading Z coord for %s: %w", s2.System, err)
	}

	line.SetCoords([]geom.Coord{origin.GetCoords(), []float64{x, y, z}})

	if line.Length() == 0 && s2.System != "Sol" {
		// c := s2.Coords()
		// log.Printf("zero-distance found: %s (%f,%f,%f)\n", s2.System, c[0], c[1], c[2])
	}
	return line.Length(), nil
}
