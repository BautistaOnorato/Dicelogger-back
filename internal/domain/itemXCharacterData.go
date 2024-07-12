package domain

type ItemXCharacterData struct {
	Character_Item_Id int `json:"character_item_id"`
	CharacterData_Id int `json:"character_data_id"`
	Item Item `json:"item"`
	Quantity int `json:"quantity"`
}