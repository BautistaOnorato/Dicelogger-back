package class

var (
	QueryInsertClass = `INSERT INTO class (name, description, proficiency_bonus, hit_dice, armor_proficiencies, weapon_proficiencies, tool_proficiencies, spellcasting_ability) VALUES(?,?,?,?,?,?,?,?);`
	QueryGetAll      = `SELECT * from class;`
	QueryGetById     = `SELECT * FROM class WHERE class_id = ?;`
	QueryUpdateClass = `UPDATE class SET name = ?, description = ?, proficiency_bonus = ?, hit_dice = ?, armor_proficiencies = ?, weapon_proficiencies = ?, tool_proficiencies = ?, spellcasting_ability = ? WHERE class_id = ?;`
	QueryDeleteClass = `DELETE FROM class WHERE class_id = ?;`
)
