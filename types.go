package companionn

import (
	"encoding/json"
	"fmt"
)

type Challenge struct {
	Cmdr  int     `json:"cmdr"`
	Codex int     `json:"codex"`
	Pct   float64 `json:"pct"`
}

func (c Challenge) String() string {
	return fmt.Sprintf("Total: %.2f", c.Pct)
}

type CodexCategory struct {
	HudCategory    string   `json:"hud_category"`
	Visited        int      `json:"visited"`
	Available      int      `json:"available"`
	TypesFound     []string `json:"types_found"`
	TypesAvailable []string `json:"types_available"`
	Percentage     float64  `json:"percentage"`
}

func (cc CodexCategory) Remaining() []string {
	m := make(map[string]struct{})
	diff := make([]string, 0)

	for _, s := range cc.TypesFound {
		m[s] = struct{}{}
	}

	for _, s := range cc.TypesAvailable {
		_, ok := m[s]
		if !ok {
			diff = append(diff, s)
		}
	}
	return diff
}

func (cc CodexCategory) String() string {
	return fmt.Sprintf("<%s: %.2f>", cc.HudCategory, cc.Percentage)
}

type ChallengeMap map[string]CodexCategory

type ChallengeReport struct {
	Challenge Challenge `json:"challenge"`
	Tourist   Challenge `json:"Tourist"`

	Objectives map[string]CodexCategory

	RawObjectives json.RawMessage
	/*
		KTypeAnomaly     CodexCategory `json:"K-Type Anomaly"`
		Anomaly          CodexCategory `json:"Anomaly"`
		ETypeAnomaly     CodexCategory `json:"E-Type Anomaly"`
		TTypeAnomaly     CodexCategory `json:"T-Type Anomaly"`
		LTypeAnomaly     CodexCategory `json:"L-Type Anomaly"`
		PTypeAnomaly     CodexCategory `json:"P-Type Anomaly"`
		QTypeAnomaly     CodexCategory `json:"Q-Type Anomaly"`
		BrainTree        CodexCategory `json:"Brain Tree"`
		Biology          CodexCategory `json:"Biology"`
		BarkMounds       CodexCategory `json:"Bark Mounds"`
		Anemone          CodexCategory `json:"Anemone"`
		Tubers           CodexCategory `json:"Tubers"`
		AmphoraPlant     CodexCategory `json:"Amphora Plant"`
		Shards           CodexCategory `json:"Shards"`
		Aleoids          CodexCategory `json:"Aleoids"`
		Bacterial        CodexCategory `json:"Bacterial"`
		Cactoid          CodexCategory `json:"Cactoid"`
		Clypeus          CodexCategory `json:"Clypeus"`
		Conchas          CodexCategory `json:"Conchas"`
		Electricae       CodexCategory `json:"Electricae"`
		Fonticulus       CodexCategory `json:"Fonticulus"`
		Fumerolas        CodexCategory `json:"Fumerolas"`
		Fungoids         CodexCategory `json:"Fungoids"`
		Osseus           CodexCategory `json:"Osseus"`
		Recepta          CodexCategory `json:"Recepta"`
		Stratum          CodexCategory `json:"Stratum"`
		Tubus            CodexCategory `json:"Tubus"`
		Shrubs           CodexCategory `json:"Shrubs"`
		Tussocks         CodexCategory `json:"Tussocks"`
		LagrangeCloud    CodexCategory `json:"Lagrange Cloud"`
		Cloud            CodexCategory `json:"Cloud"`
		StormCloud       CodexCategory `json:"Storm Cloud"`
		MineralSpheres   CodexCategory `json:"Mineral Spheres"`
		IceCrystals      CodexCategory `json:"Ice Crystals"`
		SilicateCrystals CodexCategory `json:"Silicate Crystals"`
		MetallicCrystals CodexCategory `json:"Metallic Crystals"`
		CalcitePlates    CodexCategory `json:"Calcite Plates"`
		Peduncle         CodexCategory `json:"Peduncle"`
		Aster            CodexCategory `json:"Aster"`
		Void             CodexCategory `json:"Void"`
		CollaredPod      CodexCategory `json:"Collared Pod"`
		ChalicePod       CodexCategory `json:"Chalice Pod"`
		Gyre             CodexCategory `json:"Gyre"`
		Rhizome          CodexCategory `json:"Rhizome"`
		Quadripartite    CodexCategory `json:"Quadripartite"`
		Stolon           CodexCategory `json:"Stolon"`
		Mollusc          CodexCategory `json:"Mollusc"`
		Fumarole         CodexCategory `json:"Fumarole"`
		Geology          CodexCategory `json:"Geology"`
		Geyser           CodexCategory `json:"Geyser"`
		LavaSpout        CodexCategory `json:"Lava Spout"`
		GasVent          CodexCategory `json:"Gas Vent"`
		Guardian         CodexCategory `json:"Guardian"`
		Thargoid         CodexCategory `json:"Thargoid"`
		None             CodexCategory `json:"None"`
		Planets          CodexCategory `json:"Planets"`
	*/
}
