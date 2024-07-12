package dto

type CharacterXAttackEventDto struct {
	CharacterId           int    `json:"character_id"`
	EventId               int    `json:"event_id"`
	Dmg                   int    `json:"dmg"`
	DmgRoll               string `json:"dmg_roll"`
	AttackResult          int    `json:"attack_result"`
	AttackRoll            string `json:"attack_roll"`
	ArmorClass            int    `json:"armor_class"`
}