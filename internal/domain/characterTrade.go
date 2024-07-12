package domain

type CharacterTrade struct {
	CharacterTrade_Id int `json:"character_trade_id"`
	TradeEvent_Id int `json:"trade_event_id"`
	WeaponXCharacter *int `json:"weapon_x_character"`
	ItemXCharacter *int `json:"item_x_character"` 
	ArmorXCharacter *int `json:"armor_x_character"`
	ItemOwner int `json:"item_owner"`
	ItemReciever int `json:"item_receiver"`
	Quantity *int `json:"quantity"`
	ItemName string `json:"item_name"`
	ItemType string `json:"item_type"`
}