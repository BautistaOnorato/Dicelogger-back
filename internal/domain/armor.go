package domain

type Armor struct {
	ArmorId        int    `json:"armor_id"`
	Material       string `json:"material"`
	Name           string `json:"name"`
	Weight         int    `json:"weight"`
	Price          int    `json:"price"`
	Category       string `json:"category"`
	ProtectionType string `json:"protection_type"`
	Description    string `json:"description"`
	Penalty        string `json:"penalty"`
	Strength       int    `json:"strength"`
	ArmorClass     int    `json:"armor_class"`
	DexBonus       string `json:"dex_bonus"`
	CampaignId     *int   `json:"campaign_id"`
}
