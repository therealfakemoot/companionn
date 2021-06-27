package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Coords struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type System struct {
	Name         string `json:"name"`
	Coords       Coords `json:"coords"`
	Coordslocked bool   `json:"coordsLocked"`
}

func GetSystem(system string) (System, error) {
	var s System
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
		return s, fmt.Errorf("error fetching challenge report: %w", err)
	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&s)
	if err != nil {
		return s, fmt.Errorf("error decoding challenge json: %w", err)
	}

	return s, nil
}
