package characterXspell

var (
	QueryInsert       = `INSERT INTO character_spell (character_id, spell_id) values(?,?);`
	QueryDelete       = `DELETE FROM character_spell WHERE character_spell_id=?;`
	QueryDeleteParams = `DELETE FROM character_spell WHERE character_id=? AND spell_id=?;`
	QueryDeleteByCharacterId = `DELETE FROM character_spell WHERE character_id=?;`
)
