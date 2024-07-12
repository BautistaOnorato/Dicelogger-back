package domain

type ArmorXCharacterData struct {
	ArmorXCharacterData_Id int   `json:"armorxcharacter_data_id"`
	Armor                  Armor `json:"armor"`
	CharacterData_Id       int   `json:"character_data_id"`
	Equipped               bool  `json:"equipped"`
}
