package skill

var (
	QueryCreateSkill      = `INSERT INTO skill (name, stat) VALUES(?,?);`
	QueryGetAll           = `SELECT * FROM skill;`
	QueryGetById          = `SELECT * FROM skill WHERE skill_id = ?;`
	QueryGetByCharacterId = `SELECT skill.* FROM skill INNER JOIN character_skill ON character_skill.skill_id = skill.skill_id WHERE character_skill.character_id = ?;`
	// QueryGetByCampaignId = ``
	QueryGetByClassId        = `SELECT skill.* FROM skill INNER JOIN class_skill ON class_skill.skill_id = skill.skill_id WHERE class_skill.class_id = ?;`
	QueryUpdate              = `UPDATE skill SET name = "", stat = "" WHERE skill_id = ?;`
	QueryDelete              = `DELETE FROM skill WHERE skill_id = ?;`
	QueryDeleteByCharacterId = `DELETE skill FROM skill JOIN character_skill ON character_skill.skill_id = skill.skill_id WHERE character_skill.character_id = ?;`
)
