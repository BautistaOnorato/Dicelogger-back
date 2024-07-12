package domain

type Campaign struct {
	CampaignId    int     `json:"campaign_id"`
	DungeonMaster string  `json:"dungeon_master"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Image         string  `json:"image"`
	Notes         *string `json:"notes"`
	Status        *string `json:"status"`
	Images        *string `json:"images"`
}
