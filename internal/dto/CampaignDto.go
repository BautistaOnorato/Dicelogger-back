package dto

import (
	"github.com/proyecto-dnd/backend/internal/domain"
)

type ResponseCampaignDto struct {
	CampaignId    int                        `json:"campaign_id"`
	DungeonMaster string                     `json:"dungeon_master"`
	Name          string                     `json:"name"`
	Description   string                     `json:"description"`
	Image         string                     `json:"image"`
	Notes         *string                    `json:"notes"`
	Status        *string                    `json:"status"`
	Sessions      []domain.Session           `json:"sessions"`
	Images        *string                    `json:"images"`
	Users         []UserCharacterCampaignDto `json:"users"`
}

type CreateCampaignDto struct {
	DungeonMaster string  `json:"dungeon_master"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Image         string  `json:"image"`
	Notes         *string `json:"notes"`
	Status        *string `json:"status"`
	Images        *string `json:"images"`
}

type UserCharacterCampaignDto struct {
	Id          string           `json:"id"`
	Username    string           `json:"name"`
	Email       string           `json:"email"`
	DisplayName string           `json:"displayName"`
	Image       *string          `json:"image"`
	Character   *CharacterCardDto `json:"character"`
}
