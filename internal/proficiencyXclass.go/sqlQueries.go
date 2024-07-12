package proficiencyXclass

var (
	QueryInsert = `INSERT INTO class_proficiency (class_id, proficiency_id) VALUES(?,?);`
	QueryDelete = `DELETE FROM class_proficiency WHERE class_id = ? AND proficiency_id = ?`
)
