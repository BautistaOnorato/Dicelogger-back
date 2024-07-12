package proficiency

var (
	QueryInsertProficiency    = `INSERT INTO proficiency (name, type) VALUES(?,?);`
	QueryGetAll               = `SELECT * from proficiency;`
	QueryGetById              = `SELECT * FROM proficiency WHERE proficiency_id = ?;`
	QueryGetByCharacterDataId = `SELECT proficiency.* FROM proficiency INNER JOIN character_proficiency ON character_proficiency.proficiency_id = proficiency.proficiency_id WHERE character_proficiency.character_id = ?;`
	QueryUpdate               = `UPDATE proficiency SET name = ?, type = ?;`
	QueryDelete               = `DELETE FROM proficiency WHERE proficiency_id = ?;`
)
