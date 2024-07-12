package domain

type WeaponXCharacterData struct {
	Character_Weapon_Id int `json:"character_weapon_id"`
	CharacterData_Id int `json:"character_data_id"`
	Weapon Weapon `json:"weapon"`
	Equipped bool `json:"equipped"`
}