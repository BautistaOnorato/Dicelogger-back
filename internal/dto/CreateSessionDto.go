package dto

import "time"

type CreateSessionDto struct {
	Start             time.Time `json:"start"`
	End               time.Time `json:"end"`
	Description       string    `json:"description"`
	CampaignId        *int      `json:"campaign_id"`
	CurrentEnviroment *string   `json:"current_enviroment"`
}
