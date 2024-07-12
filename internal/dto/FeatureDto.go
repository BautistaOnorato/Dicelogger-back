package dto

import "github.com/proyecto-dnd/backend/internal/domain"

type CreateFeatureDto struct {
	CharacterId int    `json:"character_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type FeatureFullResponseDto struct {
	CharacterId int              `json:"character_id"`
	Features    []domain.Feature `json:"features"`
}
