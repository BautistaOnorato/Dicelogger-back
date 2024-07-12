package dto

type CharacterCardDto struct {
	CharacterId int     `json:"character_id"`
	UserId      *string `json:"user_id"`
	CampaignID  *int    `json:"campaign_id"`
	ImageUrl    string  `json:"image_url"`
	Name        string  `json:"name"`
	Race        string  `json:"race"`
	Class       string  `json:"class"`
	Level       int     `json:"level"`
	HitPoints   int     `json:"hit_points"`
}
