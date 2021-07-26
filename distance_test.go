package helper

import (
	// "fmt"
	"math"
	"testing"
)

func withTolerance(a, b float64) bool {
	tolerance := 0.001
	diff := math.Abs(a - b)
	return diff < tolerance
}

func Test_CodexDistance(t *testing.T) {
	e := CodexEntry{IndexID: interface{}(nil), HudCategory: "Biology", EnglishName: "Clypeus Speculumi - Grey", CreatedAt: "2021-06-13 14:41:46", ReportedAt: "2021-06-13 14:41:45", Cmdrname: "thorhammer7", System: "Swoiwns DQ-D c2", X: "-739.25000", Y: "-541.34375", Z: "220.65625", Body: "Swoiwns DQ-D c2 2 b", Latitude: "29.84768", Longitude: "96.76928", Entryid: 2340306, Name: "$Codex_Ent_Clypeus_03_K_Name;", NameLocalised: "Clypeus Speculumi - Grau", Category: "$Codex_Category_Biology;", CategoryLocalised: "Biologisch und geologisch", SubCategory: "$Codex_SubCategory_Organic_Structures;", SubCategoryLocalised: "Organische Strukturen", RegionName: "$Codex_RegionName_18;", RegionNameLocalised: "Inner Orion Spur", Distance: 0}
	cache := make(map[string]EDSMSystem)
	d, err := CodexDistance("Sol", e, cache)
	if err != nil {
		t.Logf("error calculating CodexDistance: %s", err)
		t.Fail()
	}

	expected := 916.266128
	if !withTolerance(d, expected) {
		t.Logf("expected %f, got %f", expected, d)
		t.Fail()
	}

}
