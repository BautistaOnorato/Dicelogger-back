package dto

import (
	"time"

	"github.com/proyecto-dnd/backend/internal/domain"
)

type CreateAttackEventDto struct {
	Type               string     `json:"type"`
	Environment        string     `json:"environment"`
	Session_id         int        `json:"session_id"`
	EventProtagonistId int        `json:"event_protagonist_id"`
	EventResolution    string     `json:"event_resolution"`
	Weapon             *int       `json:"weapon"`
	Spell              *int       `json:"spell"`
	DmgType            *string    `json:"dmg_type"`
	Description        *string    `json:"description"`
	Timestamp          *time.Time `json:"timestamp"`
}

type RepositoryResponseAttackEvent struct {
	AttackEventId            int        `json:"attack_event_id"`
	Type                     string     `json:"type"`
	Environment              string     `json:"environment"`
	Session_id               int        `json:"session_id"`
	EventProtagonistId       int        `json:"event_protagonist_id"`
	EventResolution          string     `json:"event_resolution"`
	Weapon                   *int       `json:"weapon"`
	Spell                    *int       `json:"spell"`
	DmgType                  *string    `json:"dmg_type"`
	Description              *string    `json:"description"`
	TimeStamp                *time.Time `json:"time_stamp"`
	SessionSessionId         int        `json:"session_session_id"`
	Start                    time.Time  `json:"start"`
	End                      time.Time  `json:"end"`
	SessionDescription       string     `json:"session_description"`
	SessionCampaignId        int        `json:"session_campaign_id"`
	SessionCurrentEnviroment *string    `json:"session_current_enviroment"`
	CharacterId              int        `json:"character_id"`
	CharacterUserId          *string    `json:"character_user_id"`
	CharacterCampaignId      *int       `json:"campaign_id"`
	CharacterName            string     `json:"character_name"`
	RaceName                 string     `json:"race_name"`
	ClassName                string     `json:"class_name"`
	Level                    int        `json:"level"`
	HitPoints                int        `json:"hit_points"`
}

type ResponseEventDto struct {
	AttackEventId    int                `json:"attack_event_id"`
	Type             string             `json:"type"`
	Environment      string             `json:"environment"`
	Session          domain.Session     `json:"session"`
	EventProtagonist CharacterCardDto   `json:"event_protagonist"`
	EventResolution  string             `json:"event_resolution"`
	Weapon           *int               `json:"weapon"`
	Spell            *int               `json:"spell"`
	DmgType          *string            `json:"dmg_type"`
	Description      *string            `json:"description"`
	TimeStamp        *time.Time         `json:"time_stamp"`
	Affected         []CharacterCardDto `json:"affected"`
}
