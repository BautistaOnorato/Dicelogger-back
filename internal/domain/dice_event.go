package domain

import "time"

type DiceEvent struct {
	DiceEventId      int       `json:"dice_event_id"`
	Stat             string    `json:"stat"`
	Difficulty       int       `json:"difficulty"`
	DiceRolled       string    `json:"dice_rolled"`
	DiceResult       int       `json:"dice_result"`
	EventProtagonist int       `json:"event_protagonist"`
	Description      string    `json:"description"`
	SessionId        int       `json:"session_id"`
	TimeStamp        time.Time `json:"time_stamp"`
}
