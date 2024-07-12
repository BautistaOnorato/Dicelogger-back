package spell

var (
	QueryCreateSpell          = "INSERT INTO spell (name, description, `range`, ritual, duration, concentration, casting_time, level, damage_type, difficulty_class, aoe, school) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	QueryGetAll               = `SELECT * FROM spell`
	QueryGetById              = `SELECT * FROM spell WHERE spell_id = ?`
	QueryGetByCharacterDataId = `SELECT spell.* FROM spell left join character_spell on character_spell.spell_id = spell.spell_id WHERE character_spell.character_id = ?;`
	QueryGetByClassId         = `SELECT spell.* FROM spell left join class_spell on class_spell.spell_id = spell.spell_id WHERE class_spell.class_id = ?;`
	QueryUpdate               = "UPDATE spell SET name=?, description=?, `range`=?, ritual=?, duration=?, concentration=?, casting_time=?, level=?,damage_type=?, difficulty_class=?, aoe=?, school=? WHERE spell_id = ?"
	QueryDelete               = `DELETE FROM spell WHERE id = ?`
)
