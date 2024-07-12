package domain

import (
	"time"
)

type Session struct {
	SessionId         int       `json:"session_id"`
	Start             time.Time `json:"start"`
	End               time.Time `json:"end"`
	Description       string    `json:"description"`
	CampaignId        *int      `json:"campaign_id"`
	CurrentEnviroment *string   `json:"current_enviroment"`
}
