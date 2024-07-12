package domain

import (
	"time"
)

type AttackEvent struct {
	AttackEventId      int        `json:"attack_event_id"`
	Type               string     `json:"type"`
	Environment        string     `json:"environment"`
	Session_id         int        `json:"session_id"`
	EventProtagonistId int        `json:"event_protagonist_id"`
	EventResolution    string     `json:"event_resolution"`
	Weapon             *int       `json:"weapon"`
	Spell              *int       `json:"spell"`
	DmgType            *string    `json:"dmg_type"`
	Description        *string    `json:"description"`
	TimeStamp          *time.Time `json:"time_stamp"`
}
