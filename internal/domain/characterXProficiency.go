package domain

type CharacterXProficiency struct {
	CharacterProficiencyId int `json:"character_proficiency_id"`
	CharacterId            int `json:"character_id"`
	ProficiencyId          int `json:"proficiency_id"`
}
