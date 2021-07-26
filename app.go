package companionn

import (
	"fmt"
	// "github.com/sirupsen/logrus"
)

type System interface {
	Name() string
	Coords() []float64
}

type CoordinateRepository interface {
	Get(string) []float64
	Set(string, []float64)
}

type Helper struct {
	Name        string
	Entries     []CodexEntry
	KeyedCodex  KeyedCodex
	SystemCache CoordinateRepository
}

func NewHelper(name, challenges string, cache CoordinateRepository) (*Helper, error) {
	var nh Helper
	nh.Name = name
	entries, err := LoadEntries(challenges)
	if err != nil {
		return &nh, fmt.Errorf("couldn't load codex entries: %w", err)
	}
	nh.Entries = entries

	return &nh, nil
}
