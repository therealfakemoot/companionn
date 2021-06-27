package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func FetchChallengeMap(name string) (ChallengeMap, error) {
	cm := make(ChallengeMap)

	u, err := url.Parse("https://us-central1-canonn-api-236217.cloudfunctions.net/get_cmdr_codex")
	q := url.Values{}
	q.Set("cmdr", name)
	u.RawQuery = q.Encode()
	if err != nil {
		return cm, fmt.Errorf("error parsing hard coded url, dummy: %w", err)
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return cm, fmt.Errorf("error fetching challenge report: %w", err)
	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&cm)
	if err != nil {
		return cm, fmt.Errorf("error decoding challenge json: %w", err)
	}

	return cm, nil
}

func FetchChallengeReport(name string) (ChallengeReport, error) {
	var cr ChallengeReport

	u, err := url.Parse("https://us-central1-canonn-api-236217.cloudfunctions.net/get_cmdr_codex")
	q := url.Values{}
	q.Set("cmdr", name)
	u.RawQuery = q.Encode()
	if err != nil {
		return cr, fmt.Errorf("error parsing hard coded url, dummy: %w", err)
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return cr, fmt.Errorf("error fetching challenge report: %w", err)
	}

	enc := json.NewDecoder(resp.Body)
	err = enc.Decode(&cr)
	if err != nil {
		return cr, fmt.Errorf("error parsing challenge report: %w", err)
	}

	fmt.Printf("%#v\n", cr)

	err = json.Unmarshal(cr.RawObjectives, &cr.Objectives)
	if err != nil {
		return cr, fmt.Errorf("error parsing objective listing: %w", err)
	}

	return cr, nil
}
