package dto

type ClassDto struct {
	Name                string `json:"name"`
	Description         string `json:"description"`
	ProficiencyBonus    int    `json:"proficiency_bonus"`
	HitDice             string `json:"hit_dice"`
	ArmorProficiencies  string `json:"armor_proficiencies"`
	WeaponProficiencies string `json:"weapon_proficiencies"`
	ToolProficiencies   string `json:"tool_proficiencies"`
	SpellcastingAbility string `json:"spellcasting_ability"`
}
