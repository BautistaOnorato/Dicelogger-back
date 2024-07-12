package skillxcharacterdata

var (
	QueryCreateSkillXCharacter = `INSERT INTO character_skill (character_id, skill_id) values(?,?);`
	QueryDeleteSkillXCharacter = `DELETE FROM character_skill WHERE character_id = ? AND skill_id = ?;`
	QueryDeleteByCharacterId = `DELETE FROM character_skill WHERE character_id = ?;`
)
