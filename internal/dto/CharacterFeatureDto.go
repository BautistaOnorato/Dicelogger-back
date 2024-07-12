package dto

type CreateCharacterFeatureDto struct {
	FeatureId   int `json:"feature_id"`
	CharacterId int `json:"character_id"`
}