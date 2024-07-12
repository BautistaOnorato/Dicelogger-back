package dto

type AddCharacterToCampaignDto struct {
	CharacterId int `json:"character_id"`
	CampaignId  int `json:"campaign_id"`
}