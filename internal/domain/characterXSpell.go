package domain

type CharacterXSpell struct {
	CharacterSpellId int `json:"character_spell_id"`
	CharacterId      int `json:"character_id"`
	SpellId          int `json:"spell_id"`
}
