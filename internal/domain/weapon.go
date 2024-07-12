package domain

type Weapon struct {
	Weapon_Id int `json:"weapon_id"`
	Weapon_Type string `json:"weapon_type"`
	Name string `json:"name"`
    Weight int `json:"weight"`
	Price int `json:"price"`
    Category string `json:"category"`
	Reach string `json:"reach"`
	Description string `json:"description"`
	Damage string `json:"damage"`
	Versatile_Damage string `json:"veratile_damage"`
	Ammunition int `json:"ammunition"`
	Damage_Type string `json:"damage_type"`
	Campaign_Id *int `json:"campaign_id"`
}