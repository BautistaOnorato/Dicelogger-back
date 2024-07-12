package dto

type SpellDto struct {
	Name            string    `json:"name"`
	Description     string `json:"description"`
	Range           int    `json:"range"`
	Ritual          bool   `json:"ritual"`
	Duration        string `json:"duration"`
	Concentration   bool   `json:"concentration"`
	CastingTime     string `json:"casting_time"`
	Level           int    `json:"level"`
	DamageType      string `json:"damage_type"`
	DifficultyClass int    `json:"difficulty_class"`
	Aoe             int    `json:"aoe"`
	School          string `json:"school"`
}
