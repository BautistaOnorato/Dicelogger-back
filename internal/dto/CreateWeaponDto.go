package dto

type CreateWeaponDto struct {
	WeaponType      string `json:"weapon_type"`
	Name            string `json:"name"`
	Weight          int    `json:"weight"`
	Price           int    `json:"price"`
	Category        string `json:"category"`
	Reach           string `json:"reach"`
	Description     string `json:"description"`
	Damage          string `json:"damage"`
	VersatileDamage string `json:"versatile_damage"`
	Ammunition      int    `json:"ammunition"`
	DamageType      string `json:"damage_type"`
	Basic           int    `json:"basic"`
}
