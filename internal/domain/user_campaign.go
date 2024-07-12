package domain

type UserCampaign struct {
	UserCampaignId int    `json:"user_campaign_id"`
	CampaignId     int    `json:"campaign_id"`
	UserId         string `json:"user_id"`
	CharacterId    *int   `json:"character_id"`
	IsOwner        int    `json:"is_owner"`
}
